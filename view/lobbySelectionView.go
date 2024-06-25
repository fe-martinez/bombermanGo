package view

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	INPUT_BOX_POS_X         = (WIDTH - INPUT_BOX_WIDTH) / 2
	INPUT_BOX_POS_Y         = (HEIGHT - INPUT_BOX_HEIGHT) / 2
	LOBBY_SEL_BACK_BUTTON_X = (WIDTH - BUTTON_WIDTH) / 2
	LOBBY_SEL_BACK_BUTTON_Y = 450
	INPUT_BOX_WIDTH         = 350
	INPUT_BOX_HEIGHT        = 50
)

// Raylib no tiene cajas de texto, este es un intento de simular una
func DrawLobbySelectionScreen(lobbyID string) {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Beige)

	enterGameIDText := "Enter Game ID"
	textWidht := rl.MeasureText(enterGameIDText, 20)
	textPosX := (WIDTH - textWidht) / 2

	rl.DrawText("Enter Game ID", textPosX, INPUT_BOX_POS_Y-40, 20, rl.Maroon)
	rl.DrawRectangleLines(INPUT_BOX_POS_X, INPUT_BOX_POS_Y, INPUT_BOX_WIDTH, INPUT_BOX_HEIGHT, rl.DarkPurple)

	rl.DrawText(lobbyID, INPUT_BOX_POS_X+10, INPUT_BOX_POS_Y+15, 20, rl.Maroon)

	rl.DrawRectangle(LOBBY_SEL_BACK_BUTTON_X, LOBBY_SEL_BACK_BUTTON_Y, BUTTON_WIDTH, BUTTON_HEIGHT, rl.DarkPurple)
	rl.DrawText("Back", LOBBY_SEL_BACK_BUTTON_X+10, LOBBY_SEL_BACK_BUTTON_Y+15, 20, rl.White)

	rl.EndDrawing()
}
