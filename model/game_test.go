package model

import "testing"

const MAP_PATH = "../data/test.txt"

func TestNewGame(t *testing.T) {
	gameMap, err := CreateMap(MAP_PATH)
	if err != nil {
		t.Error("Error creating game map")
	}
	game := NewGame("1", gameMap)
	if game.GameId != "1" || game.GameMap != gameMap || len(game.Players) != 0 || game.Round != 1 {
		t.Error("Game was not created correctly")
	}
}
func TestAddPlayer(t *testing.T) {
	player := NewPlayer("106835", &Position{1, 1})
	gameMap, err := CreateMap(MAP_PATH)
	if err != nil {
		t.Error("Error creating game map")
	}
	game := NewGame("1", gameMap)
	game.AddPlayer(player)
	if len(game.Players) != 1 {
		t.Error("Player was not added to the game")
	}
}

func TestRemovePlayer(t *testing.T) {
	player := NewPlayer("106835", &Position{1, 1})
	gameMap, err := CreateMap(MAP_PATH)
	if err != nil {
		t.Error("Error creating game map")
	}
	game := NewGame("1", gameMap)
	game.AddPlayer(player)
	if len(game.Players) != 1 {
		t.Error("Player was not added to the game")
	}
	game.RemovePlayer(player.ID)
	if len(game.Players) != 0 {
		t.Error("Player was not removed from the game")
	}
}

// No sé por qué falla xd lol
func TestGameIsFull(t *testing.T) {
	player1 := NewPlayer("106835", &Position{1, 1})
	player2 := NewPlayer("106835", &Position{1, 2})
	player3 := NewPlayer("106835", &Position{1, 3})
	player4 := NewPlayer("106835", &Position{1, 4})
	gameMap, err := CreateMap(MAP_PATH)
	if err != nil {
		t.Error("Error creating game map")
	}
	game := NewGame("1", gameMap)
	game.AddPlayer(player1)
	game.AddPlayer(player2)
	game.AddPlayer(player3)
	game.AddPlayer(player4)
	if !game.IsFull() {
		t.Error("Game should be full but it's not")
	}
}

func TestGameIsNotFull(t *testing.T) {
	player1 := NewPlayer("106835", &Position{1, 1})
	player2 := NewPlayer("106835", &Position{1, 2})
	player3 := NewPlayer("106835", &Position{1, 3})
	gameMap, err := CreateMap(MAP_PATH)
	if err != nil {
		t.Error("Error creating game map")
	}
	game := NewGame("1", gameMap)
	game.AddPlayer(player1)
	game.AddPlayer(player2)
	game.AddPlayer(player3)
	if game.IsFull() {
		t.Error("Game shouldn't be full but it is")
	}

}

func TestCollidesWithWalls(t *testing.T) {
	gameMap, err := CreateMap(MAP_PATH)
	if err != nil {
		t.Error("Error creating game map")
	}
	game := NewGame("1", gameMap)
	game.GameMap.Walls = append(game.GameMap.Walls, Wall{&Position{1, 1}, false})
	if !game.collidesWithWalls(Position{1, 1}) {
		t.Error("Player should collide with wall")
	}
	if game.collidesWithWalls(Position{2, 2}) {
		t.Error("Player should not collide with wall")
	}
}

func TestIsOutOfBounds(t *testing.T) {
	gameMap, err := CreateMap(MAP_PATH)
	if err != nil {
		t.Error("Error creating game map")
	}
	game := NewGame("1", gameMap)
	if !game.isOutOfBounds(Position{-1, 1}) {
		t.Error("Position should be out of bounds")
	}
	if game.isOutOfBounds(Position{1, 1}) {
		t.Error("Position should not be out of bounds")
	}
}

func TestCanMove(t *testing.T) {
	player := NewPlayer("106835", &Position{1, 1})
	gameMap, err := CreateMap(MAP_PATH)
	if err != nil {
		t.Error("Error creating game map")
	}
	game := NewGame("1", gameMap)
	game.GameMap.Walls = append(game.GameMap.Walls, Wall{&Position{1, 1}, false})
	if game.CanMove(player, 1, 1) {
		t.Error("Player should not be able to move")
	}
	if !game.CanMove(player, 2, 2) {
		t.Error("Player should be able to move")
	}
}

func TestIsValidPosition(t *testing.T) {
	gameMap, err := CreateMap(MAP_PATH)
	if err != nil {
		t.Error("Error creating game map")
	}
	game := NewGame("1", gameMap)
	if !game.IsValidPosition(Position{1, 1}) {
		t.Error("Position should be valid")
	}
	if game.IsValidPosition(Position{-1, 1}) {
		t.Error("Position should be invalid")
	}
	if game.IsValidPosition(Position{0, 0}) {
		t.Error("Position should be invalid")
	}
}

func TestPutBomb(t *testing.T) {
	player := NewPlayer("106835", &Position{1, 1})
	gameMap, err := CreateMap(MAP_PATH)
	if err != nil {
		t.Error("Error creating game map")
	}
	game := NewGame("1", gameMap)
	game.AddPlayer(player)
	game.PutBomb(player)
	if len(game.GameMap.Bombs) != 1 {
		t.Error("Bomb was not placed")
	}
}

func TestExplodeBomb(t *testing.T) {
	player := NewPlayer("106835", &Position{1, 1})
	gameMap, err := CreateMap(MAP_PATH)
	if err != nil {
		t.Error("Error creating game map")
	}
	game := NewGame("1", gameMap)
	game.AddPlayer(player)
	game.PutBomb(player)
	game.ExplodeBomb(&game.GameMap.Bombs[0])
	if len(game.GameMap.Bombs) != 0 {
		t.Error("Bomb was not exploded")
	}
}

func TestTransferPowerUpToPlayer(t *testing.T) {

}

func TestGrabPowerUp(t *testing.T) {
	player := NewPlayer("106835", &Position{1, 1})
	gameMap, err := CreateMap(MAP_PATH)
	if err != nil {
		t.Error("Error creating game map")
	}
	gameMap.AddPowerUp(&Position{1, 1})
	game := NewGame("1", gameMap)
	game.AddPlayer(player)
	game.GrabPowerUp(player.ID)
	if len(game.Players[player.ID].PowerUps) != 1 {
		t.Error("Power up was not grabbed")
	}
	if len(game.GameMap.PowerUps) != 0 {
		t.Error("Power up was not removed from the map")
	}
}

func TestPowerUpSpawn(t *testing.T) {
	gameMap, err := CreateMap(MAP_PATH)
	if err != nil {
		t.Error("Error creating game map")
	}
	game := NewGame("1", gameMap)
	game.PowerUpSpawn()
	game.PowerUpSpawn()
	game.PowerUpSpawn()
	game.PowerUpSpawn()
	game.PowerUpSpawn()
	if len(game.GameMap.PowerUps) != 4 {
		t.Error("Power ups were not spawned correctly")
	}
}
