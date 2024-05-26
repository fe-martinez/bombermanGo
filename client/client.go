package client

import (
	"bombman/model"
	"bombman/utils"
	"bombman/view"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
)

const SERVER_ADDRESS = "localhost:8080"

var mu sync.Mutex

type Client struct {
	connection net.Conn
	playerID   string
	game       model.Game
	gameState  string
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
		gameState:  "main-menu",
	}
}

func (c *Client) sendMessages(input string) {
	if input != "none" {
		c.sendInput(input)
	}
}

func (c *Client) sendInput(input string) {
	if input == "bomb" {
		SendBombMessage(c.connection, c.playerID)
	} else if input == "create" {
		SendCreateGameMessage(c.connection, c.playerID)
	} else if len(input) == 3 { // Esto esta mal. TO DO
		lobbyID := input
		fmt.Println("Lobby code:", lobbyID)
		SendJoinGameMessage(c.connection, c.playerID, lobbyID)
	} else {
		SendMoveMessage(input, c.connection, c.playerID)
	}
}

func (c *Client) sendLeaveMessage() {
	SendLeaveMessage(c.connection, c.playerID)
}

func (c *Client) handleMainMenu() {
	view.DrawMainMenuScreen()

	input := handleMainMenuInput()
	if input == "create" {
		go updateGame(c.connection, &c.game)
		c.gameState = "game"
		c.sendMessages("create")
	} else if input == "join" {
		// go updateGame(c.connection, &c.game)
		c.gameState = "lobby-selection"
	}
}

func (c *Client) handleLobbySelection() {
	view.DrawLobbySelectionScreen()

	userInput, lobbyID := handleLobbySelectionInput()
	if userInput != "none" && len(lobbyID) == 3 {
		c.sendMessages(lobbyID)
		c.gameState = "game"
		go updateGame(c.connection, &c.game)
	}

}

func (c *Client) handleGame() {
	fmt.Println(c.game)
	view.DrawGame(c.game)
	input := handleInput()
	c.sendMessages(input)
	if view.WindowShouldClose() {
		c.sendLeaveMessage()
	}
}

func (c *Client) Start() {
	defer c.connection.Close()
	view.InitWindow()

	for !view.WindowShouldClose() {
		if c.gameState == "main-menu" {
			c.handleMainMenu()
		} else if c.gameState == "lobby-selection" {
			c.handleLobbySelection()
		} else if c.gameState == "game" {
			c.handleGame()
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
	buffer := make([]byte, 9000)
	n, err := conn.Read(buffer)
	fmt.Println(n)
	if err != nil {
		return nil, fmt.Errorf("error al leer del servidor: %s", err)
	}
	mu.Lock()
	decodedGame, err := utils.DecodeGame(buffer[:n])
	mu.Unlock()
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
