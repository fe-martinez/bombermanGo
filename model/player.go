package model

const PLAYER_LIVES = 6
const PLAYER_BOMBS = 1

type Player struct {
	ID       string
	Position *Position
	Lives    int8
	Bombs    int8
	PowerUps []PowerUp
}

func NewPlayer(id string, position *Position) *Player {
	return &Player{
		ID:       id,
		Position: position,
		Lives:    PLAYER_LIVES,
		Bombs:    PLAYER_BOMBS,
		PowerUps: []PowerUp{},
	}
}
