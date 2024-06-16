package client

import (
	"bombman/view"
)

type PlayingState struct{}

func (p *PlayingState) Handle(c *Client) {
	view.DrawGame(c.game)
	input := handleInput()
	c.sendGameInput(input)
	if view.WindowShouldClose() {
		c.sendLeaveMessage()
	}

	if c.game.State == "finished" {
		c.gameState = &GameOverState{}
	}
}
