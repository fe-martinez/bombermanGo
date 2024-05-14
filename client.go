package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func client() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	decoder := json.NewDecoder(conn)
	var clientID int
	err = decoder.Decode(&clientID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Enter message: ")
	message := Message{
		Action: "join",
	}

	encoder := json.NewEncoder(conn)
	err = encoder.Encode(message)
	if err != nil {
		fmt.Println("Error encoding message:", err)
		return
	}

	rl.InitWindow(800, 800, "raylib [core] example - basic window")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)
	var game Game

	go func() {
		decoder := json.NewDecoder(conn)
		for {
			err := decoder.Decode(&game)
			if err != nil {
				fmt.Println("Error decoding game object:", err)
				return
			}
		}
	}()

	for !rl.WindowShouldClose() {
		drawGame(game)
		move := handleInput()

		if move == "none" {
			message := Message{
				Action:   "update",
				Data:     nil,
				PlayerID: clientID,
			}

			err = encoder.Encode(message)
			if err != nil {
				fmt.Println("Error encoding message:", err)
				return
			}
		} else {
			moveData := map[string]interface{}{
				"direction": move,
			}

			message := Message{
				Action:   "move",
				Data:     moveData,
				PlayerID: clientID,
			}

			err = encoder.Encode(message)
			if err != nil {
				fmt.Println("Error encoding message:", err)
				return
			}
		}

	}
}

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
	return "none"
}
