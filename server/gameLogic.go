package server

import (
	"bombman/model"
	"bombman/utils"
	"fmt"
)

func handlePlayerAction(msg utils.ClientMessage, game *model.Game) {
	switch msg.Action {
	case utils.ActionMove:
		direction := msg.Data.(string)  
		fmt.Println("Received: ", direction)
		movePlayer(game.Players[msg.ID], direction)
	case utils.ActionBomb:
		fmt.Println("Received: ", msg.Action)
	default:
		fmt.Println("Action unknown")
	}
}

func movePlayer(player *model.Player, direction string) {
	switch direction {
	case "up":
		player.Speed.Y = -0.1
		player.Position.Y += player.Speed.Y
	case "right":
		player.Speed.X = 0.1
		player.Position.X += player.Speed.X
	case "down":
		player.Speed.Y = 0.1
		player.Position.Y += player.Speed.Y
	case "left":
		player.Speed.X = -0.1
		player.Position.X += player.Speed.X
	
	default:
		player.Speed.X = 0
	}
}