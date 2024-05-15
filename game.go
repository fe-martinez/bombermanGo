package main

import (
	"fmt"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const BOMBTIME = 3

type Game struct {
	GameId  string
	Level   int8
	Players []Player
	GameMap Map
}

type Position struct {
	X float32
	Y float32
}

type Player struct {
	id string
	//position Position no anda el juego si uso esto xd lol
	X        float32
	Y        float32
	lives    int8
	maxBombs int8
	bombs    []Bomb
	powerUps []string
}

func initGame(gameMap Map, gameId string) Game {
	game := Game{
		GameId:  gameId,
		Level:   1,
		Players: make([]Player, 0),
		GameMap: gameMap,
	}

	return game
}

func createPlayer(game *Game, playerId string) {
	player := Player{}
	for {
		player.id = playerId
		player.X = float32(rand.Intn(game.GameMap.size))
		player.Y = float32(rand.Intn(game.GameMap.size))
		player.lives = 6
		if !checkCollision(player.X, player.Y, *game, playerId) {
			break
		}
	}
	game.Players = append(game.Players, player)
}

func disconnectPlayer(game *Game, playerId string) {
	playerPosition := getPlayerPositionInList(game, playerId)
	if playerPosition == -1 {
		return
	}
	game.Players = append(game.Players[:playerPosition], game.Players[playerPosition+1:]...)
}

func getPlayerPositionInList(game *Game, id string) int {
	for i := range game.Players {
		if game.Players[i].id == id {
			return i
		}
	}
	return -1
}

func getPlayerPosition(clientID string, game Game) Position {
	position := Position{}
	for i := range game.Players {
		if game.Players[i].id == clientID {
			position.X = game.Players[i].X
			position.Y = game.Players[i].Y
		}
	}
	return position
}

func placeBomb(position Position, playerId string) {
	// bomb := Bomb{
	// 	X:     position.X,
	// 	Y:     position.Y,
	// 	Timer: BOMBTIME,
	// }
	fmt.Println("Placing bomb... TO-DO!")
}

func movePlayer(game *Game, direction string, playerId string) {
	playerPosition := getPlayerPositionInList(game, playerId)
	if playerPosition == -1 {
		return
	}

	switch direction {
	case "up":
		if !checkCollision(game.Players[playerPosition].X, game.Players[playerPosition].Y-0.05, *game, playerId) {
			game.Players[playerPosition].Y = game.Players[playerPosition].Y - 0.05
		}
	case "down":
		if !checkCollision(game.Players[playerPosition].X, game.Players[playerPosition].Y+0.05, *game, playerId) {
			game.Players[playerPosition].Y = game.Players[playerPosition].Y + 0.05
		}
	case "left":
		if !checkCollision(game.Players[playerPosition].X-0.05, game.Players[playerPosition].Y, *game, playerId) {
			game.Players[playerPosition].X = game.Players[playerPosition].X - 0.05
		}
	case "right":
		if !checkCollision(game.Players[playerPosition].X+0.05, game.Players[playerPosition].Y, *game, playerId) {
			game.Players[playerPosition].X = game.Players[playerPosition].X + 0.05
		}
	}
}

func checkCollision(x float32, y float32, game Game, playerID string) bool {
	// playerPosition := getPlayerPositionInList(&game, playerID)
	// if playerPosition == -1 {
	// 	return false
	// }

	playerRect := rl.NewRectangle(x*50, y*50, 50, 50)

	for _, wall := range game.GameMap.Walls {
		wallRect := rl.NewRectangle(wall.X*50, wall.Y*50, 45, 45)
		if rl.CheckCollisionRecs(playerRect, wallRect) {
			return true
		}
	}
	//Ni idea, no funciona
	// for _, player := range game.Players {
	// 	otherPlayerRect := rl.NewRectangle(player.X*50, player.Y*50, 50, 50)
	// 	if rl.CheckCollisionRecs(playerRect, otherPlayerRect) && player.id != game.Players[playerPosition].id {
	// 		return true
	// 	}
	// }

	return false
}
