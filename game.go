package main

import (
	"fmt"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const BOMBTIME = 3.00

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

type PowerUp struct {
	name  string
	timer float32
}

type Bomb struct {
	X       float32
	Y       float32
	Timer   float32
	Alcance int8
	Power   int8
}

type Player struct {
	id string
	//position Position no anda el juego si uso esto xd lol
	X        float32
	Y        float32
	lives    int8
	maxBombs int8
	bombs    []Bomb
	powerUps []PowerUp
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
		player.maxBombs = 1
		player.bombs = make([]Bomb, 0)
		player.powerUps = make([]PowerUp, 0)
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

func getPlayerPowerUps(game *Game, playerId string) []PowerUp {
	i := getPlayerPositionInList(game, playerId)
	if i == -1 {
		return nil
	}
	powerUps := game.Players[i].powerUps
	return powerUps
}

func powerUpIs(PowerUpName string, powerUp []PowerUp) bool {
	for i := range powerUp {
		if powerUp[i].name == PowerUpName {
			return true
		}
	}
	return false
}

func insertBombInPlayer(game *Game, playerId string, bomb Bomb) {

	i := getPlayerPositionInList(game, playerId)
	if i == -1 {
		return
	}
	game.Players[i].bombs = append(game.Players[i].bombs, bomb)
}

func getPlayer(playerId string, game Game) Player {
	for i := range game.Players {
		if game.Players[i].id == playerId {
			return game.Players[i]
		}
	}
	return Player{}
}

// Esto hay que refactorizarlo
func placeBomb(position Position, playerId string, game *Game) {
	powerUps := getPlayerPowerUps(game, playerId)
	if powerUps == nil {
		return
	}

	player := getPlayer(playerId, *game)
	if int8(len(player.bombs)) >= player.maxBombs {
		return
	}

	var alcance = 1
	var power = 1
	if powerUpIs("alcance_mejorado", powerUps) {
		alcance = 2
	}
	if powerUpIs("potencia_mejorada", powerUps) {
		power = 2
	}
	bomb := Bomb{
		X:       position.X,
		Y:       position.Y,
		Timer:   BOMBTIME,
		Alcance: int8(alcance),
		Power:   int8(power),
	}

	insertBombInPlayer(game, playerId, bomb)
	fmt.Println("Bomb placed")
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
