package client

import "bombman/view"

type MainMenuState struct{}

func (m *MainMenuState) Handle(c *Client) {
	view.DrawMainMenuScreen()

	input := handleMainMenuInput()
	if input == "create" {
		c.sendCreateGameMessage()
		c.receiveLobbyID()
		c.gameState = &WaitingMenuState{}
	} else if input == "join" {
		c.gameState = &LobbySelectionState{}
	} else if input == "rules" {
		c.gameState = &RulesState{}
	} else if input == "controls" {
		c.gameState = &ControlRulesState{}
	}
}
