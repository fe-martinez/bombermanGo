package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args)

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [client|server]")
		return
	}

	arg := os.Args[1]

	switch arg {
	case "client":
		runClient()
	case "server":
		runServer()
	default:
		fmt.Println("Invalid argument. Usage: go run main.go [client|server]")
	}
}

func runClient() {
	client()
	fmt.Println("Running client...")
}

func runServer() {
	server()
	fmt.Println("Running server...")
}
