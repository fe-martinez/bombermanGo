package model

import rl "github.com/gen2brain/raylib-go/raylib"

type Wall struct {
	Position       *Position
	Indestructible bool
}

func (w Wall) GetPosition() Position {
	return *w.Position
}

func (w Wall) GetSize() float32 {
	return 65
}

func (w Wall) GetRect() rl.Rectangle {
	return rl.NewRectangle(w.Position.X*65, w.Position.Y*65, 65, 65)
}
