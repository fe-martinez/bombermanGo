package client

import (
	"bombman/view"
	"fmt"
)

type PlayingState struct{}

func (p *PlayingState) Handle(c *Client) {
	fmt.Println(c.game)
	view.DrawGame(c.game)
	input := handleInput()
	c.sendGameInput(input)
	if view.WindowShouldClose() {
		c.sendLeaveMessage()
	}
}
