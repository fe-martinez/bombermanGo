package model

import (
	"log"

	petname "github.com/dustinkirkland/golang-petname"
)

const PLAYER_LIVES = 2
const PLAYER_BOMBS = 1
const PLAYER_BOMB_REACH_BASE = 2
const START_DIRECTION = "up"

type Player struct {
	Username   string
	ID         string
	Position   *Position
	Lives      int8
	Invencible bool
	Bombs      int8
	PowerUps   []PowerUp
	Speed      Speed
	Direction  string
	BombReach  int
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
		Username:   generateRandomUsername(),
		ID:         id,
		Position:   position,
		Lives:      PLAYER_LIVES,
		Invencible: false,
		Bombs:      PLAYER_BOMBS,
		PowerUps:   []PowerUp{},
		Speed:      Speed{X: 0.0, Y: 0.0},
		Direction:  START_DIRECTION,
		BombReach:	PLAYER_BOMB_REACH_BASE,
	}
}

func (p *Player) CanPlantBomb() bool {
	return p.Bombs > 0
}

func (p *Player) AddPowerUp(powerUp PowerUp) {
	log.Println("Adding power up to player", powerUp)
	p.PowerUps = append(p.PowerUps, powerUp)
}

func (p *Player) RemovePowerUp(powerUp PowerUp) {
	for i, power := range p.PowerUps {
		if powerUp == power {
			p.PowerUps = append(p.PowerUps[:i], p.PowerUps[i+1:]...)
		}
	}
}
