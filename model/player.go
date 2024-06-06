package model

import (
	petname "github.com/dustinkirkland/golang-petname"
)

const PLAYER_LIVES = 2
const PLAYER_BOMBS = 1
const START_DIRECTION = "up"

type Player struct {
	Username  string
	ID        string
	Position  *Position
	Lives     int8
	Bombs     int8
	PowerUps  []PowerUp
	Speed     Speed
	Direction string
}

type Speed struct {
	X float32
	Y float32
}

func generateRandomUsername() string {
	return petname.Generate(2, "-") // Genera un nombre con 2 palabras separadas por un guión
}

func NewPlayer(id string, position *Position) *Player {
	return &Player{
		Username:  generateRandomUsername(),
		ID:        id,
		Position:  position,
		Lives:     PLAYER_LIVES,
		Bombs:     PLAYER_BOMBS,
		PowerUps:  []PowerUp{},
		Speed:     Speed{X: 0.0, Y: 0.0},
		Direction: START_DIRECTION,
	}
}

func (p *Player) AddPowerUp(powerUp PowerUp) {
	p.PowerUps = append(p.PowerUps, powerUp)
}
