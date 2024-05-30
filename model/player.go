package model

import petname "github.com/dustinkirkland/golang-petname"

const PLAYER_LIVES = 6
const PLAYER_BOMBS = 1

type Player struct {
	Username string
	ID       string
	Position *Position
	Lives    int8
	Bombs    int8
	PowerUps []PowerUp
}

func generateRandomUsername() string {
	return petname.Generate(2, "-") // Genera un nombre con 2 palabras separadas por un guión
}

func NewPlayer(id string, position *Position) *Player {
	return &Player{
		Username: generateRandomUsername(),
		ID:       id,
		Position: position,
		Lives:    PLAYER_LIVES,
		Bombs:    PLAYER_BOMBS,
		PowerUps: []PowerUp{},
	}
}
