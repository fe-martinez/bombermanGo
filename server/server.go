package server

import (
	"bombman/model"
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
	clients  map[string]Client
}

type Lobby struct {
	ownerID string
	id      string
	clients map[string]net.Conn
	game    *model.Game
}

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
		clients:  make(map[string]Client),
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

	s.clients[playerID] = client
	sendId(conn, playerID)

	for {
		s.handleMessages(&client)
	}
}

func (s *Server) handleMessages(client *Client) {
	message, err := readClientMessage(client.connection)
	if err != nil {
		return
	}

	if client.state == MainMenu {
		s.handleMainMenuAction(message, client)
	} else {
		lobby := s.lobbies[client.lobbyID]
		handlePlayerAction(message, lobby.game)
	}
}

func (s *Server) handleMainMenuAction(msg utils.ClientMessage, client *Client) {
	switch msg.Action {
	case utils.ActionCreateGame:
		lobby := s.createLobby(client)
		if lobby == nil {
			log.Println("Error creating lobby")
			return
		}
	case utils.ActionJoinGame:
		lobbyID := msg.Data.(string)
		lobby := s.joinLobby(lobbyID, client)
		if lobby == nil {
			log.Println("Error joining lobby")
			return
		}
	default:
		fmt.Println("Action unknown")
	}
}

func (s *Server) createUniqueLobbyID() string {
	randomValue := strconv.Itoa(rand.Intn(1000))
	for s.lobbies[randomValue] != nil {
		randomValue = strconv.Itoa(rand.Intn(1000))
	}
	return randomValue
}

func (s *Server) addClientToLobby(lobbyID string, client *Client) {
	lobby := s.lobbies[lobbyID]
	player := model.NewPlayer(client.clientID, lobby.game.GenerateValidPosition(lobby.game.GameMap.Size))

	lobby.game.AddPlayer(player)
	lobby.clients[client.clientID] = client.connection
	s.lobbies[lobbyID] = lobby

	client.state = Game
	s.clients[client.clientID] = *client
}

func (s *Server) createLobby(client *Client) *Lobby {
	lobbyID := s.createUniqueLobbyID()
	gameMap := model.CreateMap(15, 16)
	lobby := &Lobby{
		ownerID: client.clientID,
		id:      lobbyID,
		clients: make(map[string]net.Conn),
		game:    model.NewGame(lobbyID, gameMap),
	}
	s.lobbies[lobbyID] = lobby
	s.addClientToLobby(lobbyID, client)
	return lobby
}

func (s *Server) joinLobby(lobbyID string, client *Client) *Lobby {
	lobby := s.lobbies[lobbyID]
	if lobby == nil {
		log.Println("Lobby", lobbyID, "not found")
		return nil
	}
	s.addClientToLobby(lobbyID, client)
	return lobby
}

func (s *Server) disconnectClient(clientID string) {
	client := s.clients[clientID]
	if client.state == Game {
		lobby := s.lobbies[client.lobbyID]
		lobby.game.RemovePlayer(clientID)
		delete(lobby.clients, clientID)
		if len(lobby.clients) == 0 {
			delete(s.lobbies, client.lobbyID)
		}
	}
	delete(s.clients, clientID)
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
		gameState := lobby.game
		for _, conn := range lobby.clients {
			sendMessageToClient(conn, gameState)
		}
	}
}

func listen(serverAddress string) (net.Listener, error) {
	listener, err := net.Listen("tcp", serverAddress)
	if err != nil {
		return nil, err
	}
	return listener, nil
}
