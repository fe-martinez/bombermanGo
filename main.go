package main

import (
	"bombman/client"
	"bombman/server"
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	ServerAddress string `json:"server_address"`
}

func validateArgs() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . [client|server]")
		fmt.Println("P.S.: remember to run the server before you try to run the client!")
		os.Exit(1)
	}
}

func getConfig() Config {
	file, err := os.Open("config.json")

	if err != nil {
		fmt.Println("Error while opening config file: ", err)
		return Config{}
	}

	defer file.Close()

	var config Config
	err = json.NewDecoder(file).Decode(&config)

	if err != nil {
		fmt.Println("Error while decoding config file: ", err)
		return Config{}
	}

	return config
}

func runClient() {
	config := getConfig()
	address := config.ServerAddress
	if address == "" {
		address = "localhost:8080"
	}

	client := client.NewClient(address)
	client.Start()
}

func runServer() {
	config := getConfig()
	address := config.ServerAddress
	if address == "" {
		address = "localhost:8080"
	}

	server, err := server.NewServer(address, 4)
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
