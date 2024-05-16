package main

import (
	"fmt"
	"net"
	"os"
	// "encoding/json"
	// "fmt"
	// "net"
	// "os"
)

const SERVER_ADDRESS = "localhost:8080"

func informUser(serverAddress string) {
	fmt.Println("Starting game server at", serverAddress)
}

func listen(serverAddress string) net.Listener {
	listener, err := net.Listen("tcp", serverAddress)
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
	return listener
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

func handleConnection(conn net.Conn, Game *Game) {
	defer conn.Close()
	playerID := createRandomUid()
	//Generar posición
	if !connectPlayer(playerID, Game) {
		fmt.Println("Game is full!")
		return
	}
	fmt.Println("Player connected:", playerID)

	sendId(conn, playerID)

	for {
		handleMessages(conn, Game)
	}
}

func readClientMessage(conn net.Conn) (ClientMessage, error) {
	// Leer los datos enviados por el cliente
	buffer := make([]byte, 1024)
	_, err := conn.Read(buffer)
	if err != nil {
		return ClientMessage{}, fmt.Errorf("error al leer del cliente: %s", err)
	}

	// Decodificar el mensaje recibido
	clientMsg, err := decodeClientMessage(buffer)
	if err != nil {
		return ClientMessage{}, fmt.Errorf("error al decodificar el mensaje del cliente: %s", err)
	}

	return clientMsg, nil
}

func sendMessageToClient(conn net.Conn, game *Game) {
	encodedGame, err := encodeGame(*game)
	if err != nil {
		fmt.Println("Error encoding game:", err)
		return
	}
	_, err = conn.Write(encodedGame)
	if err != nil {
		fmt.Println("Error al enviar el juego al cliente:", err)
		return
	}

	fmt.Println("Juego enviado exitosamente al cliente.")
}

func respondToClient(conn net.Conn, message ClientMessage, game *Game) {
	switch message.Action {
	case "bomb":
		//Place the bomb
	case "move":
		//Mover personaje
	case "update":
		//Enviar juego
		sendMessageToClient(conn, game)
	case "leave":
		game.removePlayer(message.ID)
		fmt.Println("Player left:", message.ID)
	default:
		fmt.Println("Unknown message action: ", message.Action)
	}
}

func handleMessages(conn net.Conn, game *Game) {
	message, err := readClientMessage(conn)
	if err != nil {
		return
	}
	fmt.Println("Message received: ", message)
	respondToClient(conn, message, game)
}

func sendId(conn net.Conn, playerID string) {
	_, err := conn.Write([]byte(playerID))
	if err != nil {
		fmt.Println("Error sending player ID to client: ", err)
	}
}

func connectPlayer(playerID string, Game *Game) bool {
	playerPosition := Game.generateValidPosition(15)
	player := createPlayer(playerID, playerPosition, Game)
	if player != nil {
		Game.addPlayer(player)
		return true
	}
	return false
}

func createPlayer(playerID string, position *Position, Game *Game) *Player {
	if Game.isFull() {
		return nil
	} else {
		return NewPlayer(playerID, position)
	}
}

func startServer() {
	informUser(SERVER_ADDRESS)

	listener := listen(SERVER_ADDRESS)

	defer listener.Close()

	GameMap := createMap(15, 16)
	GameId := createRandomUid()
	Game := createGame(GameId, GameMap)
	handleConnections(listener, Game)
}

func createGame(GameId string, GameMap *GameMap) *Game {
	return NewGame(GameId, GameMap)
}
