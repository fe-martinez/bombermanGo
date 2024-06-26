package model

import (
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const EXPLODE_TIME = 3

type Bomb struct {
	Position    Position
	Alcance     int
	PlantedTime time.Time
	ExplodeTime time.Duration
	Owner       *Player
}

func (b *Bomb) IsOwner(playerID string) bool {
	return b.Owner.ID == playerID
}

func NewBomb(x, y float32, alcance int, owner Player) *Bomb {
	return &Bomb{
		Position:    Position{X: x, Y: y},
		Alcance:     alcance,
		PlantedTime: time.Now(),
		ExplodeTime: EXPLODE_TIME * time.Second,
		Owner:       &owner,
	}
}

func (b Bomb) GetPosition() Position {
	return b.Position
}

func (b Bomb) GetSize() float32 {
	return 65
}

func (b Bomb) GetRect() rl.Rectangle {
	return rl.NewRectangle(b.Position.X*65, b.Position.Y*65, 65, 65)
}
