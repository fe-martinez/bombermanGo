package client

import (
	"bombman/view"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func handleInput() []string {
	var inputs []string

	if rl.IsKeyDown(rl.KeyW) {
		inputs = append(inputs, "up")
	}
	if rl.IsKeyDown(rl.KeyS) {
		inputs = append(inputs, "down")
	}
	if rl.IsKeyDown(rl.KeyA) {
		inputs = append(inputs, "left")
	}
	if rl.IsKeyDown(rl.KeyD) {
		inputs = append(inputs, "right")
	}
	if rl.IsKeyDown(rl.KeyB) {
		inputs = append(inputs, "bomb")
	}

	return inputs
}

func isMouseInJoinButton() bool {
	return rl.GetMouseX() > view.JOIN_GAME_POS_X && rl.GetMouseX() < view.JOIN_GAME_POS_X+view.BUTTON_WIDTH && rl.GetMouseY() > view.JOIN_GAME_POS_Y && rl.GetMouseY() < view.JOIN_GAME_POS_Y+view.BUTTON_HEIGHT
}

func isMouseInCreateButton() bool {
	return rl.GetMouseX() > view.CREATE_GAME_POS_X && rl.GetMouseX() < view.CREATE_GAME_POS_X+view.BUTTON_WIDTH && rl.GetMouseY() > view.CREATE_GAME_POS_Y && rl.GetMouseY() < view.CREATE_GAME_POS_Y+view.BUTTON_HEIGHT
}

func handleMainMenuInput() string {
	if rl.IsMouseButtonDown(0) {
		if isMouseInJoinButton() {
			return "join"
		} else if isMouseInCreateButton() {
			return "create"
		}
	}
	return "none"
}

var userInput string

func handleLobbySelectionInput() (string, string) {
	key := rl.GetCharPressed()

	if key != 0 {
		userInput += string(key)
	}

	if rl.IsKeyDown(rl.KeyEnter) {
		result := userInput
		userInput = ""
		return "join", result
	}

	return "none", userInput
}

func isMouseInStartGameButton() bool {
	return rl.GetMouseX() > view.START_GAME_POS_X && rl.GetMouseX() < view.START_GAME_POS_X+view.BUTTON_WIDTH && rl.GetMouseY() > view.START_GAME_POS_Y && rl.GetMouseY() < view.START_GAME_POS_Y+view.BUTTON_HEIGHT
}

func handleWaitingMenuInput() string {
	if rl.IsMouseButtonDown(0) {
		if isMouseInStartGameButton() {
			return "start"
		}
	}
	return "none"
}
