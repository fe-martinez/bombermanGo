package server

import (
	"bombman/utils"
	"errors"
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
	s.handleConnections()
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
			s.disconnectClient(playerID)
			return
		}
	}
}

func (s *Server) handleMessage(client *Client) error {
	message, err := readClientMessage(client.connection)
	if err != nil {
		return err
	}

	switch client.state {
	case MainMenu:
		return s.handleMainMenuAction(message, client)
	case WaitingMenu, Game, WaitingInGame:
		return s.handleGameAction(message, client)
	default:
		return fmt.Errorf("unknown client state: %v", client.state)
	}
}

func (s *Server) handleMainMenuAction(msg utils.ClientMessage, client *Client) error {
	if client.lobbyID != "" {
		log.Println("Client is already in another lobby")
	}

	switch msg.Action {
	case utils.ActionCreateGame:
		s.createLobby(client)
	case utils.ActionJoinGame:
		lobbyID := msg.Data.(string)
		s.joinLobby(lobbyID, client)
	default:
		fmt.Println("This action unknown")
	}

	log.Println("Client joined lobby", client.lobbyID)
	return nil
}

func (s *Server) handleGameAction(message utils.ClientMessage, client *Client) error {
	if message.Action == utils.ActionLeave {
		s.disconnectClient(client.clientID)
		return nil
	}
	lobby, exists := s.lobbies[client.lobbyID]

	if !exists {
		return fmt.Errorf("lobby not found: %s", client.lobbyID)
	}

	if message.Action == utils.ActionMainMenu {
		s.removeClientFromLobby(client.clientID)
		return nil
	}

	if message.Action == utils.ActionStartGame && client.clientID == lobby.ownerID {
		lobby.startGame()
		return nil
	}

	select {
	case lobby.updates <- message:
	default:
		log.Printf("Input channel full for lobby %s", client.lobbyID)
	}

	return nil
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

	log.Println("Created lobby", lobbyID, "succesfully")
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
	client, exists := s.clients[clientID]
	if !exists || client == nil {
		return
	}

	if client.state == Game {
		lobby, exists := s.lobbies[client.lobbyID]
		if !exists || lobby == nil {
			return
		}

		lobby.RemoveClient(client)
		if len(lobby.clients) == 0 {
			lobby.game.Stop()
			lobby.Close()
			delete(s.lobbies, client.lobbyID)
			log.Println("Lobby", client.lobbyID, "deleted succesfully")
		}
	}
	delete(s.clients, clientID)
}

func (s *Server) removeClientFromLobby(clientID string) error {
	client, exists := s.clients[clientID]
	if !exists || client == nil {
		return errors.New("Client " + clientID + " was not found")
	}

	lobby, exists := s.lobbies[client.lobbyID]
	if !exists || lobby == nil {
		return errors.New("Lobby " + client.lobbyID + " was not found")
	}

	lobby.RemoveClient(client)
	if len(lobby.clients) == 0 {
		lobby.game.Stop()
		lobby.Close()
		delete(s.lobbies, client.lobbyID)
		log.Println("Lobby", client.lobbyID, "deleted succesfully")
		log.Println("Remaining lobbies:", len(s.lobbies))
	}

	client.state = MainMenu
	client.lobbyID = ""

	log.Println("Client", clientID, "returned to main menu")
	return nil
}

func (s *Server) createUniqueLobbyID() string {
	randomValue := strconv.Itoa(rand.Intn(900) + 100)
	for s.lobbies[randomValue] != nil {
		randomValue = strconv.Itoa(rand.Intn(900) + 100)
	}
	return randomValue
}

func listen(serverAddress string) (net.Listener, error) {
	listener, err := net.Listen("tcp", serverAddress)
	if err != nil {
		return nil, err
	}
	return listener, nil
}
