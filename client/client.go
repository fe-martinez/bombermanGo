package client

import (
	"bombman/model"
	"bombman/view"
	"fmt"
	"net"
	"os"
	"strings"
)

const SERVER_ADDRESS = "localhost:8080"

type Client struct {
	connection net.Conn
	playerID   string
	game       model.Game
}

func NewClient() *Client {
	connection := dial(SERVER_ADDRESS)
	playerID, err := receiveId(connection)
	if err != nil {
		fmt.Println("Error while receiving player id")
		return nil
	}
	fmt.Println("Connected to server with ID:", playerID)

	var game model.Game
	//	loadGame(connection, &game)

	return &Client{
		connection: connection,
		playerID:   playerID,
		game:       game,
	}
}

func (c *Client) sendMessages(input string) {
	if input != "none" {
		c.sendInput(input)
	} else {
		SendUpdateMessage(c.connection, c.playerID)
	}
}

func (c *Client) sendInput(input string) {
	if input == "bomb" {
		SendBombMessage(c.connection, c.playerID)
	} else {
		SendMoveMessage(input, c.connection, c.playerID)
	}
}

func (c *Client) sendLeaveMessage() {
	SendLeaveMessage(c.connection, c.playerID)
}

func (c *Client) Start() {
	defer c.connection.Close()
	view.InitWindow()
	go updateGame(c.connection, &c.game)

	for !view.WindowShouldClose() {
		view.DrawGame(c.game)
		input := handleInput()
		c.sendMessages(input)
		if view.WindowShouldClose() {
			c.sendLeaveMessage()
		}
	}
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
	buffer := make([]byte, 2048)
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
		updatedGame, err := receiveMessageFromServer(conn)
		if err != nil {
			fmt.Println("Error al recibir el juego actualizado:", err)
			return
		}
		*game = *updatedGame
	}
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
