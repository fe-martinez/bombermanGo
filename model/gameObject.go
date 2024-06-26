package model

import rl "github.com/gen2brain/raylib-go/raylib"

type GameObject struct {
	Position Position
	Size     float32
}

func (g GameObject) GetPosition() Position {
	return g.Position
}

func (g GameObject) GetSize() float32 {
	return g.Size
}

func (g GameObject) GetRect() rl.Rectangle {
	return rl.NewRectangle(g.Position.X*65, g.Position.Y*65, 65, 65)
}
