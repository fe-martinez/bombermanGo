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

func isMouseInButton(buttonX int32, buttonY int32) bool {
	return rl.GetMouseX() > buttonX && rl.GetMouseX() < buttonX+view.BUTTON_WIDTH && rl.GetMouseY() > buttonY && rl.GetMouseY() < buttonY+view.BUTTON_HEIGHT
}

func handleMainMenuInput() string {
	if rl.IsMouseButtonDown(0) {
		if isMouseInButton(view.JOIN_GAME_POS_X, view.JOIN_GAME_POS_Y) {
			return "join"
		} else if isMouseInButton(view.CREATE_GAME_POS_X, view.CREATE_GAME_POS_Y) {
			return "create"
		} else if isMouseInButton(view.RULES_POS_X, view.RULES_POS_Y) {
			return "rules"
		} else if isMouseInButton(view.CONTROLS_POS_X, view.CONTROLS_POS_Y) {
			return "controls"
		}
	}
	return "none"
}

func handleControlRulesInput() string {
	if rl.IsMouseButtonDown(0) {
		if isMouseInButton(view.BACK_BUTTON_X, view.BACK_BUTTON_Y) {
			return "back"
		}
	}
	return "none"
}

func handleRulesInput() string {
	if rl.IsMouseButtonDown(0) {
		if isMouseInButton(view.BACK_BUTTON_X, view.BACK_BUTTON_Y) {
			return "back"
		}
	}
	return "none"
}

func handleLobbySelectionInput() (string, string) {
	key := rl.GetCharPressed()

	if key != 0 && len(userInput) < 6 {
		userInput += string(key)
	}

	if rl.IsMouseButtonDown(0) && isMouseInButton(view.LOBBY_SEL_BACK_BUTTON_X, view.LOBBY_SEL_BACK_BUTTON_Y) {
		return "back", ""
	}

	if rl.IsKeyDown(rl.KeyEnter) {
		result := userInput
		userInput = ""
		return "join", result
	}

	return "none", userInput
}

func handleWaitingMenuInput() string {
	if rl.IsMouseButtonDown(0) {
		if isMouseInButton(view.START_GAME_POS_X, view.START_GAME_POS_Y) {
			return "start"
		}
		if isMouseInButton(view.BACK_TO_MENU_POS_X, view.BACK_TO_MENU_POS_Y) {
			return "back"
		}
	}
	return "none"
}

func handleGameOverInput() string {
	if rl.IsMouseButtonDown(0) {
		if isMouseInButton(view.RETURN_MAIN_MENU_BUTTON_X, view.RETURN_MAIN_MENU_BUTTON_Y) {
			return "return"
		}
	}
	return "none"
}
