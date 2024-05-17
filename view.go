package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func initWindow() {
	rl.InitWindow(800, 800, "Bomberman Go!")
	rl.SetTargetFPS(60)
}

func WindowShouldClose() bool {
	return rl.WindowShouldClose()
}

// Función de prueba para dibujar en la ventana
func drawGame2() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)
	rl.DrawText("Welcome to Bomberman Go!", 190, 200, 20, rl.Maroon)
	rl.EndDrawing()
}

// Después se van a dibujar diferenciados, no todos iguales
func drawPlayers(game Game) {
	for _, player := range game.Players {
		rl.DrawRectangle(int32(player.Position.X*50), int32(player.Position.Y*50), 50, 50, rl.Red)
	}
}

func drawBombs(game Game) {
	for _, bomb := range game.GameMap.Bombs {
		rl.DrawRectangle(int32(bomb.Position.X*50), int32(bomb.Position.Y*50), 50, 50, rl.Blue)
	}
}

// Después va a tener que dibujar los distintos powerups según el tipo
func drawPowerUps(game Game) {
	for _, powerUp := range game.GameMap.PowerUps {
		rl.DrawRectangle(int32(powerUp.Position.X*50), int32(powerUp.Position.Y*50), 50, 50, rl.Brown)
	}
}

func drawWalls(game Game) {
	for _, wall := range game.GameMap.Walls {
		if wall.Indestructible {
			rl.DrawRectangle(int32(wall.Position.X*50), int32(wall.Position.Y*50), 50, 50, rl.DarkGray)
		} else {
			rl.DrawRectangle(int32(wall.Position.X*50), int32(wall.Position.Y*50), 50, 50, rl.Gray)
		}
	}
}

func drawGame(game Game) {
	if len(game.Players) == 0 {
		fmt.Println("No players")
		return
	}

	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)

	drawWalls(game)

	drawPlayers(game)

	drawBombs(game)

	drawPowerUps(game)

	rl.EndDrawing()
}
