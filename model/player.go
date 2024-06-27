package model

import (
	"log"

	petname "github.com/dustinkirkland/golang-petname"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const PLAYER_LIVES = 2
const PLAYER_BOMBS = 1
const BOMB_REACH_BASE = 2
const START_DIRECTION = "down"

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
		BombReach:  BOMB_REACH_BASE,
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

func (p *Player) IsInvencible() bool {
	return p.Invencible
}

func (p *Player) LoseHealth() int8 {
	p.Lives--
	return p.Lives
}

func (p *Player) GetFirstPowerUp() *PowerUp {
	if len(p.PowerUps) > 0 {
		return &p.PowerUps[0]
	}
	return nil
}

func (p *Player) AddBomb() bool {
	firstPowerUp := p.GetFirstPowerUp()
	if firstPowerUp != nil && firstPowerUp.Name == MasBombasEnSimultaneo {
		log.Println("First PowerUp:", firstPowerUp.Name)
		if p.Bombs <= 4 {
			p.Bombs++
		}
	} else if p.Bombs == 0 {
		log.Println("No PowerUps available")
		p.Bombs++
	}
	return true
}

func (p *Player) ApplyPowerUpBenefit(powerUp PowerUpType) {
	switch powerUp {
	case Invencibilidad:
		log.Println("Invencibilidad")
		p.Invencible = true
	case MasBombasEnSimultaneo:
		log.Println("Mas bombas en simultaneo")
		p.Bombs = 5
	case AlcanceMejorado:
		log.Println("Alcance mejorado")
		p.BombReach = BOMB_REACH_MODIFED
	default:
	}
}

func (p *Player) RemovePowerUpBenefit(powerUp PowerUpType) {
	switch powerUp {
	case Invencibilidad:
		log.Println("Removiendo invencibilidad")
		p.Invencible = false
	case MasBombasEnSimultaneo:
		log.Println("Removiendo mas bombas en simultaneo")
		p.Bombs = 1
	case AlcanceMejorado:
		log.Println("Removiendo alcance mejorado")
		p.BombReach = BOMB_REACH_BASE
	default:
	}
}

func (p Player) GetPosition() Position {
	return *p.Position
}

func (p Player) GetSize() float32 {
	return 55
}

func (p Player) GetRect() rl.Rectangle {
	return rl.NewRectangle(p.Position.X*65+5, p.Position.Y*65+5, 55, 55)
}
