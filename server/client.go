package server

import (
	"net"
)

type Client struct {
	clientID   string
	connection net.Conn
	state      ClientState
	lobbyID    string
}

type ClientState string

const (
	MainMenu ClientState = "main-menu"
	Game     ClientState = "game"
)
