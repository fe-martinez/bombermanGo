package model

import "testing"

// Acá va a haber que testear que los mapas se creen según su pre-definición
func TestNewGameMap(t *testing.T) {
	gameMap := CreateMap(15, 16)
	if len(gameMap.Walls) != 15 || len(gameMap.PowerUps) != 0 || len(gameMap.Bombs) != 0 || gameMap.Size != 16 || len(gameMap.Bombs) != 0 && len(gameMap.PowerUps) != 0 {
		t.Error("GameMap was not created correctly")
	}
}
