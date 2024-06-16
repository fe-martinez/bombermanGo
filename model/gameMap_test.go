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

func TestNewGameMapError(t *testing.T) {
	_, err := CreateMap("notfound.txt")
	if err == nil {
		t.Error("Error was expected")
	}
}

func TestNewGameMapInvalidCharacter(t *testing.T) {
	_, err := CreateMap("test_resources/invalid_map.txt")
	if err == nil {
		t.Error("Error was expected")
	}
}

func TestAddPowerUp(t *testing.T) {
	gameMap, _ := CreateMap(MAP_PATH)
	gameMap.AddPowerUp(&Position{X: 1, Y: 1})
	if len(gameMap.PowerUps) != 1 {
		t.Error("PowerUp was not added correctly")
	}
}

func TestAddPowerUpInExistingPositionFails(t *testing.T) {
	gameMap, _ := CreateMap(MAP_PATH)
	gameMap.AddPowerUp(&Position{X: 1, Y: 1})
	if len(gameMap.PowerUps) != 1 {
		t.Error("PowerUp was not added correctly")
	}
	gameMap.AddPowerUp(&Position{X: 1, Y: 1})
	if len(gameMap.PowerUps) == 2 {
		t.Error("PowerUp was added when it shouldn't be")
	}
}

func TestRemovePowerUp(t *testing.T) {
	gameMap, _ := CreateMap(MAP_PATH)
	gameMap.AddPowerUp(&Position{X: 1, Y: 1})
	if len(gameMap.PowerUps) != 1 {
		t.Error("PowerUp was not added correctly")
	}
	gameMap.RemovePowerUp(Position{X: 1, Y: 1})
	if len(gameMap.PowerUps) != 0 {
		t.Error("PowerUp was not removed correctly")
	}
}

func TestRemoveNonExistingPowerUp(t *testing.T) {
	gameMap, _ := CreateMap(MAP_PATH)
	gameMap.AddPowerUp(&Position{X: 1, Y: 1})
	if len(gameMap.PowerUps) != 1 {
		t.Error("PowerUp was not added correctly")
	}
	gameMap.RemovePowerUp(Position{X: 0, Y: 0})
	if len(gameMap.PowerUps) != 1 {
		t.Error("PowerUp was not removed correctly")
	}
}

func TestIsUnbreakableWall_PositiveCase(t *testing.T) {
	gameMap, _ := CreateMap(MAP_PATH)
	if gameMap.isUnbreakableWall(Position{X: 0, Y: 0}) != true {
		t.Error("Wall is not unbreakable")
	}
}

func TestIsUnbreakableWall_NegativeCase(t *testing.T) {
	gameMap, _ := CreateMap(MAP_PATH)
	if gameMap.isUnbreakableWall(Position{X: 0, Y: 1}) != false {
		t.Error("Wall is not unbreakable")
	}
}


