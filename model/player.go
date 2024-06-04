package model

const PLAYER_LIVES = 6
const PLAYER_BOMBS = 1
const START_DIRECTION = "up"

type Player struct {
	ID       string
	Position *Position
	Lives    int8
	Bombs    int8
	PowerUps []PowerUp
	Speed 	 Speed
	Direction string
}

type Speed struct {
	X 		float32
	Y		float32
}

func NewPlayer(id string, position *Position) *Player {
	return &Player{
		ID:        id,
		Position:  position,
		Lives:     PLAYER_LIVES,
		Bombs:     PLAYER_BOMBS,
		PowerUps:  []PowerUp{},
		Speed:	   Speed{X: 0.0, Y: 0.0},
		Direction: START_DIRECTION,
	}
}

