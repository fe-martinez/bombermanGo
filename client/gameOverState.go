package client

import (
	"bombman/model"
	"bombman/view"
)

type GameOverState struct{}

func (g *GameOverState) Handle(c *Client) {
	view.DrawGameOverScreen(c.game)

	if handleGameOverInput() == "return" {
		c.sendMainMenuMessage()
		c.game = model.Game{}
		c.lobbyId = ""
		c.gameState = &MainMenuState{}
	}

}
