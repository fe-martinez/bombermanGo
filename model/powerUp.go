package model

import (
	"log"
	"time"
)

const POWERUP_EXPIRE_TIME = 10 * time.Second //segundos
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
	ExpireTime time.Time
	Duration   time.Duration
}

func NewPowerUp(position Position, name PowerUpType) PowerUp {
	return PowerUp{
		Position:   position,
		name:       name,
		Duration: POWERUP_EXPIRE_TIME,
	}
}

func (p *PowerUp) SetPowerUpStartTime() {
	p.StartTime = time.Now()
	p.ExpireTime = p.StartTime.Add(p.Duration)
	log.Println("PowerUp Start Time: ", p.StartTime)
	log.Println("PowerUp Expiry Time: ", p.ExpireTime)
}
