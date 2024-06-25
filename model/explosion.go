package model

import (
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Explosion struct {
	Position        Position
	AffectedTiles   []Position
	AffectedPlayers []string
	ExplosionTime   time.Time
}

func NewExplosion(position Position, radius int, game Game) *Explosion {
	return &Explosion{
		Position:      position,
		AffectedTiles: getAffectedTiles(position, radius, game),
		ExplosionTime: time.Now(),
	}
}

func lookForAffectedTiles(game Game, position Position, radius int) []Position {
	var affectedTiles []Position
	affectedTiles = append(affectedTiles, position)

	directions := []Position{
		{X: -1, Y: 0},
		{X: 1, Y: 0},
		{X: 0, Y: -1},
		{X: 0, Y: 1},
	}

	for _, dir := range directions {
		for i := 1; i <= radius; i++ {
			newPos := Position{X: position.X + float32(i)*dir.X, Y: position.Y + float32(i)*dir.Y}
			if !game.GameMap.isUnbreakableWall(newPos) {
				affectedTiles = append(affectedTiles, newPos)
			} else {
				break
			}
		}
	}
	return affectedTiles
}

func removeWalls(affectedTiles []Position, game Game) {
	for i, tile := range affectedTiles {
		if !game.GameMap.isUnbreakableWall(tile) {
			game.GameMap.RemoveWall(tile)
			affectedTiles[i] = Position{X: tile.X, Y: tile.Y}
		}
	}
}

func getAffectedTiles(position Position, radius int, game Game) []Position {
	affectedTiles := lookForAffectedTiles(game, position, radius)
	removeWalls(affectedTiles, game)
	return affectedTiles
}

func (e *Explosion) IsExpired() bool {
	return time.Since(e.ExplosionTime) > 500*time.Millisecond
}

func (e *Explosion) IsTileInRange(position Position) bool {
	playerPos := rl.NewRectangle(position.X*65+5, position.Y*65+5, 50, 50)
	for _, tile := range e.AffectedTiles {
		explosionRect := rl.NewRectangle(tile.X*65, tile.Y*65, 65, 65)
		if rl.CheckCollisionRecs(playerPos, explosionRect) {
			return true
		}
	}
	return false
}

func (e *Explosion) AddAffectedPlayer(playerID string) {
	e.AffectedPlayers = append(e.AffectedPlayers, playerID)
}

func (e *Explosion) IsPlayerAlreadyAffected(playerID string) bool {
	for _, pID := range e.AffectedPlayers {
		if pID == playerID {
			return true
		}
	}
	return false
}
