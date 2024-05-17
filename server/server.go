package server

import (
	"bombman/model"
	"bombman/utils"
	"fmt"
	"net"
	"time"
)

type Server struct {
	address  string
	listener net.Listener
	game     *model.Game
	clients  map[string]net.Conn
}

const SERVER_ADDRESS = "localhost:8080"

func NewServer(address string, maxPlayers int) (*Server, error) {
	listener, err := listen(SERVER_ADDRESS)
	if err != nil {
		return nil, err
	}

	gameMap := model.CreateMap(15, 16)
	gameId := utils.CreateRandomUid()
	game := model.NewGame(gameId, gameMap)

	return &Server{
		address:  address,
		listener: listener,
		game:     game,
		clients:  make(map[string]net.Conn),
	}, nil
}

func (s *Server) Start() {
	fmt.Println("Starting game server at", SERVER_ADDRESS)
	go s.handleConnections()
	go s.broadcastLoop()
	select {}
}

func (s *Server) handleConnections() {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			fmt.Printf("Error accepting connection: %v\n", err)
			continue
		}

		go s.handleConnection(conn)
	}
}

func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()
	if s.game.IsFull() {
		return
	}

	playerID := utils.CreateRandomUid()
	mapSize := s.game.GameMap.Size
	player := model.NewPlayer(playerID, s.game.GenerateValidPosition(mapSize))
	s.game.AddPlayer(player)

	s.clients[playerID] = conn
	sendId(conn, playerID)

	for {
		s.handleMessages(conn, s.game)
	}
}

func (s *Server) handleMessages(conn net.Conn, game *model.Game) {
	message, err := readClientMessage(conn)
	if err != nil {
		return
	}
	handlePlayerMove(message, game)
}

func (s *Server) broadcastLoop() {
	ticker := time.NewTicker(33 * time.Millisecond)
	defer ticker.Stop()

	for {
		<-ticker.C
		s.broadcastGameState()
	}
}

func (s *Server) broadcastGameState() {
	gameState := s.game
	for _, conn := range s.clients {
		sendMessageToClient(conn, gameState)
	}
}

func listen(serverAddress string) (net.Listener, error) {
	listener, err := net.Listen("tcp", serverAddress)
	if err != nil {
		return nil, err
	}
	return listener, nil
}
