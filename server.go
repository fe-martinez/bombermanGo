package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
)

type Message struct {
	Action   string
	Data     map[string]interface{}
	PlayerID int
}

const (
	serverAddress = "localhost:8080"
)

func server() {
	fmt.Println("Starting game server at", serverAddress)

	listener, err := net.Listen("tcp", serverAddress)
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
	defer listener.Close()

	game := initGame(15, 16)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go handleConnection(conn, &game)
	}
}

func newPlayer(game *Game) int {
	return createPlayer(game)
}

func handleMessage(conn net.Conn, msg *Message, game *Game) {
	switch msg.Action {
	case "join":
		json.NewEncoder(conn).Encode(game)
	case "leave":
		//Handle leave
	case "move":
		var moveData struct {
			Direction string
		}
		moveData.Direction = msg.Data["direction"].(string)
		fmt.Println(moveData)
		fmt.Println("move: ", moveData.Direction)

		movePlayer(game, moveData.Direction, msg.PlayerID)
		json.NewEncoder(conn).Encode(game)
	case "update":
		json.NewEncoder(conn).Encode(game)
	default:
		fmt.Println("Unknown message action: ", msg.Action)
	}
}

func handleConnection(conn net.Conn, game *Game) {
	defer conn.Close()

	// When a player connects, we first return his ID.
	playerID := newPlayer(game)
	encoder := json.NewEncoder(conn)
	encoder.Encode(playerID)

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
