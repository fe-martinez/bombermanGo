package client

import (
	"bombman/view"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var userInput string

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

func isMouseInRulesButton() bool {
	return rl.GetMouseX() > view.RULES_POS_X && rl.GetMouseX() < view.RULES_POS_X+view.BUTTON_WIDTH && rl.GetMouseY() > view.RULES_POS_Y && rl.GetMouseY() < view.RULES_POS_Y+view.BUTTON_HEIGHT
}

func isMouseInControlsButton() bool {
	return rl.GetMouseX() > view.CONTROLS_POS_X && rl.GetMouseX() < view.CONTROLS_POS_X+view.BUTTON_WIDTH && rl.GetMouseY() > view.CONTROLS_POS_Y && rl.GetMouseY() < view.CONTROLS_POS_Y+view.BUTTON_HEIGHT
}

func handleMainMenuInput() string {
	if rl.IsMouseButtonDown(0) {
		if isMouseInJoinButton() {
			return "join"
		} else if isMouseInCreateButton() {
			return "create"
		} else if isMouseInRulesButton() {
			return "rules"
		} else if isMouseInControlsButton() {
			return "controls"
		}
	}
	return "none"
}

func isMouseInBackButton() bool {
	return rl.GetMouseX() > view.BACK_BUTTON_X && rl.GetMouseX() < view.BACK_BUTTON_X+view.BUTTON_WIDTH && rl.GetMouseY() > view.BACK_BUTTON_Y && rl.GetMouseY() < view.BACK_BUTTON_Y+view.BUTTON_HEIGHT
}

func handleControlRulesInput() string {
	if rl.IsMouseButtonDown(0) {
		if isMouseInBackButton() {
			return "back"
		}
	}
	return "none"
}

func handleRulesInput() string {
	if rl.IsMouseButtonDown(0) {
		if isMouseInBackButton() {
			return "back"
		}
	}
	return "none"
}

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

func isMouseInGameOverReturnButton() bool {
	return rl.GetMouseX() > view.RETURN_MAIN_MENU_BUTTON_X && rl.GetMouseX() < view.RETURN_MAIN_MENU_BUTTON_X+view.BUTTON_WIDTH && rl.GetMouseY() > view.RETURN_MAIN_MENU_BUTTON_Y && rl.GetMouseY() < view.RETURN_MAIN_MENU_BUTTON_Y+view.BUTTON_HEIGHT
}

func handleGameOverInput() string {
	if rl.IsMouseButtonDown(0) {
		if isMouseInGameOverReturnButton() {
			return "return"
		}
	}
	return "none"
}
