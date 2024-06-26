package model

import rl "github.com/gen2brain/raylib-go/raylib"

type Positionable interface {
	GetPosition() Position
	GetSize() float32
	GetRect() rl.Rectangle
}
