package main

import (
	"bombman/client"
	"bombman/server"
	"fmt"
	"os"
)

const SERVER_ADDRESS = "192.168.0.2:8080"

func validateArgs() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . [client|server]")
		fmt.Println("P.S.: remember to run the server before you try to run the client!")
		os.Exit(1)
	}
}

func runClient() {
	client := client.NewClient()
	client.Start()
}

func runServer() {
	server, err := server.NewServer(SERVER_ADDRESS, 4)
	if err != nil {
		fmt.Println("Error while starting server: ", err)
	}

	server.Start()
}

func run(arg string) {
	switch arg {
	case "client":
		runClient()
	case "server":
		runServer()
	default:
		fmt.Println("Invalid argument. Usage: go run . [client|server]")
	}
}

func main() {
	validateArgs()

	arg := os.Args[1]

	run(arg)
}
