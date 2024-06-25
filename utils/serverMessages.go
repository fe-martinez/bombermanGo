package utils

type ServerMessage struct {
	Action MessageType
	Data   interface{}
}

type MessageType int

const (
	LobbyIDMessage  MessageType = 0
	PlayerIDMessage MessageType = 1
	JoinLobbyAck    MessageType = 2
)
