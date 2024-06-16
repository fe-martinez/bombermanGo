package model

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

type GameMap struct {
	Walls      []Wall
	PowerUps   []PowerUp
	Bombs      []Bomb
	Explosions []Explosion
	RowSize    int
	ColumnSize int
}

type Wall struct {
	Position       *Position
	Indestructible bool
}

func CreateMap(filepath string) (*GameMap, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var walls []Wall
	row := 0
	var column = 0

	for scanner.Scan() {
		line := scanner.Text()
		for col, char := range line {
			x, y := float32(col), float32(row)
			switch char {
			case 'W':
				walls = append(walls, Wall{
					Position:       &Position{x, y},
					Indestructible: true,
				})
			case 'D':
				walls = append(walls, Wall{
					Position:       &Position{x, y},
					Indestructible: false,
				})
			case '-':
			default:
				return nil, fmt.Errorf("carácter no válido encontrado en la línea %d, columna %d: %c", row+1, col+1, char)
			}
			column = col + 1
		}
		row++
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	gameMap := GameMap{
		Walls:      walls,
		PowerUps:   []PowerUp{},
		Bombs:      []Bomb{},
		ColumnSize: int(column),
		RowSize:    int(row),
	}
	return &gameMap, nil
}

func GetRoundGameMap(roundNumber int8) *GameMap {
	filepath := fmt.Sprintf("data/round%dmap.txt", roundNumber)
	gameMap, err := CreateMap(filepath)
	if err != nil {
		panic(err)
	}
	return gameMap
}

func (m *GameMap) GetPowerUp(powerUpPosition Position) *PowerUp {
	for _, powerUp := range m.PowerUps {
		if powerUp.Position == powerUpPosition {
			return &powerUp
		}
	}
	return nil
}

func (m *GameMap) RemovePowerUp(powerUpPosition Position) {
	for i, powerUp := range m.PowerUps {
		if powerUp.Position == powerUpPosition {
			m.PowerUps = append(m.PowerUps[:i], m.PowerUps[i+1:]...)
		}
	}
}

func GetPowerUpType() PowerUpType {
	number := rand.Intn(100)
	switch {
	case number < 20:
		return Invencibilidad
	case number < 55:
		return AlcanceMejorado
	default:
		return MasBombasEnSimultaneo
	}
}

func (m *GameMap) existingPowerUpInPosition(position Position) bool {
	for _, powerUp := range m.PowerUps {
		if powerUp.Position == position {
			return true
		}
	}
	return false
}

func (m *GameMap) AddPowerUp(powerUpPosition *Position) {
	PowerUp := NewPowerUp(*powerUpPosition, GetPowerUpType())
	if m.existingPowerUpInPosition(*powerUpPosition) {
		return
	}
	m.PowerUps = append(m.PowerUps, PowerUp)
}

func (m *GameMap) PlaceBomb(bomb *Bomb) {
	m.Bombs = append(m.Bombs, *bomb)
}

func (m *GameMap) RemoveBomb(explodedBomb *Bomb) {
	for i, bomb := range m.Bombs {
		if explodedBomb.Position == bomb.Position {
			if len(m.Bombs) == 1 {
				m.Bombs = []Bomb{}
			} else {
				m.Bombs = append(m.Bombs[:i], m.Bombs[i+1:]...)
			}
		}
	}
}

func (m *GameMap) RemoveExplosion(explosion *Explosion) {
	for i, exp := range m.Explosions {
		if exp.Position == explosion.Position {
			if len(m.Explosions) == 1 {
				m.Explosions = []Explosion{}
			} else {
				m.Explosions = append(m.Explosions[:i], m.Explosions[i+1:]...)
			}
		}
	}
}

func (m *GameMap) isUnbreakableWall(position Position) bool {
	for _, wall := range m.Walls {
		if *wall.Position == position && wall.Indestructible {
			return true
		}
	}
	return false
}

func (m *GameMap) RemoveWall(position Position) {
	for i, wall := range m.Walls {
		if *wall.Position == position {
			m.Walls = append(m.Walls[:i], m.Walls[i+1:]...)
		}
	}
}
