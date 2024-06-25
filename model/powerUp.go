package model

import (
	"time"
)

const POWERUP_EXPIRE_TIME = 10 * time.Second //segundos
const POWER_UPS_AMOUNT = 3

type PowerUpType int8

const (
	Invencibilidad PowerUpType = iota
	AlcanceMejorado
	MasBombasEnSimultaneo
)

type PowerUp struct {
	Position   Position
	Name       PowerUpType
	StartTime  time.Time
	ExpireTime time.Time
	Duration   time.Duration
}

func NewPowerUp(position Position, name PowerUpType) PowerUp {
	return PowerUp{
		Position: position,
		Name:     name,
		Duration: POWERUP_EXPIRE_TIME,
	}
}

func (p *PowerUp) SetPowerUpStartTime() {
	p.StartTime = time.Now()
	p.ExpireTime = p.StartTime.Add(p.Duration)
}
