package server

import (
	"bombman/model"
	"bombman/utils"
	"log"
	"sync"
	"time"
)

const firstMapPath = "data/round1map.txt"

type Lobby struct {
	ownerID string
	id      string
	clients map[string]*Client
	updates chan utils.ClientMessage
	done    chan struct{}
	game    *model.Game
	mu      sync.RWMutex
}

func NewLobby(ownerID string, id string) *Lobby {
	gameMap, error := model.CreateMap(firstMapPath)
	if error != nil {
		log.Println("Error creating map")
		return nil
	}

	lobby := &Lobby{
		ownerID: ownerID,
		id:      id,
		clients: make(map[string]*Client),
		updates: make(chan utils.ClientMessage, 1000),
		done:    make(chan struct{}),
		game:    model.NewGame(id, gameMap),
	}
	go lobby.processInput()
	return lobby
}

func (l *Lobby) AddClient(client *Client) {
	l.mu.Lock()
	defer l.mu.Unlock()
	player := model.NewPlayer(client.clientID, l.game.GetPlayerPosition(len(l.game.Players)))
	l.game.AddPlayer(player)
	l.clients[client.clientID] = client

	client.lobbyID = l.id
	client.state = Game
}

func (l *Lobby) LeavingClientIsOwner(clientID string) bool {
	return clientID == l.ownerID
}

func (l *Lobby) AsignNewOwnerId(clientID string) {
	if !l.game.IsEmpty() {
		l.ownerID = l.game.RandomPlayerId()
	}
}

func (l *Lobby) RemoveClient(client *Client) {
	l.mu.Lock()
	defer l.mu.Unlock()

	delete(l.clients, client.clientID)
	l.game.RemovePlayer(client.clientID)
	if l.LeavingClientIsOwner(client.clientID) {
		l.AsignNewOwnerId(client.clientID)
	}
}

func (l *Lobby) Close() {
	l.game = nil
	close(l.done)
	time.Sleep(1 * time.Second)
	close(l.updates)
}

func (l *Lobby) startGame() {
	if l.game.State == "not-started" {
		l.game.Start()
	}
}

func (l *Lobby) handlePlayerInput(input utils.ClientMessage) {
	handlePlayerAction(input, l.game)
}

func (l *Lobby) processInput() {
	ticker := time.NewTicker(GAME_SPEED)
	defer ticker.Stop()

	for {
		select {
		case <-l.done:
			return
		case input := <-l.updates:
			l.handlePlayerInput(input)
		case <-ticker.C:
			if l.game.State == model.Finished {
				return
			}
			l.game.Update()
			l.BroadcastGameState()
		}
	}
}

func (l *Lobby) BroadcastGameState() {
	l.mu.RLock()
	defer l.mu.RUnlock()

	encodedGame, err := utils.EncodeGame(*l.game)
	if err != nil {
		log.Println("Error encoding game:", err)
		return
	}

	for _, client := range l.clients {
		sendGameMessageToClient(client.connection, encodedGame)
	}
}
