package model

import "math/rand"

type Position struct {
	X float32
	Y float32
}

func getRandomPosition(rowSize int, columnSize int) *Position {
	return &Position{float32(rand.Intn(rowSize)), float32(rand.Intn(columnSize))}
}
