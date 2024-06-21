package client

import (
	"bombman/model"
	"bombman/utils"
	"bombman/view"
	"fmt"
	"net"
	"os"
	"slices"
	"strings"
)

const SERVER_ADDRESS = "localhost:8080"

type Client struct {
	connection net.Conn
	playerID   string
	lobbyId    string
	game       model.Game
	gameState  ClientState
}

func NewClient() *Client {
	connection := dial(SERVER_ADDRESS)
	playerID, err := receivePlayerID(connection)
	if err != nil {
		fmt.Println("Error while receiving player id")
		return nil
	}
	fmt.Println("Connected to server with ID:", playerID)

	var game model.Game

	return &Client{
		connection: connection,
		playerID:   playerID,
		lobbyId:    "",
		game:       game,
		gameState:  &MainMenuState{},
	}
}

func (c *Client) sendGameInput(input []string) {
	if slices.Contains(input, "bomb") {
		SendBombMessage(c.connection, c.playerID)
	} else {
		SendMoveMessage(input, c.connection, c.playerID)
	}
}

func (c *Client) sendLeaveMessage() {
	SendLeaveMessage(c.connection, c.playerID)
}

func (c *Client) sendCreateGameMessage() {
	SendCreateGameMessage(c.connection, c.playerID)
}

func (c *Client) sendJoinGameMessage(lobbyID string) {
	SendJoinGameMessage(c.connection, c.playerID, lobbyID)
}

func (c *Client) sendStartGameMessage() {
	SendStartGameMessage(c.connection, c.playerID)
}

func (c *Client) Start() {
	defer c.connection.Close()
	view.InitWindow()

	for !view.WindowShouldClose() {
		c.gameState.Handle(c)
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

func receiveGameFromServer(conn net.Conn) (*model.Game, error) {
	buffer := make([]byte, 9000)
	n, err := conn.Read(buffer)
	if err != nil {
		return nil, fmt.Errorf("error al leer del servidor: %s", err)
	}
	decodedGame, err := utils.DecodeGame(buffer[:n])
	if err != nil {
		return nil, fmt.Errorf("error al decodificar el juego del servidor: %s", err)
	}

	return decodedGame, nil
}

func (c *Client) receiveLobbyID() {
	buffer := make([]byte, 37)
	n, err := c.connection.Read(buffer)
	if err != nil {
		fmt.Println(err)
	}
	id := string(buffer[:n])
	cleanId := strings.TrimSpace(id)
	c.lobbyId = cleanId
}

func updateGame(conn net.Conn, game *model.Game) {
	for {
		updatedGame, err := receiveGameFromServer(conn)
		if err != nil {
			fmt.Println("Error al recibir el juego actualizado:", err)
			return
		}
		*game = *updatedGame
	}
}

func receivePlayerID(conn net.Conn) (string, error) {
	buffer := make([]byte, 37)
	n, err := conn.Read(buffer)
	if err != nil {
		return "", err
	}
	id := string(buffer[:n])
	return strings.TrimSpace(id), nil
}
