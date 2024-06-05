package client

import (
	"bombman/view"
	"encoding/gob"
	"net"
)

type LobbySelectionState struct{}

func (l *LobbySelectionState) Handle(c *Client) {
	userInput, lobbyID := handleLobbySelectionInput()
	if userInput != "none" && len(lobbyID) == 3 {
		c.sendJoinGameMessage(lobbyID)
		ack, err := readLobbyAcknowledgeMessage(c.connection)

		if err != nil {
			view.DrawLobbySelectionScreen("Error while joining lobby")
			return
		}

		if !ack.Success {
			view.DrawLobbySelectionScreen(ack.Message)
			return
		} else {
			c.lobbyId = ack.LobbyID
			go updateGame(c.connection, &c.game)
			c.gameState = &WaitingMenuState{}
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
	var msg JoinLobbyAck

	dec := gob.NewDecoder(connection)
	err := dec.Decode(&msg)
	if err != nil {
		return JoinLobbyAck{}, err
	}

	return msg, nil
}
