package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"

	"github.com/google/uuid"
)

// Action: conect or move
// Data: the command (up, down, left, right, bomb, etc)
type Message struct {
	Action   string
	Data     map[string]interface{}
	PlayerID string
}

const (
	serverAddress = "localhost:8080"
)

func informUser() {
	fmt.Println("Starting game server at", serverAddress)
}

func listen() net.Listener {
	listener, err := net.Listen("tcp", serverAddress)
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
	return listener
}
func createGame(gameId string) Game {
	gameMap := createMap(15, 16)
	game := initGame(gameMap, gameId)
	return game
}

func handleConnections(listener net.Listener, game *Game) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go handleConnection(conn, game)
	}
}

func startServer() {
	informUser()

	listener := listen()

	defer listener.Close()

	gameId := createRandomUid()
	game := createGame(gameId)

	handleConnections(listener, &game)
}

func createNewPlayer(game *Game, playerId string) bool {
	if len(game.Players) < 4 {
		createPlayer(game, playerId)
		return true
	}
	return false
}

func move(msg *Message, game *Game) {
	var moveData struct {
		Direction string
	}
	moveData.Direction = msg.Data["direction"].(string)
	// fmt.Println(moveData)
	// fmt.Println("move: ", moveData.Direction)

	movePlayer(game, moveData.Direction, msg.PlayerID)
}

func handleMessage(conn net.Conn, msg *Message, game *Game) {
	switch msg.Action {
	case "join":
		json.NewEncoder(conn).Encode(game)
	case "bomb":
		position := getPlayerPosition(msg.PlayerID, *game)
		placeBomb(position, msg.PlayerID, game)
		json.NewEncoder(conn).Encode(game)
	case "move":
		move(msg, game)
		json.NewEncoder(conn).Encode(game)
	case "update":
		json.NewEncoder(conn).Encode(game)
	case "leave":
		disconnectPlayer(game, msg.PlayerID)
	default:
		fmt.Println("Unknown message action: ", msg.Action)
	}
}

func createRandomUid() string {
	id, err := uuid.NewRandom()
	if err != nil {
		fmt.Println("Error generating UUID: ", err)
	}
	return id.String()
}

func handleMessages(conn net.Conn, game *Game) {
	decoder := json.NewDecoder(conn)
	for {
		var msg Message
		err := decoder.Decode(&msg)
		if err != nil {
			fmt.Println("Error decoding message: ", err)
			return
		}
		fmt.Println(game.Players)
		handleMessage(conn, &msg, game)
	}
}

func handleConnection(conn net.Conn, game *Game) {
	defer conn.Close()
	playerID := createRandomUid()
	created := createNewPlayer(game, playerID)
	if !created {
		fmt.Println("Game is full")
		return
	}
	encoder := json.NewEncoder(conn)
	encoder.Encode(playerID)

	handleMessages(conn, game)

}
