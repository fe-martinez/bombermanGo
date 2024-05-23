package client

import (
	"bombman/view"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func handleInput() string {
	if rl.IsKeyDown(rl.KeyW) {
		return "up"
	}
	if rl.IsKeyDown(rl.KeyS) {
		return "down"
	}
	if rl.IsKeyDown(rl.KeyA) {
		return "left"
	}
	if rl.IsKeyDown(rl.KeyD) {
		return "right"
	}
	if rl.IsKeyDown(rl.KeyB) {
		return "bomb"
	}

	return "none"
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

func handleLobbySelectionInput() string {
	if rl.IsKeyDown(rl.KeyEnter) {
		return "join"
	}
	return "none"
}
