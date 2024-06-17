package model

import "testing"

func TestNewExplosion(t *testing.T) {
	gameMap, err := CreateMap(MAP_PATH)
	if err != nil {
		t.Error("Error creating game map")
	}
	game := NewGame("1", gameMap)
	if game == nil {
		t.Error("Error creating game")
	}
	explosion := NewExplosion(Position{1, 1}, 1, *game)
	if explosion == nil {
		t.Error("Error creating explosion")
	}
}

func TestLookForAffectedTiles(t *testing.T) {
	gameMap, err := CreateMap(MAP_PATH)
	if err != nil {
		t.Error("Error creating game map")
	}
	game := NewGame("1", gameMap)
	if game == nil {
		t.Error("Error creating game")
	}
	position := Position{1, 1}
	var affectedTiles = lookForAffectedTiles(*game, position, 1)

	//Esta no es la condición real, cambiarla
	if len(affectedTiles) != 1 {
		t.Error("Error looking for affected tiles")
	}
}

func TestGetAffectedTiles(t *testing.T) {
	gameMap, err := CreateMap(MAP_PATH)
	if err != nil {
		t.Error("Error creating game map")
	}
	game := NewGame("1", gameMap)
	if game == nil {
		t.Error("Error creating game")
	}
	affectedTiles := getAffectedTiles(Position{1, 1}, 1, *game)
	if len(affectedTiles) != 1 {
		t.Error("Error getting affected tiles")
	}
}

func TestIsTileInRange(t *testing.T) {
	gameMap, err := CreateMap(MAP_PATH)
	if err != nil {
		t.Error("Error creating game map")
	}
	game := NewGame("1", gameMap)
	if game == nil {
		t.Error("Error creating game")
	}
	position := Position{1, 1}
	explosion := NewExplosion(position, 1, *game)
	if explosion == nil {
		t.Error("Error creating explosion")
	}
	if !explosion.IsTileInRange(position) {
		t.Error("Tile should be in range")
	}
}

func TestAddAffectedPlayers(t *testing.T) {
	gameMap, err := CreateMap(MAP_PATH)
	if err != nil {
		t.Error("Error creating game map")
	}
	game := NewGame("1", gameMap)
	if game == nil {
		t.Error("Error creating game")
	}
	position := Position{1, 1}
	explosion := NewExplosion(position, 1, *game)
	if explosion == nil {
		t.Error("Error creating explosion")
	}
	explosion.AddAffectedPlayer("1")
	if len(explosion.AffectedPlayers) != 1 {
		t.Error("Error adding affected players")
	}
}

func TestIsPlayerAlreadyAffected(t *testing.T) {
	gameMap, err := CreateMap(MAP_PATH)
	if err != nil {
		t.Error("Error creating game map")
	}
	game := NewGame("1", gameMap)
	if game == nil {
		t.Error("Error creating game")
	}
	position := Position{1, 1}
	explosion := NewExplosion(position, 1, *game)
	if explosion == nil {
		t.Error("Error creating explosion")
	}
	explosion.AddAffectedPlayer("1")
	if len(explosion.AffectedPlayers) != 1 {
		t.Error("Error adding affected players")
	}

	if !explosion.IsPlayerAlreadyAffected("1") {
		t.Error("Player should be already affected")
	}
}
