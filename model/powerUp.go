package model

const POWERUP_DURATION = 10 //segundos

type PowerUpType int8

const (
	Invencibilidad PowerUpType = iota
	CaminarSobreParedes
	AlcanceMejorado
	MasBombasEnSimultaneo
)

type PowerUp struct {
	Position Position
	name     PowerUpType
	duration float32
}

func NewPowerUp(position Position, name PowerUpType) PowerUp {
	return PowerUp{
		Position: position,
		name:     name,
		duration: POWERUP_DURATION,
	}
}
