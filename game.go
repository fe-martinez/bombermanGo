package main

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	Players []Player
	Walls   []Wall
	Bombs   []Bomb
	size    int
}

type Player struct {
	X float32
	Y float32
}

type Bomb struct {
	X     float32
	Y     float32
	Timer float32
}

type Wall struct {
	X float32
	Y float32
}

func initGame(wallsAmount int, mapSize int) Game {
	walls := []Wall{}
	for i := 0; i < wallsAmount; i++ {
		walls = append(walls, Wall{
			X: float32(rand.Intn(mapSize)),
			Y: float32(rand.Intn(mapSize)),
		})
	}

	game := Game{
		Players: make([]Player, 0),
		Walls:   walls,
		Bombs:   make([]Bomb, 0),
		size:    mapSize,
	}

	return game
}

func createPlayer(game *Game) int {
	player := Player{}
	for {
		player.X = float32(rand.Intn(game.size))
		player.Y = float32(rand.Intn(game.size))
		if !checkCollision(player.X, player.Y, *game) {
			break
		}
	}
	game.Players = append(game.Players, player)
	return len(game.Players) - 1
}

func drawGame(game Game) {
	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)

	for _, player := range game.Players {
		rl.DrawRectangle(int32(player.X*50), int32(player.Y*50), 50, 50, rl.Red)
	}

	for _, wall := range game.Walls {
		rl.DrawRectangle(int32(wall.X*50), int32(wall.Y*50), 50, 50, rl.Gray)
	}

	for _, bomb := range game.Bombs {
		rl.DrawRectangle(int32(bomb.X*50), int32(bomb.Y*50), 50, 50, rl.Blue)
	}

	rl.EndDrawing()
}

func movePlayer(game *Game, direction string, playerNumber int) {
	switch direction {
	case "up":
		if !checkCollision(game.Players[playerNumber].X, game.Players[playerNumber].Y-0.05, *game) {
			game.Players[playerNumber].Y = game.Players[playerNumber].Y - 0.05
		}
	case "down":
		if !checkCollision(game.Players[playerNumber].X, game.Players[playerNumber].Y+0.05, *game) {
			game.Players[playerNumber].Y = game.Players[playerNumber].Y + 0.05
		}
	case "left":
		if !checkCollision(game.Players[playerNumber].X-0.05, game.Players[playerNumber].Y, *game) {
			game.Players[playerNumber].X = game.Players[playerNumber].X - 0.05
		}
	case "right":
		if !checkCollision(game.Players[playerNumber].X+0.05, game.Players[playerNumber].Y, *game) {
			game.Players[playerNumber].X = game.Players[playerNumber].X + 0.05
		}
	}
}

func checkCollision(x, y float32, game Game) bool {
	playerRect := rl.NewRectangle(x*50, y*50, 50, 50)

	for _, wall := range game.Walls {
		wallRect := rl.NewRectangle(wall.X*50, wall.Y*50, 45, 45)
		if rl.CheckCollisionRecs(playerRect, wallRect) {
			return true
		}
	}

	return false
}
