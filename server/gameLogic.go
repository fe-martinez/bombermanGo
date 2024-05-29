package server

import (
	"bombman/model"
	"bombman/utils"
	"fmt"
)

func handlePlayerAction(msg utils.ClientMessage, game *model.Game) {
	switch msg.Action {
	case utils.ActionMove:
		fmt.Println("Received: ", msg.Data)
	case utils.ActionBomb:
		fmt.Println("Received: ", msg.Action)
	default:
		fmt.Println("Action unknown")
	}
}
