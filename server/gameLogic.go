package server

import (
	"bombman/model"
	"bombman/utils"
	"fmt"
)

const speedIncrement = 0.1
const baseSpeed = 0

type Direction struct {
    X, Y float32
}

var directionMap = map[string]Direction{
    "up":    {0, -1},
    "right": {1, 0},
    "down":  {0, 1},
    "left":  {-1, 0},
}

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
	if dir, ok := directionMap[direction]; ok {
		player.Speed.X, player.Speed.Y = dir.X*speedIncrement, dir.Y*speedIncrement
		player.Position.X, player.Position.Y = player.Position.X+dir.X*speedIncrement, player.Position.Y+dir.Y*speedIncrement 
		player.Direction = direction
        
    } else {
        player.Speed.X, player.Speed.Y = baseSpeed, baseSpeed
    }
}