package client

import (
	"bombman/view"
	"slices"
)

type PlayingState struct{}

func (p *PlayingState) Handle(c *Client) {
	updateGame(c.connection, &c.game)
	view.DrawGame(c.game)
	input := handleInput()

	if len(input) > 0 {
		if slices.Contains(input, "bomb") {
			c.EmitEvent(EventPlaceBomb, "")
		}

		c.EmitEvent(EventMove, input)
	}

	if view.WindowShouldClose() {
		c.EmitEvent(EventLeave, "")
	}

	if c.game.State == "finished" {
		c.gameState = &GameOverState{}
	}
}
