package main

import (
	// "encoding/json"
	"fmt"
	"os"

	// "log"
	"net"
)

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
func receiveMessageFromServer(conn net.Conn) (*Game, error) {
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		return nil, fmt.Errorf("error al leer del servidor: %s", err)
	}

	decodedGame, err := decodeGame(buffer[:n])
	if err != nil {
		return nil, fmt.Errorf("error al decodificar el juego del servidor: %s", err)
	}

	return decodedGame, nil
}

func updateGame(conn net.Conn, game *Game) {
	go func() {
		fmt.Println("Updating game...")

		updatedGame, err := receiveMessageFromServer(conn)
		if err != nil {
			fmt.Println("Error al recibir el juego actualizado:", err)
			return
		}

		*game = *updatedGame
	}()
}

func sendMessage(ClientMessage ClientMessage, connection net.Conn) {
	encoded, err := encodeClientMessage(ClientMessage)
	if err != nil {
		fmt.Println("Error encoding message:", err)
		return
	}
	_, err = connection.Write(encoded)
	if err != nil {
		fmt.Println("Error al enviar el mensaje:", err)
		return
	}

	fmt.Println("Mensaje enviado exitosamente.")

}

func sendMove(move string, connection net.Conn, playerID string) {
	ClientMessage := ClientMessage{Action: "move", Data: move, ID: playerID}
	sendMessage(ClientMessage, connection)
}

func sendBomb(connection net.Conn, playerID string) {
	ClientMessage := ClientMessage{Action: "bomb", Data: nil, ID: playerID}
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
	ClientMessage := ClientMessage{Action: "update", Data: nil, ID: playerID}
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
	ClientMessage := ClientMessage{Action: "leave", Data: nil, ID: playerID}
	sendMessage(ClientMessage, connection)
}

func startClient() {
	connection := dial(SERVER_ADDRESS)

	defer connection.Close()

	playerID := receiveId(connection)
	welcomeUser(playerID)

	initWindow()

	var game Game
	updateGame(connection, &game)

	for !WindowShouldClose() {
		//drawGame(game)
		drawGame2()
		sendMessages(connection, playerID)
		if WindowShouldClose() {
			sendLeaveMessage(connection, playerID)
		}
	}
}

func receiveId(conn net.Conn) string {
	buffer := make([]byte, 37)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading from server:", err)
		return ""
	}
	return string(buffer[:n])
}
