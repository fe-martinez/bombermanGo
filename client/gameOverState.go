package client

import (
	"bombman/view"
)

type GameOverState struct{}

func (g *GameOverState) Handle(c *Client) {
	view.DrawGameOverScreen(c.game)
}
