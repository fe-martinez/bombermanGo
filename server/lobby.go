package server

import (
	"bombman/model"
	"log"
)

const firstMapPath = "data/round1map.txt"

type Lobby struct {
	ownerID string
	id      string
	clients map[string]*Client
	game    *model.Game
}

func NewLobby(ownerID string, id string) *Lobby {
	gameMap, error := model.CreateMap(firstMapPath)
	if error != nil {
		log.Println("Error creating map")
		return nil
	}

	return &Lobby{
		ownerID: ownerID,
		id:      id,
		clients: make(map[string]*Client),
		game:    model.NewGame(id, gameMap),
	}
}

func (l *Lobby) AddClient(client *Client) {
	player := model.NewPlayer(client.clientID, l.game.GenerateValidPosition(l.game.GameMap.RowSize))
	l.game.AddPlayer(player)
	l.clients[client.clientID] = client

	client.lobbyID = l.id
	client.state = Game
}

func (l *Lobby) RemoveClient(client *Client) {
	delete(l.clients, client.clientID)
	l.game.RemovePlayer(client.clientID)
}

func (l *Lobby) startGame() {
	if l.game.State == "not-started" {
		l.game.Start()
	}
}

func (l *Lobby) BroadcastGameState() {
	for _, client := range l.clients {
		sendGameMessageToClient(client.connection, l.game)
	}
}
