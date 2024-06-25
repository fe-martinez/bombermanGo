package view

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	START_TITLE_POS_X  = 400
	START_TITLE_POS_Y  = 200
	START_GAME_POS_X   = 370
	START_GAME_POS_Y   = 350
	BACK_TO_MENU_POS_X = 370
	BACK_TO_MENU_POS_Y = 450
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

	rl.DrawRectangle(BACK_TO_MENU_POS_X, BACK_TO_MENU_POS_Y, BUTTON_WIDTH, BUTTON_HEIGHT, rl.DarkPurple)
	rl.DrawText("Back to menu", BACK_TO_MENU_POS_X+10, BACK_TO_MENU_POS_Y+15, 20, rl.White)
	rl.EndDrawing()
}
