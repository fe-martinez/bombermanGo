package model

import "testing"

func TestNewGameMap(t *testing.T) {
	gameMap, err := CreateMap(MAP_PATH)
	if err != nil {
		t.Error("Error creating game map")
	}
	if len(gameMap.Walls) != 34 || len(gameMap.PowerUps) != 0 || len(gameMap.Bombs) != 0 || gameMap.RowSize != 10 || gameMap.ColumnSize != 16 || len(gameMap.Bombs) != 0 && len(gameMap.PowerUps) != 0 {
		t.Error("GameMap was not created correctly")
	}
	if gameMap.Walls[0].Position.X != 0 || gameMap.Walls[0].Position.Y != 0 || gameMap.Walls[0].Indestructible != true {
		t.Error("Wall position or type is not correct")
	}
}
