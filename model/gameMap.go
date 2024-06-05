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
					Indestructible: false,
				})
			case 'D':
				walls = append(walls, Wall{
					Position:       &Position{x, y},
					Indestructible: true,
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

func (m *GameMap) GetPowerUp(powerUpPosition Position) *PowerUp {
	for _, powerUp := range m.PowerUps {
		if powerUp.Position == powerUpPosition {
			return &powerUp
		}
	}
	return nil
}

func (m *GameMap) RemovePowerUp(powerUpPosition Position) *PowerUp {
	powerUp := m.GetPowerUp(powerUpPosition)
	for i, powerUp := range m.PowerUps {
		if powerUp.Position == powerUpPosition {
			m.PowerUps = append(m.PowerUps[:i], m.PowerUps[i+1:]...)
		}
	}
	return powerUp
}

func (m *GameMap) AddPowerUp(powerUpPosition *Position) {
	PowerUp := NewPowerUp(*powerUpPosition, PowerUpType(rand.Intn(5)))
	m.PowerUps = append(m.PowerUps, PowerUp)
}
