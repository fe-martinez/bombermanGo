package client

import (
	"bombman/view"
	"encoding/gob"
	"io"
	"log"
	"net"
)

type LobbySelectionState struct{}

func (l *LobbySelectionState) Handle(c *Client) {
	userInput, lobbyID := handleLobbySelectionInput()
	if userInput == "back" {
		c.gameState = &MainMenuState{}
	}

	if userInput != "none" && len(lobbyID) == 3 {
		c.sendJoinGameMessage(lobbyID)
		ack, err := readLobbyAcknowledgeMessage(c.connection)

		if err != nil {
			view.DrawLobbySelectionScreen("Error while joining lobby")
			log.Println(err)
			return
		}

		if !ack.Success {
			view.DrawLobbySelectionScreen(ack.Message)
			return
		} else {
			c.lobbyId = ack.LobbyID
			c.gameState = &WaitingMenuState{}
			updateGame(c.connection, &c.game)
		}
	}
	view.DrawLobbySelectionScreen(lobbyID)
}

type JoinLobbyAck struct {
	Success bool
	LobbyID string
	Message string
}

func readLobbyAcknowledgeMessage(connection net.Conn) (JoinLobbyAck, error) {
	dec := gob.NewDecoder(connection)

	for {
		var msg JoinLobbyAck
		err := dec.Decode(&msg)
		if err != nil {
			if err == io.EOF {
				log.Println("Reached EOF without finding a lobby ack")
				return JoinLobbyAck{}, err
			}
			log.Println("Error decoding message")
			continue
		}

		log.Println("Recieved JoinLobbyAck message")
		return msg, nil
	}
}
