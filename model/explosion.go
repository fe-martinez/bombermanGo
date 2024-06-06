package model

import "time"

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

	for i := 1; i <= radius; i++ {
		// Check left side
		if !game.GameMap.isUnbreakableWall(Position{X: position.X - float32(i), Y: position.Y}) {
			affectedTiles = append(affectedTiles, Position{X: position.X - float32(i), Y: position.Y})
		} else {
			break
		}
		// Check right side
		if !game.GameMap.isUnbreakableWall(Position{X: position.X + float32(i), Y: position.Y}) {
			affectedTiles = append(affectedTiles, Position{X: position.X + float32(i), Y: position.Y})
		} else {
			break
		}
		// Check top side
		if !game.GameMap.isUnbreakableWall(Position{X: position.X, Y: position.Y - float32(i)}) {
			affectedTiles = append(affectedTiles, Position{X: position.X, Y: position.Y - float32(i)})
		} else {
			break
		}
		// Check bottom side
		if !game.GameMap.isUnbreakableWall(Position{X: position.X, Y: position.Y + float32(i)}) {
			affectedTiles = append(affectedTiles, Position{X: position.X, Y: position.Y + float32(i)})
		} else {
			break
		}
	}

	return affectedTiles
}
