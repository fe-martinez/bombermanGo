package model

import "time"

type Bomb struct {
	Position    Position
	Alcance     int8
	PlantedTime time.Time
	ExplodeTime time.Duration
	Owner       *Player
}

func NewBomb(x, y float32, alcance int8, owner Player) *Bomb {
	return &Bomb{
		Position:    Position{X: x, Y: y},
		Alcance:     alcance,
		PlantedTime: time.Now(),
		ExplodeTime: 3 * time.Second,
		Owner:       &owner,
	}
}
