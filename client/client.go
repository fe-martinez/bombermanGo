package client

import (
	"bombman/model"
	"bombman/utils"
	"bombman/view"
	"encoding/gob"
	"fmt"
	"io"
	"log"
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
	err := SendCreateGameMessage(c.connection, c.playerID)
	if err != nil {
		log.Println(err)
	}
}

func (c *Client) sendJoinGameMessage(lobbyID string) {
	err := SendJoinGameMessage(c.connection, c.playerID, lobbyID)
	if err != nil {
		log.Println(err)
	}
}

func (c *Client) sendStartGameMessage() {
	SendStartGameMessage(c.connection, c.playerID)
}

func (c *Client) sendMainMenuMessage() {
	SendMainMenuMessage(c.connection, c.playerID)
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

func (c *Client) receiveLobbyID() error {
	dec := gob.NewDecoder(c.connection)
	//c.connection.SetReadDeadline(time.Now().Add(15 * time.Second))

	for {
		var msg utils.ServerMessage
		err := dec.Decode(&msg)
		if err != nil {
			if err == io.EOF {
				log.Println("No more messages to read")
				return err
			}
			// if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
			// 	log.Println("Timeout waiting for lobbyID")
			// 	return err
			// }

			log.Println("Error decoding lobby id message", err)
			continue
		}

		if msg.Action == utils.LobbyIDMessage {
			lobbyID, ok := msg.Data.(string)
			if !ok {
				log.Println("Lobby ID message is not a string")
				continue
			}
			c.lobbyId = strings.TrimSpace(lobbyID)
			log.Println("Received lobby ID:", c.lobbyId)
			return err
		}

		fmt.Println("Received non-lobby ID message, action:", msg.Action)
	}
}

func updateGame(conn net.Conn, game *model.Game) {
	log.Println("Starting to receive game updates")
	for {
		if game.State == model.Finished {
			log.Println("Ending game updates")
			return
		}

		updatedGame, err := receiveGameFromServer(conn)
		if err != nil {
			fmt.Println("Error al recibir el juego actualizado:", err)
			return
		}
		*game = *updatedGame
	}
}

func receivePlayerID(conn net.Conn) (string, error) {
	dec := gob.NewDecoder(conn)
	//conn.SetReadDeadline(time.Now().Add(15 * time.Second))

	for {
		var msg utils.ServerMessage
		err := dec.Decode(&msg)
		if err != nil {
			log.Println("Error decoding player id message", err)
			continue
		}

		// if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
		// 	log.Println("Timeout waiting for lobbyID")
		// 	return "", netErr
		// }

		if msg.Action == utils.PlayerIDMessage {
			playerID, ok := msg.Data.(string)
			if !ok {
				log.Println("Player ID message is not a string")
				continue
			}
			playerIDString := strings.TrimSpace(playerID)
			log.Println("Received lobby ID:", playerIDString)
			return playerIDString, nil
		}

		fmt.Println("Received non-lobby ID message, action:", msg.Action)
	}
}
