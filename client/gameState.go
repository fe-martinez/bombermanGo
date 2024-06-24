package client

import (
	"bombman/view"
)

type PlayingState struct{}

func (p *PlayingState) Handle(c *Client) {
	view.DrawGame(c.game)
	input := handleInput()

	if len(input) > 0 {
		c.sendGameInput(input)
	}

	if view.WindowShouldClose() {
		c.sendLeaveMessage()
	}

	if c.game.State == "finished" {
		c.gameState = &GameOverState{}
	}
}
