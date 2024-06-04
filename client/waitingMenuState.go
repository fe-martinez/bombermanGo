package client

import (
	"bombman/view"
	"slices"
)

type WaitingMenuState struct{}

func showDrawingMenu(c *Client) {
	var players []string

	for _, value := range c.game.Players {
		players = append(players, value.Username)
	}

	slices.Sort(players)

	view.DrawWaitingMenu(players, c.lobbyId)
}

func (w *WaitingMenuState) Handle(c *Client) {
	showDrawingMenu(c)
	input := handleWaitingMenuInput()

	if input == "start" {
		c.sendStartGameMessage()
	}

	if c.game.State == "started" {
		c.gameState = &PlayingState{}
	}
}
