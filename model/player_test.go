package model

import "testing"

func TestNewPlayer(t *testing.T) {
	player := NewPlayer("1", &Position{1, 1})
	if player.ID != "1" || player.Lives != 6 || player.Bombs != 1 || len(player.PowerUps) != 0 || player.Position.X != 1 || player.Position.Y != 1 {
		t.Error("Player was not created correctly")
	}
}

func TestLoseHealth(t *testing.T) {
	player := NewPlayer("1", &Position{1, 1})
	var lives = player.Lives
	player.LoseHealth()
	if player.Lives != lives-1 {
		t.Error("Player did not lose health")
	}
}

func TestCanPlantBomb(t *testing.T) {
	player := NewPlayer("1", &Position{1, 1})
	if !player.CanPlantBomb() {
		t.Error("Player should be able to plant bomb")
	}
	player.Bombs = 0
	if player.CanPlantBomb() {
		t.Error("Player should not be able to plant bomb")
	}
}
