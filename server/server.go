package server

import (
	"bombman/utils"
	"fmt"
	"log"
	"math/rand"
	"net"
	"strconv"
	"time"
)

type Server struct {
	address  string
	listener net.Listener
	lobbies  map[string]*Lobby
	clients  map[string]*Client
}

const SERVER_ADDRESS = "localhost:8080"
const GAME_SPEED = 33 * time.Millisecond

func NewServer(address string, maxPlayers int) (*Server, error) {
	listener, err := listen(SERVER_ADDRESS)
	if err != nil {
		return nil, err
	}

	return &Server{
		address:  address,
		listener: listener,
		clients:  make(map[string]*Client),
		lobbies:  make(map[string]*Lobby),
	}, nil
}

func (s *Server) Start() {
	log.Println("Starting game server at", SERVER_ADDRESS)
	go s.handleConnections()
	go s.broadcastLoop()
	select {}
}

func (s *Server) handleConnections() {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v\n", err)
			continue
		}

		go s.handleConnection(conn)
	}
}

func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()

	playerID := utils.CreateRandomUid()

	client := Client{clientID: playerID, connection: conn, state: MainMenu, lobbyID: ""}

	s.clients[playerID] = &client
	sendPlayerId(conn, playerID)

	for {
		err := s.handleMessage(&client)
		if err != nil {
			log.Printf("Error handling client message: %v\n", err)
		}
	}
}

func (s *Server) handleMessage(client *Client) error {
	message, err := readClientMessage(client.connection)
	if err != nil {
		return err
	}

	if client.state == MainMenu {
		s.handleMainMenuAction(message, client)
	} else {
		lobby := s.lobbies[client.lobbyID]
		if message.Action == utils.ActionLeave {
			s.disconnectClient(client.clientID)
			return nil
		}
		handlePlayerAction(message, lobby.game)
	}

	return nil
}

func (s *Server) handleMainMenuAction(msg utils.ClientMessage, client *Client) {
	switch msg.Action {
	case utils.ActionCreateGame:
		s.createLobby(client)
	case utils.ActionJoinGame:
		lobbyID := msg.Data.(string)
		s.joinLobby(lobbyID, client)
	default:
		fmt.Println("Action unknown")
	}
}

func (s *Server) createLobby(client *Client) {
	lobbyID := s.createUniqueLobbyID()
	lobby := NewLobby(client.clientID, lobbyID)
	lobby.AddClient(client)
	s.lobbies[lobbyID] = lobby

	if lobby == nil {
		log.Println("Failed to create lobby")
	} else {
		sendLobbyId(client.connection, lobbyID)
	}
}

func (s *Server) joinLobby(lobbyID string, client *Client) {
	lobby := s.lobbies[lobbyID]
	if lobby == nil {
		log.Println("Lobby", lobbyID, "not found")
		sendJoinLobbyFailure(client.connection, lobbyID)
	} else {
		sendJoinLobbySuccess(client.connection, lobbyID)
		lobby.AddClient(client)
	}
}

func (s *Server) disconnectClient(clientID string) {
	client := s.clients[clientID]
	if client.state == Game {
		lobby := s.lobbies[client.lobbyID]
		lobby.RemoveClient(client)
		if len(lobby.clients) == 0 {
			delete(s.lobbies, client.lobbyID)
		}
	}
	delete(s.clients, clientID)
}

func (s *Server) createUniqueLobbyID() string {
	randomValue := strconv.Itoa(rand.Intn(1000))
	for s.lobbies[randomValue] != nil {
		randomValue = strconv.Itoa(rand.Intn(1000))
	}
	return randomValue
}

func (s *Server) broadcastLoop() {
	ticker := time.NewTicker(GAME_SPEED)
	defer ticker.Stop()

	for {
		<-ticker.C
		s.broadcastGameState()
	}
}

func (s *Server) broadcastGameState() {
	for _, lobby := range s.lobbies {
		lobby.BroadcastGameState()
	}
}

func listen(serverAddress string) (net.Listener, error) {
	listener, err := net.Listen("tcp", serverAddress)
	if err != nil {
		return nil, err
	}
	return listener, nil
}
