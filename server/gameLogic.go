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
	if game.Players[msg.ID].Lives == 0 {
		return
	}

	switch msg.Action {
	case utils.ActionMove:
		direction := msg.Data.([]string)
		movePlayer(game.Players[msg.ID], direction, game)
	case utils.ActionBomb:
		game.PutBomb(game.Players[msg.ID])
	default:
		fmt.Println("Action unknown")
	}
}

func movePlayer(player *model.Player, directions []string, game *model.Game) {
	for _, direction := range directions {
		if dir, ok := directionMap[direction]; ok {
			newX := player.Position.X + dir.X*SPEED_INCREMENT
			newY := player.Position.Y + dir.Y*SPEED_INCREMENT
			game.MovePlayer(player, newX, newY)
			player.Direction = directions[0]
		} else {
			player.Speed.X, player.Speed.Y = BASE_SPEED, BASE_SPEED
		}
	}
}
