package client

import "bombman/view"

type MainMenuState struct{}

func (m *MainMenuState) Handle(c *Client) {
	view.DrawMainMenuScreen()

	input := handleMainMenuInput()
	if input == "create" {
		//go updateGame(c.connection, &c.game)
		//c.gameState = &PlayingState{}
		c.sendCreateGameMessage()
		c.receiveLobbyID()
		c.gameState = &WaitingMenuState{}
	} else if input == "join" {
		c.gameState = &LobbySelectionState{}
	}
}
