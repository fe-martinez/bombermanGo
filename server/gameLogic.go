package server

import (
	"bombman/model"
	"bombman/utils"
	"fmt"
)

const SPEED_INCREMENT = 0.1
const BASE_SPEED = 0

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
		movePlayer(game.Players[msg.ID], direction)
	default:
		fmt.Println("Action unknown")
	}
}

func movePlayer(player *model.Player, direction string) {
	if dir, ok := directionMap[direction]; ok {
		player.Speed.X, player.Speed.Y = dir.X*SPEED_INCREMENT, dir.Y*SPEED_INCREMENT
		player.Position.X, player.Position.Y = player.Position.X+dir.X*SPEED_INCREMENT, player.Position.Y+dir.Y*SPEED_INCREMENT 
		player.Direction = direction
        
    } else {
        player.Speed.X, player.Speed.Y = BASE_SPEED, BASE_SPEED
    }
}