package view

import rl "github.com/gen2brain/raylib-go/raylib"

const LOBBY_TITLE_POS_X = WIDTH/2 - 70
const LOBBY_TITLE_POS_Y = 50

const CREATE_GAME_POS_X = (WIDTH - BUTTON_WIDTH) / 2
const CREATE_GAME_POS_Y = 200

const JOIN_GAME_POS_X = (WIDTH - BUTTON_WIDTH) / 2
const JOIN_GAME_POS_Y = 300

const RULES_POS_X = (WIDTH - BUTTON_WIDTH) / 2
const RULES_POS_Y = 400

const CONTROLS_POS_X = (WIDTH - BUTTON_WIDTH) / 2
const CONTROLS_POS_Y = 500

const BUTTON_WIDTH = 350
const BUTTON_HEIGHT = 50

func DrawMainMenuScreen() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Beige)

	title := "Lobby screen"
	titleFontSize := 40
	titleWidth := rl.MeasureText(title, int32(titleFontSize))

	rl.DrawText(title, (WIDTH-titleWidth)/2, LOBBY_TITLE_POS_Y, int32(titleFontSize), rl.Maroon)

	rl.DrawRectangle(CREATE_GAME_POS_X, CREATE_GAME_POS_Y, BUTTON_WIDTH, BUTTON_HEIGHT, rl.DarkPurple)
	rl.DrawText("Create game", CREATE_GAME_POS_X+10, CREATE_GAME_POS_Y+15, 20, rl.White)
	rl.DrawRectangle(CREATE_GAME_POS_X, JOIN_GAME_POS_Y, BUTTON_WIDTH, BUTTON_HEIGHT, rl.DarkPurple)
	rl.DrawText("Join game", CREATE_GAME_POS_X+10, JOIN_GAME_POS_Y+15, 20, rl.White)
	rl.DrawRectangle(RULES_POS_X, RULES_POS_Y, BUTTON_WIDTH, BUTTON_HEIGHT, rl.DarkPurple)
	rl.DrawText("Rules", RULES_POS_X+10, RULES_POS_Y+15, 20, rl.White)
	rl.DrawRectangle(CONTROLS_POS_X, CONTROLS_POS_Y, BUTTON_WIDTH, BUTTON_HEIGHT, rl.DarkPurple)
	rl.DrawText("Controls", CONTROLS_POS_X+10, CONTROLS_POS_Y+15, 20, rl.White)

	rl.EndDrawing()
}

const INSTRUCTIONS_TITLE_Y = 50

const BOX_X = 100
const BOX_Y = 150
const BOX_WIDTH = WIDTH - (BOX_X * 2)
const BOX_HEIGHT = 400

const BACK_BUTTON_X = (WIDTH - BUTTON_WIDTH) / 2
const BACK_BUTTON_Y = HEIGHT - 100

func DrawGameRules() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Beige)

	title := "Game instructions"
	titleFontSize := 20
	titleWidth := rl.MeasureText(title, int32(titleFontSize))
	rl.DrawText(title, (WIDTH-titleWidth)/2, LOBBY_TITLE_POS_Y, int32(titleFontSize), rl.Maroon)

	rl.DrawRectangleLines(BOX_X, BOX_Y, BOX_WIDTH, BOX_HEIGHT, rl.Black)

	rl.DrawRectangle(BACK_BUTTON_X, BACK_BUTTON_Y, BUTTON_WIDTH, BUTTON_HEIGHT, rl.DarkPurple)
	rl.DrawText("Back", BACK_BUTTON_X+10, BACK_BUTTON_Y+15, 20, rl.White)
	rl.EndDrawing()
}

func DrawControlsRules() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Beige)

	title := "Game Controls"
	titleFontSize := 20
	titleWidth := rl.MeasureText(title, int32(titleFontSize))
	rl.DrawText(title, (WIDTH-titleWidth)/2, LOBBY_TITLE_POS_Y, int32(titleFontSize), rl.Maroon)

	rl.DrawRectangleLines(BOX_X, BOX_Y, BOX_WIDTH, BOX_HEIGHT, rl.Black)

	rl.DrawRectangle(BACK_BUTTON_X, BACK_BUTTON_Y, BUTTON_WIDTH, BUTTON_HEIGHT, rl.DarkPurple)
	rl.DrawText("Back", BACK_BUTTON_X+10, BACK_BUTTON_Y+15, 20, rl.White)
	rl.EndDrawing()
}
