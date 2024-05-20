package model

import "testing"

func TestAddPlayer(t *testing.T) {
	player := NewPlayer("106835", &Position{1, 1})
	gameMap := CreateMap(15, 16)
	game := NewGame("1", gameMap)
	game.AddPlayer(player)
	if len(game.Players) != 1 {
		t.Error("Player was not added to the game")
	}
}
