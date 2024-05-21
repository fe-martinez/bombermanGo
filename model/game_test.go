package model

import "testing"

func TestNewGame(t *testing.T) {
	gameMap := CreateMap(15, 16)
	game := NewGame("1", gameMap)
	if game.GameId != "1" || game.GameMap != gameMap || len(game.Players) != 0 || game.Level != 1 {
		t.Error("Game was not created correctly")
	}
}
func TestAddPlayer(t *testing.T) {
	player := NewPlayer("106835", &Position{1, 1})
	gameMap := CreateMap(15, 16)
	game := NewGame("1", gameMap)
	game.AddPlayer(player)
	if len(game.Players) != 1 {
		t.Error("Player was not added to the game")
	}
}

func TestRemovePlayer(t *testing.T) {
	player := NewPlayer("106835", &Position{1, 1})
	gameMap := CreateMap(15, 16)
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
	gameMap := CreateMap(15, 16)
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
	gameMap := CreateMap(15, 16)
	game := NewGame("1", gameMap)
	game.AddPlayer(player1)
	game.AddPlayer(player2)
	game.AddPlayer(player3)
	if game.IsFull() {
		t.Error("Game shouldn't be full but it is")
	}

}

func TestCollidesWithWalls(t *testing.T) {
	gameMap := CreateMap(15, 16)
	game := NewGame("1", gameMap)
	game.GameMap.Walls = append(game.GameMap.Walls, Wall{&Position{1, 1}, false})
	if !game.collidesWithWalls(Position{1, 1}) {
		t.Error("Player should collide with wall")
	}
	if game.collidesWithWalls(Position{2, 2}) {
		t.Error("Player should not collide with wall")
	}
}
