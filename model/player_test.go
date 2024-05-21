package model

import "testing"

func TestNewPlayer(t *testing.T) {
	player := NewPlayer("1", &Position{1, 1})
	if player.ID != "1" || player.Lives != 6 || player.Bombs != 1 || len(player.PowerUps) != 0 || player.Position.X != 1 || player.Position.Y != 1 {
		t.Error("Player was not created correctly")
	}
}
