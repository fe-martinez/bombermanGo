package client

import (
	"bombman/model"
	"bombman/utils"
	"bombman/view"
	"fmt"
	"net"
	"os"
	"strings"
)

const SERVER_ADDRESS = "localhost:8080"

func welcomeUser(playerID string) {
	fmt.Println("Connected to server with ID:", playerID)
}

func dial(serverAddress string) net.Conn {
	connection, err := net.Dial("tcp", serverAddress)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		os.Exit(1)
	}
	return connection
}
func receiveMessageFromServer(conn net.Conn) (*model.Game, error) {
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		return nil, fmt.Errorf("error al leer del servidor: %s", err)
	}

	decodedGame, err := model.DecodeGame(buffer[:n])
	if err != nil {
		return nil, fmt.Errorf("error al decodificar el juego del servidor: %s", err)
	}

	return decodedGame, nil
}

func updateGame(conn net.Conn, game *model.Game) {
	for {
		fmt.Println("Updating game...")

		updatedGame, err := receiveMessageFromServer(conn)
		if err != nil {
			fmt.Println("Error al recibir el juego actualizado:", err)
			return
		}

		*game = *updatedGame
	}
}

func sendMessage(ClientMessage utils.ClientMessage, connection net.Conn) {
	encoded, err := utils.EncodeClientMessage(ClientMessage)
	if err != nil {
		fmt.Println("Error encoding message:", err)
		return
	}
	_, err = connection.Write(encoded)
	if err != nil {
		fmt.Println("Error al enviar el mensaje:", err)
		return
	}
}

func sendMove(move string, connection net.Conn, playerID string) {
	ClientMessage := utils.ClientMessage{Action: "move", Data: move, ID: playerID}
	sendMessage(ClientMessage, connection)
}

func sendBomb(connection net.Conn, playerID string) {
	ClientMessage := utils.ClientMessage{Action: "bomb", Data: nil, ID: playerID}
	sendMessage(ClientMessage, connection)
}

func sendInput(input string, connection net.Conn, playerID string) {
	if input == "bomb" {
		sendBomb(connection, playerID)
	} else {
		sendMove(input, connection, playerID)
	}
}

func askForUpdates(connection net.Conn, playerID string) {
	ClientMessage := utils.ClientMessage{Action: "update", Data: nil, ID: playerID}
	sendMessage(ClientMessage, connection)
}

func sendMessages(connection net.Conn, playerID string) {
	input := handleInput()
	if input != "none" {
		sendInput(input, connection, playerID)
	} else {
		askForUpdates(connection, playerID)
	}
}

func sendLeaveMessage(connection net.Conn, playerID string) {
	ClientMessage := utils.ClientMessage{Action: "leave", Data: nil, ID: playerID}
	sendMessage(ClientMessage, connection)
}

func receiveId(conn net.Conn) (string, error) {
	buffer := make([]byte, 37)
	n, err := conn.Read(buffer)
	if err != nil {
		return "", err
	}
	id := string(buffer[:n])
	return strings.TrimSpace(id), nil
}

func StartClient() {
	connection := dial(SERVER_ADDRESS)
	defer connection.Close()

	playerID, err := receiveId(connection)
	if err != nil {
		fmt.Println("Error while receiving player ID: ", err)
	}
	welcomeUser(playerID)

	view.InitWindow()

	var game model.Game
	go updateGame(connection, &game)

	for !view.WindowShouldClose() {
		view.DrawGame(game)
		//drawGame2()
		sendMessages(connection, playerID)
		if view.WindowShouldClose() {
			sendLeaveMessage(connection, playerID)
		}
	}
}
