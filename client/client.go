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
	"strings"
)

const SERVER_ADDRESS = "192.168.0.2:8080"

type Client struct {
	connection   net.Conn
	playerID     string
	lobbyId      string
	game         model.Game
	gameState    ClientState
	eventEmitter *EventEmitter
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
		connection:   connection,
		playerID:     playerID,
		lobbyId:      "",
		game:         game,
		gameState:    &MainMenuState{},
		eventEmitter: &EventEmitter{connection: connection},
	}
}

func (c *Client) Start() {
	defer c.connection.Close()
	view.InitWindow()

	for !view.WindowShouldClose() {
		c.gameState.Handle(c)
	}
}

func (c *Client) receiveLobbyID() error {
	dec := gob.NewDecoder(c.connection)

	for {
		var msg utils.ServerMessage
		err := dec.Decode(&msg)
		if err != nil {
			if err == io.EOF {
				log.Println("No more messages to read")
				return err
			}

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

func (c *Client) EmitEvent(eventType EventType, payload interface{}) error {
	return c.eventEmitter.Emit(GameEvent{
		Type:     eventType,
		PlayerID: c.playerID,
		Payload:  payload,
	})
}

func dial(serverAddress string) net.Conn {
	connection, err := net.Dial("tcp", serverAddress)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		os.Exit(1)
	}
	connection.(*net.TCPConn).SetNoDelay(true)
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

func updateGame(conn net.Conn, game *model.Game) error {
	updatedGame, err := receiveGameFromServer(conn)
	if err != nil {
		log.Println("Error al recibir el juego actualizado:", err)
	}
	*game = *updatedGame
	return nil
}

func receivePlayerID(conn net.Conn) (string, error) {
	dec := gob.NewDecoder(conn)

	for {
		var msg utils.ServerMessage
		err := dec.Decode(&msg)
		if err != nil {
			log.Println("Error decoding player id message", err)
			continue
		}

		if msg.Action == utils.PlayerIDMessage {
			playerID, ok := msg.Data.(string)
			if !ok {
				log.Println("Player ID message is not a string")
				continue
			}
			playerIDString := strings.TrimSpace(playerID)
			log.Println("Received playerID:", playerIDString)
			return playerIDString, nil
		}

		fmt.Println("Received non-playerID message, action:", msg.Action)
	}
}
