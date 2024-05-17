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
	case utils.ActionUpdate:
		fmt.Println("No action was sent")
	case utils.ActionLeave:
		fmt.Println("Player disconnected")
		game.RemovePlayer(msg.ID)
	default:
		fmt.Println("Action unknown")
	}
}
