package main

import (
	"math/rand"
)

type Map struct {
	Walls   []Wall
	Bombs   []Bomb
	Enemies []Enemy
	size    int
}

type Enemy struct {
	X     float32
	Y     float32
	lives int8
}

type Wall struct {
	X              float32
	Y              float32
	indestructible bool
}

type BombQualities struct {
	alcance float32
	power   int8
}

type Bomb struct {
	X             float32
	Y             float32
	Timer         float32
	bombQualities BombQualities
}

// Esta función hay que cambiarla para que ponga algún mapa predefinido
// Podemos hacer algoritmos de mapas según el nivel así que recibiríamos el nivel
func createMap(wallsAmount int, mapSize int) Map {
	walls := []Wall{}
	for i := 0; i < wallsAmount; i++ {
		walls = append(walls, Wall{
			X:              float32(rand.Intn(mapSize)),
			Y:              float32(rand.Intn(mapSize)),
			indestructible: false,
		})
	}

	enemies := []Enemy{}
	for i := 0; i < 4; i++ {
		enemies = append(enemies, Enemy{
			X:     float32(rand.Intn(mapSize)),
			Y:     float32(rand.Intn(mapSize)),
			lives: 1,
		})
	}

	gameMap := Map{
		Walls:   walls,
		Bombs:   []Bomb{},
		Enemies: enemies,
		size:    mapSize,
	}

	return gameMap
}
