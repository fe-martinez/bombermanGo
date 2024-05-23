package view

import (
	"bombman/model"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func InitWindow() {
	rl.InitWindow(800, 800, "Bomberman Go!")
	rl.SetTargetFPS(30)
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
func drawPlayers(game model.Game) {
	for _, player := range game.Players {
		rl.DrawRectangle(int32(player.Position.X*50), int32(player.Position.Y*50), 50, 50, rl.Red)
	}
}

func drawBombs(game model.Game) {
	for _, bomb := range game.GameMap.Bombs {
		rl.DrawRectangle(int32(bomb.Position.X*50), int32(bomb.Position.Y*50), 50, 50, rl.Blue)
	}
}

// Después va a tener que dibujar los distintos powerups según el tipo
func drawPowerUps(game model.Game) {
	for _, powerUp := range game.GameMap.PowerUps {
		rl.DrawRectangle(int32(powerUp.Position.X*50), int32(powerUp.Position.Y*50), 50, 50, rl.Brown)
	}
}

func drawWalls(game model.Game) {
	for _, wall := range game.GameMap.Walls {
		if wall.Indestructible {
			rl.DrawRectangle(int32(wall.Position.X*50), int32(wall.Position.Y*50), 50, 50, rl.DarkGray)
		} else {
			rl.DrawRectangle(int32(wall.Position.X*50), int32(wall.Position.Y*50), 50, 50, rl.Gray)
		}
	}
}

func DrawGame(game model.Game) {
	if len(game.Players) == 0 {
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

const LOBBY_TITLE_POS_X = 350
const LOBBY_TITLE_POS_Y = 200

const CREATE_GAME_POS_X = 250
const CREATE_GAME_POS_Y = 350

const JOIN_GAME_POS_X = 250
const JOIN_GAME_POS_Y = 450

const BUTTON_WIDTH = 350
const BUTTON_HEIGHT = 50

func DrawMainMenuScreen() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Beige)
	rl.DrawText("Lobby screen", LOBBY_TITLE_POS_X, LOBBY_TITLE_POS_Y, 20, rl.Maroon)
	rl.DrawRectangle(CREATE_GAME_POS_X, CREATE_GAME_POS_Y, BUTTON_WIDTH, BUTTON_HEIGHT, rl.DarkPurple)
	rl.DrawText("Create game", CREATE_GAME_POS_X+10, CREATE_GAME_POS_Y+15, 20, rl.White)
	rl.DrawRectangle(JOIN_GAME_POS_X, JOIN_GAME_POS_Y, BUTTON_WIDTH, BUTTON_HEIGHT, rl.DarkPurple)
	rl.DrawText("Join game", JOIN_GAME_POS_X+10, JOIN_GAME_POS_Y+15, 20, rl.White)
	rl.EndDrawing()
}

const MAX_INPUT_CHARS = 4

var InputChars [MAX_INPUT_CHARS + 1]byte
var letterCount = 0

const (
	INPUT_BOX_POS_X  = 250
	INPUT_BOX_POS_Y  = 450
	INPUT_BOX_WIDTH  = 350
	INPUT_BOX_HEIGHT = 50
)

// Raylib no tiene cajas de texto, este es un intento de simular una
func DrawLobbySelectionScreen() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Beige)

	rl.DrawText("Enter Game ID", INPUT_BOX_POS_X, INPUT_BOX_POS_Y-40, 20, rl.Maroon)
	rl.DrawRectangleLines(INPUT_BOX_POS_X, INPUT_BOX_POS_Y, INPUT_BOX_WIDTH, INPUT_BOX_HEIGHT, rl.DarkPurple)

	key := rl.GetKeyPressed()
	if key >= 32 && key <= 125 && letterCount < MAX_INPUT_CHARS {
		InputChars[letterCount] = byte(key)
		letterCount++
	} else if key == 8 && letterCount > 0 {
		letterCount--
		InputChars[letterCount] = 0
	}

	rl.DrawText(string(InputChars[:letterCount]), INPUT_BOX_POS_X+5, INPUT_BOX_POS_Y+8, 20, rl.Maroon)
	rl.EndDrawing()
}
