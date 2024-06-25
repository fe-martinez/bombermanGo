package client

import (
	"bombman/model"
	"bombman/view"
	"slices"
)

type WaitingMenuState struct{}

func (c *Client) gameShouldStart(input string) bool {
	return input == "start" && len(c.game.Players) > 1
}

func showDrawingMenu(c *Client) {
	var players []string

	for _, value := range c.game.Players {
		players = append(players, value.Username)
	}

	slices.Sort(players)

	view.DrawWaitingMenu(players, c.lobbyId)
}

func (w *WaitingMenuState) Handle(c *Client) {
	updateGame(c.connection, &c.game)
	showDrawingMenu(c)
	input := handleWaitingMenuInput()

	if c.gameShouldStart(input) {
		c.sendStartGameMessage()
	}

	if input == "back" {
		c.sendMainMenuMessage()
		c.game = model.Game{}
		c.lobbyId = ""
		c.gameState = &MainMenuState{}
		return
	}

	if c.game.State == "started" {
		c.gameState = &PlayingState{}
	}
}
