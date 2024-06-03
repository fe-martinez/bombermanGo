package server

import (
	"bombman/model"
	"log"
)

type Lobby struct {
	ownerID string
	id      string
	clients map[string]*Client
	game    *model.Game
}

func NewLobby(ownerID string, id string) *Lobby {
	return &Lobby{
		ownerID: ownerID,
		id:      id,
		clients: make(map[string]*Client),
		game:    model.NewGame(id, model.CreateMap(15, 16)),
	}
}

func (l *Lobby) AddClient(client *Client) {
	player := model.NewPlayer(client.clientID, l.game.GenerateValidPosition(l.game.GameMap.Size))
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
		log.Println("Sending game message to client", client.clientID)
		sendGameMessageToClient(client.connection, l.game)
	}
}
