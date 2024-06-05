package model

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const MAX_PLAYERS = 4
const MAX_ROUNDS = 5

type Game struct {
	State   string
	GameId  string
	Round   int8
	Players map[string]*Player
	GameMap *GameMap
}

func NewGame(id string, GameMap *GameMap) *Game {
	return &Game{
		State:   "not-started",
		GameId:  id,
		Round:   1,
		Players: make(map[string]*Player),
		GameMap: GameMap,
	}
}

func (g *Game) collidesWithWalls(position Position) bool {
	playerRect := rl.NewRectangle(position.X*50, position.Y*50, 50, 50)

	for _, wall := range g.GameMap.Walls {
		wallRect := rl.NewRectangle(wall.Position.X*50, wall.Position.Y*50, 45, 45)
		if rl.CheckCollisionRecs(playerRect, wallRect) {
			return true
		}
	}
	return false
}

func (g *Game) GenerateValidPosition(mapSize int) *Position {
	var ValidPosition = getRandomPosition(mapSize)
	for g.collidesWithWalls(*ValidPosition) {
		ValidPosition = getRandomPosition(mapSize)
	}
	return ValidPosition
}

func (g *Game) IsFull() bool {
	return len(g.Players) == MAX_PLAYERS
}

func (g *Game) AddPlayer(player *Player) {
	g.Players[player.ID] = player
}

func (g *Game) RemovePlayer(playerID string) {
	delete(g.Players, playerID)
}

func (g *Game) Start() {
	g.State = "started"
}

func (g *Game) passRound() {
	if g.Round < MAX_ROUNDS {
		g.Round++
	} else {
		g.State = "finished"
	}
}
