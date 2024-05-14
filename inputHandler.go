package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func handleInput() string {
	if rl.IsKeyDown(rl.KeyW) {
		return "up"
	}
	if rl.IsKeyDown(rl.KeyS) {
		return "down"
	}
	if rl.IsKeyDown(rl.KeyA) {
		return "left"
	}
	if rl.IsKeyDown(rl.KeyD) {
		return "right"
	}
	if rl.IsKeyDown(rl.KeyB) {
		return "bomb"
	}

	return "none"
}
