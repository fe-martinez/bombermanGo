package main

import (
	"math/rand"
)

type Map struct {
	Walls []Wall
	size  int
}

type Wall struct {
	X              float32
	Y              float32
	indestructible bool
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

	gameMap := Map{
		Walls: walls,
		size:  mapSize,
	}

	return gameMap
}
