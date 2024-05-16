package main

import (
	"bytes"
	"encoding/gob"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const MAX_PLAYERS = 4

type Game struct {
	GameId  string
	Level   int8
	Players map[string]*Player
	GameMap *GameMap
}

func NewGame(id string, GameMap *GameMap) *Game {
	return &Game{
		GameId:  id,
		Level:   1,
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

func (g *Game) generateValidPosition(mapSize int) *Position {
	var ValidPosition = getRandomPosition(mapSize)
	for g.collidesWithWalls(*ValidPosition) {
		ValidPosition = getRandomPosition(mapSize)
	}
	return ValidPosition
}

func (g *Game) isFull() bool {
	return len(g.Players) == MAX_PLAYERS
}

func (g *Game) addPlayer(player *Player) {
	g.Players[player.ID] = player
}

func (g *Game) removePlayer(playerID string) {
	delete(g.Players, playerID)
}

func encodeGame(game Game) ([]byte, error) {
	buf := new(bytes.Buffer)

	enc := gob.NewEncoder(buf)

	err := enc.Encode(game)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func decodeGame(encodedGame []byte) (*Game, error) {
	buf := bytes.NewBuffer(encodedGame)
	
	dec := gob.NewDecoder(buf)
	
	var game Game
	
	err := dec.Decode(&game)
	if err != nil {
		return nil, err
	}
	
	return &game, nil
}
