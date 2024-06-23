package client

import "bombman/view"

type ControlRulesState struct{}

func (co *ControlRulesState) Handle(c *Client) {
	view.DrawControlsRules()
	input := handleControlRulesInput()
	if input == "back" {
		c.gameState = &MainMenuState{}
	}
}
