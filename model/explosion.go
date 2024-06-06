package model

import (
	"time"
)

type Explosion struct {
	Position      Position
	AffectedTiles []Position
	ExplosionTime time.Time
}

func NewExplosion(position Position, radius int, game Game) *Explosion {
	return &Explosion{
		Position:      position,
		AffectedTiles: getAffectedTiles(position, radius, game),
		ExplosionTime: time.Now(),
	}
}

func getAffectedTiles(position Position, radius int, game Game) []Position {
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

	for i, tile := range affectedTiles {
		if game.GameMap.isBreakableWall(tile) {
			game.GameMap.RemoveWall(tile)
			affectedTiles[i] = Position{X: tile.X, Y: tile.Y}
		}
	}

	return affectedTiles
}
