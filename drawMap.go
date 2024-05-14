package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func drawGame(game Game) {
	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)

	for _, player := range game.Players {
		rl.DrawRectangle(int32(player.X*50), int32(player.Y*50), 50, 50, rl.Red)
	}

	for _, wall := range game.GameMap.Walls {
		rl.DrawRectangle(int32(wall.X*50), int32(wall.Y*50), 50, 50, rl.Gray)
	}

	for _, bomb := range game.GameMap.Bombs {
		rl.DrawRectangle(int32(bomb.X*50), int32(bomb.Y*50), 50, 50, rl.Blue)
	}

	for _, enemy := range game.GameMap.Enemies {
		rl.DrawRectangle(int32(enemy.X*50), int32(enemy.Y*50), 50, 50, rl.Green)
	}

	rl.EndDrawing()
}
