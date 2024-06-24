package client

import "bombman/view"

type RulesState struct{}

func (p *RulesState) Handle(c *Client) {
	view.DrawGameRules()
	input := handleRulesInput()
	if input == "back" {
		c.gameState = &MainMenuState{}
	}
}
