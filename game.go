package main

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const BOMBTIME = 3

type Game struct {
	Players []Player
	GameMap Map
	level   int8
}

type Position struct {
	X float32
	Y float32
}

type Player struct {
	id string
	//position Position no anda el juego si uso esto xd lol
	X             float32
	Y             float32
	lives         int8
	maxBombs      int8
	bombQualities BombQualities
}

func initGame(gameMap Map) Game {
	game := Game{
		Players: make([]Player, 0),
		GameMap: gameMap,
	}

	return game
}

func createPlayer(game *Game, playerId string) {
	player := Player{}
	for {
		player.id = playerId
		player.X = float32(rand.Intn(game.GameMap.size))
		player.Y = float32(rand.Intn(game.GameMap.size))
		player.lives = 6
		if !checkCollision(player.X, player.Y, *game) {
			break
		}
	}
	game.Players = append(game.Players, player)
}

func getPlayerPositionInList(game *Game, id string) int {
	for i := range game.Players {
		if game.Players[i].id == id {
			return i
		}
	}
	return -1
}

func getPlayerPosition(clientID string, game Game) Position {
	position := Position{}
	for i := range game.Players {
		if game.Players[i].id == clientID {
			position.X = game.Players[i].X
			position.Y = game.Players[i].Y
		}
	}
	return position
}

func placeBomb(position Position, game *Game) {
	bomb := Bomb{
		X:     position.X,
		Y:     position.Y,
		Timer: BOMBTIME,
	}
	game.GameMap.Bombs = append(game.GameMap.Bombs, bomb)
}

func movePlayer(game *Game, direction string, playerId string) {
	playerPosition := getPlayerPositionInList(game, playerId)
	if playerPosition == -1 {
		return
	}

	switch direction {
	case "up":
		if !checkCollision(game.Players[playerPosition].X, game.Players[playerPosition].Y-0.05, *game) {
			game.Players[playerPosition].Y = game.Players[playerPosition].Y - 0.05
		}
	case "down":
		if !checkCollision(game.Players[playerPosition].X, game.Players[playerPosition].Y+0.05, *game) {
			game.Players[playerPosition].Y = game.Players[playerPosition].Y + 0.05
		}
	case "left":
		if !checkCollision(game.Players[playerPosition].X-0.05, game.Players[playerPosition].Y, *game) {
			game.Players[playerPosition].X = game.Players[playerPosition].X - 0.05
		}
	case "right":
		if !checkCollision(game.Players[playerPosition].X+0.05, game.Players[playerPosition].Y, *game) {
			game.Players[playerPosition].X = game.Players[playerPosition].X + 0.05
		}
	}
}

func checkCollision(x, y float32, game Game) bool {
	playerRect := rl.NewRectangle(x*50, y*50, 50, 50)

	for _, wall := range game.GameMap.Walls {
		wallRect := rl.NewRectangle(wall.X*50, wall.Y*50, 45, 45)
		if rl.CheckCollisionRecs(playerRect, wallRect) {
			return true
		}
	}

	//Chequear si esto está bien please
	for _, bomb := range game.GameMap.Bombs {
		bombRect := rl.NewRectangle(bomb.X*50, bomb.Y*50, 50, 50)
		if rl.CheckCollisionRecs(playerRect, bombRect) {
			return true
		}
	}

	// for _, enemy := range game.GameMap.Enemies {
	// 	enemyRect := rl.NewRectangle(enemy.X*50, enemy.Y*50, 50, 50)
	// 	if rl.CheckCollisionRecs(playerRect, enemyRect) {
	// 		return true
	// 	}
	// }

	return false
}
