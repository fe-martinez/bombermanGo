package main

import "math/rand"

type Position struct {
	X float32
	Y float32
}

func getRandomPosition(mapSize int) *Position {
	return &Position{float32(rand.Intn(mapSize)), float32(rand.Intn(mapSize))}
}
