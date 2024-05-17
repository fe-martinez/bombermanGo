package model

import "math/rand"

type GameMap struct {
	Walls    []Wall
	PowerUps []PowerUp
	Bombs    []Bomb
	Size     int
}

type Wall struct {
	Position       *Position
	Indestructible bool
}

// Por ahora crea un mapa con paredes en posiciones aleatorias. Hay que armar bien los algoritmos acá jeje
func CreateMap(wallsAmount int, mapSize int) *GameMap {
	walls := []Wall{}
	for i := 0; i < wallsAmount; i++ {
		WallPosition := &Position{float32(rand.Intn(mapSize)), float32(rand.Intn(mapSize))}
		walls = append(walls, Wall{
			Position:       WallPosition,
			Indestructible: false,
		})
	}

	gameMap := GameMap{
		Walls:    walls,
		PowerUps: []PowerUp{},
		Bombs:    []Bomb{},
		Size:     mapSize,
	}

	return &gameMap
}
