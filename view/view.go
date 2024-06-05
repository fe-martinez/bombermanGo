package view

import (
	"bombman/model"
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const TILE_SIZE = 65

func InitWindow() {
	rl.InitWindow(TILE_SIZE*16, TILE_SIZE*10, "Bomberman Go!")
	rl.SetTargetFPS(30)
}

func WindowShouldClose() bool {
	return rl.WindowShouldClose()
}

// Después se van a dibujar diferenciados, no todos iguales
func drawPlayers(game model.Game) {
	for _, player := range game.Players {
		rl.DrawRectangle(int32(player.Position.X*TILE_SIZE), int32(player.Position.Y*TILE_SIZE), TILE_SIZE, TILE_SIZE, rl.Red)
	}
}

func drawBombs(game model.Game) {
	for _, bomb := range game.GameMap.Bombs {
		rl.DrawRectangle(int32(bomb.Position.X*TILE_SIZE), int32(bomb.Position.Y*TILE_SIZE), TILE_SIZE, TILE_SIZE, rl.Blue)
	}
}

// Después va a tener que dibujar los distintos powerups según el tipo
func drawPowerUps(game model.Game) {
	for _, powerUp := range game.GameMap.PowerUps {
		rl.DrawRectangle(int32(powerUp.Position.X*TILE_SIZE), int32(powerUp.Position.Y*TILE_SIZE), TILE_SIZE, TILE_SIZE, rl.Brown)
	}
}

func drawWalls(game model.Game) {
	for _, wall := range game.GameMap.Walls {
		if wall.Indestructible {
			rl.DrawRectangle(int32(wall.Position.X*TILE_SIZE), int32(wall.Position.Y*TILE_SIZE), TILE_SIZE, TILE_SIZE, rl.DarkGray)
		} else {
			rl.DrawRectangle(int32(wall.Position.X*TILE_SIZE), int32(wall.Position.Y*TILE_SIZE), TILE_SIZE, TILE_SIZE, rl.Gray)
		}
	}
}

func DrawGame(game model.Game) {
	if len(game.Players) == 0 {
		return
	}

	fmt.Println("Players in game:", len(game.Players))

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

const (
	INPUT_BOX_POS_X  = 250
	INPUT_BOX_POS_Y  = 450
	INPUT_BOX_WIDTH  = 350
	INPUT_BOX_HEIGHT = 50
)

// Raylib no tiene cajas de texto, este es un intento de simular una
func DrawLobbySelectionScreen(lobbyID string) {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Beige)

	rl.DrawText("Enter Game ID", INPUT_BOX_POS_X, INPUT_BOX_POS_Y-40, 20, rl.Maroon)
	rl.DrawRectangleLines(INPUT_BOX_POS_X, INPUT_BOX_POS_Y, INPUT_BOX_WIDTH, INPUT_BOX_HEIGHT, rl.DarkPurple)

	rl.DrawText(lobbyID, INPUT_BOX_POS_X+10, INPUT_BOX_POS_Y+15, 20, rl.Maroon)

	rl.EndDrawing()
}

const (
	START_TITLE_POS_X  = 350
	START_TITLE_POS_Y  = 200
	START_GAME_POS_X   = 250
	START_GAME_POS_Y   = 450
	PLAYER_START_POS_X = 100
	PLAYER_START_POS_Y = 150
	PLAYER_GAP_Y       = 30
	GAME_ID_POS_X      = 50  // Posición X para el texto del ID del juego
	GAME_ID_POS_Y      = 50  // Posición Y para el texto del ID del juego
	TEXTBOX_WIDTH      = 200 // Ancho del recuadro para el ID del juego
	TEXTBOX_HEIGHT     = 30  // Alto del recuadro para el ID del juego
)

func DrawWaitingMenu(players []string, lobbyId string) {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Beige)
	rl.DrawText("Are you ready for the game?", START_TITLE_POS_X, START_TITLE_POS_Y, 20, rl.Maroon)
	// Draw Game ID
	rl.DrawText("Game ID:", GAME_ID_POS_X, GAME_ID_POS_Y, 20, rl.Black)
	rl.DrawRectangle(GAME_ID_POS_X, GAME_ID_POS_Y+25, TEXTBOX_WIDTH, TEXTBOX_HEIGHT, rl.DarkGray)
	rl.DrawText(lobbyId, GAME_ID_POS_X+5, GAME_ID_POS_Y+30, 20, rl.Black)
	// Draw Connected players
	rl.DrawText("Connected players:", PLAYER_START_POS_X, PLAYER_START_POS_Y-30, 20, rl.Black)
	// Draw players
	for i, player := range players {
		yPos := PLAYER_START_POS_Y + int32(i)*PLAYER_GAP_Y
		rl.DrawText(player, PLAYER_START_POS_X, yPos, 20, rl.Black)
	}
	// Draw Start Game button
	rl.DrawRectangle(START_GAME_POS_X, START_GAME_POS_Y, BUTTON_WIDTH, BUTTON_HEIGHT, rl.DarkPurple)
	rl.DrawText("Start Game", START_GAME_POS_X+10, START_GAME_POS_Y+15, 20, rl.White)
	rl.EndDrawing()
}
