package server

import (
	"bombman/model"
	"bombman/utils"
	"log"
	"time"
)

const firstMapPath = "data/round1map.txt"

type Lobby struct {
	ownerID string
	id      string
	clients map[string]*Client
	updates chan utils.ClientMessage
	game    *model.Game
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
		updates: make(chan utils.ClientMessage),
		game:    model.NewGame(id, gameMap),
	}
	go lobby.processInput()
	return lobby
}

func (l *Lobby) AddClient(client *Client) {
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
		l.ownerID = l.game.GetAPlayerId()
	}
}

func (l *Lobby) RemoveClient(client *Client) {
	delete(l.clients, client.clientID)
	l.game.RemovePlayer(client.clientID)
	if l.LeavingClientIsOwner(client.clientID) {
		l.AsignNewOwnerId(client.clientID)
	}
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
		case input := <-l.updates:
			l.handlePlayerInput(input)
		case <-ticker.C:
			l.game.Update()
			l.BroadcastGameState()
		}
	}
}

func (l *Lobby) BroadcastGameState() {
	for _, client := range l.clients {
		sendGameMessageToClient(client.connection, l.game)
	}
}
