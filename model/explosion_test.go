package model

import (
	"log"
	"testing"
)

func TestNewExplosion(t *testing.T) {
	gameMap, err := CreateMap(MAP_PATH)
	if err != nil {
		t.Error("Error creating game map")
	}
	game := NewGame("1", gameMap)
	if game == nil {
		t.Error("Error creating game")
	} else {
		explosion := NewExplosion(Position{1, 1}, 1, *game)
		if explosion == nil {
			t.Error("Error creating explosion")
		}
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
	} else {
		position := Position{3, 3}
		var affectedTiles = lookForAffectedTiles(*game, position, 2)
		if len(affectedTiles) != 9 {
			t.Error("Error looking for affected tiles")
		}

		var expectedAffectedTiles = []Position{
			{1, 3},
			{2, 3},
			{3, 3},
			{4, 3},
			{5, 3},
			{3, 1},
			{3, 2},
			{3, 4},
			{3, 5},
		}

		for _, expectedTile := range expectedAffectedTiles {
			found := false
			for _, tile := range affectedTiles {
				if tile == expectedTile {
					found = true
					break
				}
			}
			if !found {
				t.Error("Error looking for affected tiles")
			}

		}
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
	} else {
		affectedTiles := getAffectedTiles(Position{1, 1}, 1, *game)
		log.Println(affectedTiles)
		if len(affectedTiles) != 4 {
			t.Error("Error getting affected tiles")
		}
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
	} else {
		position := Position{1, 1}
		explosion := NewExplosion(position, 1, *game)
		if explosion == nil {
			t.Error("Error creating explosion")
		}
		if !explosion.IsTileInRange(position) {
			t.Error("Tile should be in range")
		}
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
	} else {
		position := Position{1, 1}
		explosion := NewExplosion(position, 1, *game)
		if explosion == nil {
			t.Error("Error creating explosion")
		} else {
			explosion.AddAffectedPlayer("1")
			if len(explosion.AffectedPlayers) != 1 {
				t.Error("Error adding affected players")
			}
		}
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
	} else {
		position := Position{1, 1}
		explosion := NewExplosion(position, 1, *game)
		if explosion == nil {
			t.Error("Error creating explosion")
		} else {
			explosion.AddAffectedPlayer("1")
			if len(explosion.AffectedPlayers) != 1 {
				t.Error("Error adding affected players")
			}

			if !explosion.IsPlayerAlreadyAffected("1") {
				t.Error("Player should be already affected")
			}
		}
	}
}
