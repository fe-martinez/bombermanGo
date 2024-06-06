package model

import "time"

const EXPLODE_TIME = 3

type Bomb struct {
	Position    Position
	Alcance     int8
	PlantedTime time.Time
	ExplodeTime time.Duration
	Owner       *Player
}

func (b *Bomb) IsOwner(playerID string) bool {
	return b.Owner.ID == playerID
}

func NewBomb(x, y float32, alcance int8, owner Player) *Bomb {
	return &Bomb{
		Position:    Position{X: x, Y: y},
		Alcance:     alcance,
		PlantedTime: time.Now(),
		ExplodeTime: EXPLODE_TIME * time.Second,
		Owner:       &owner,
	}
}
