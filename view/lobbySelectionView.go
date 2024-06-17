package view

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	INPUT_BOX_POS_X  = WIDTH/2 - 100
	INPUT_BOX_POS_Y  = HEIGHT/2 - 25
	INPUT_BOX_WIDTH  = 350
	INPUT_BOX_HEIGHT = 50
)

// Raylib no tiene cajas de texto, este es un intento de simular una
func DrawLobbySelectionScreen(lobbyID string) {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Beige)

	rl.DrawText("Enter Game ID", INPUT_BOX_POS_X, INPUT_BOX_POS_Y-40, 20, rl.Maroon)
	rl.DrawRectangleLines(INPUT_BOX_POS_X-95, INPUT_BOX_POS_Y, INPUT_BOX_WIDTH, INPUT_BOX_HEIGHT, rl.DarkPurple)

	rl.DrawText(lobbyID, INPUT_BOX_POS_X-90, INPUT_BOX_POS_Y+15, 20, rl.Maroon)

	rl.EndDrawing()
}
