package main

import (
	"fmt"
	"os"
)

func validateArgs() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go [client|server]")
		fmt.Println("P.S.: remember to run the server before you try to run the client!")
		os.Exit(1)
	}
}

func runClient() {
	startClient()
}

func runServer() {
	startServer()
}

func run(arg string) {
	switch arg {
	case "client":
		runClient()
	case "server":
		runServer()
	default:
		fmt.Println("Invalid argument. Usage: go run main.go [client|server]")
	}
}

func main() {
	validateArgs()

	arg := os.Args[1]

	run(arg)
}
