package client

import "bombman/view"

type WaitingMenuState struct{}

func showDrawingMenu(c *Client) {
	var players []string

	for _, value := range c.game.Players {
		players = append(players, value.Username)
	}
	view.DrawWaitingMenu(players)
}

func (w *WaitingMenuState) Handle(c *Client) {
	showDrawingMenu(c)
	input := handleWaitingMenuInput()
	if input == "start" {
		go updateGame(c.connection, &c.game)
		c.gameState = &PlayingState{}
	}
}
