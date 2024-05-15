package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func startClient() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	decoder := json.NewDecoder(conn)
	var clientID string
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

		if move == "bomb" {
			message := Message{
				Action:   "bomb",
				Data:     nil,
				PlayerID: clientID,
			}

			err = encoder.Encode(message)
			if err != nil {
				fmt.Println("Error encoding message:", err)
				return
			}
		} else if move == "none" {
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

		if rl.WindowShouldClose() {
			message := Message{
				Action:   "leave",
				Data:     nil,
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
