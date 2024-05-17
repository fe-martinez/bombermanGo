package server

import (
	"bombman/model"
	"bombman/utils"
	"fmt"
)

func handleMove(msg utils.ClientMessage, game *model.Game) {
	switch msg.Action {
	case utils.ActionMove:
		fmt.Println("Received: ", msg.Data)
	case utils.ActionBomb:
		fmt.Println("Received: ", msg.Action)
	case utils.ActionUpdate:
		fmt.Println("No action was sent")
	case utils.ActionLeave:
		fmt.Println("Player disconnected")
	default:
		fmt.Println("Action unknown")
	}
}
