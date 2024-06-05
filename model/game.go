package model

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const MAX_PLAYERS = 4
const MAX_ROUNDS = 5

var colors = NewQueue()

type Game struct {
	State        string
	GameId       string
	Round        int8
	Players      map[string]*Player
	PlayerColors map[string]string
	GameMap      *GameMap
}

func initializeColors() {
	colors.Enqueue("Orange")
	colors.Enqueue("Green")
	colors.Enqueue("Violet")
	colors.Enqueue("Blue")
}

func NewGame(id string, GameMap *GameMap) *Game {
	initializeColors()
	return &Game{
		State:        "not-started",
		GameId:       id,
		Round:        1,
		Players:      make(map[string]*Player),
		PlayerColors: make(map[string]string),
		GameMap:      GameMap,
	}
}

func (g *Game) collidesWithWalls(position Position) bool {
	playerRect := rl.NewRectangle(position.X*65, position.Y*65, 65, 65)

	for _, wall := range g.GameMap.Walls {
		wallRect := rl.NewRectangle(wall.Position.X*65, wall.Position.Y*65, 55, 55)
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
	color, ok := colors.Dequeue()
	if ok {
		g.PlayerColors[player.ID] = color
	}
}

func (g *Game) GetPlayerColors() map[string]string {
	return g.PlayerColors
}

func (g *Game) GetPlayerColor(id string) string {
	return g.PlayerColors[id]
}

func (g *Game) RemovePlayer(playerID string) {
	delete(g.Players, playerID)
	color := g.GetPlayerColor(playerID)
	colors.Enqueue(color)
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
