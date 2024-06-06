package model

import "time"

const POWERUP_EXPIRTE_TIME = 10 //segundos
const POWER_UPS_AMOUNT = 4

type PowerUpType int8

const (
	Invencibilidad PowerUpType = iota
	CaminarSobreParedes
	AlcanceMejorado
	MasBombasEnSimultaneo
)

type PowerUp struct {
	Position   Position
	name       PowerUpType
	StartTime  time.Time
	ExpireTime time.Duration
}

func NewPowerUp(position Position, name PowerUpType) PowerUp {
	return PowerUp{
		Position:   position,
		name:       name,
		StartTime:  time.Time{},
		ExpireTime: POWERUP_EXPIRTE_TIME * time.Second,
	}
}

func (p *PowerUp) SetPowerUpStartTime() {
	p.StartTime = time.Now()
}
