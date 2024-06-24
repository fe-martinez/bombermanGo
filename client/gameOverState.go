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
		c.gameState = &MainMenuState{}
		c.lobbyId = ""
		c.game = model.Game{}
	}

}
