package client

import "bombman/view"

type LobbySelectionState struct{}

func (l *LobbySelectionState) Handle(c *Client) {
	userInput, lobbyID := handleLobbySelectionInput()
	if userInput != "none" && len(lobbyID) == 3 {
		c.sendJoinGameMessage(lobbyID)
		c.gameState = &PlayingState{}
		go updateGame(c.connection, &c.game)
	}

	view.DrawLobbySelectionScreen(lobbyID)
}
