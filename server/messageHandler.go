package server

import (
	"bombman/model"
	"bombman/utils"
	"encoding/gob"
	"fmt"
	"net"
)

type JoinLobbyAck struct {
	Success bool
	LobbyID string
	Message string
}

func readClientMessage(conn net.Conn) (utils.ClientMessage, error) {
	buffer := make([]byte, 1024)
	_, err := conn.Read(buffer)
	if err != nil {
		return utils.ClientMessage{}, fmt.Errorf("error al leer del cliente: %s", err)
	}

	clientMsg, err := utils.DecodeClientMessage(buffer)
	if err != nil {
		return utils.ClientMessage{}, fmt.Errorf("error al decodificar el mensaje del cliente: %s", err)
	}

	return clientMsg, nil
}

func sendMessageToClient(conn net.Conn, game *model.Game) {
	encodedGame, err := utils.EncodeGame(*game)
	if err != nil {
		fmt.Println("Error encoding game:", err)
		return
	}
	_, err = conn.Write(encodedGame)
	if err != nil {
		fmt.Println("Error al enviar el juego al cliente:", err)
		return
	}
}

func sendJoinLobbySuccess(conn net.Conn, lobbyID string) {
	var msg JoinLobbyAck
	msg.LobbyID = lobbyID
	msg.Success = true
	msg.Message = "Joined lobby " + lobbyID

	enc := gob.NewEncoder(conn)
	err := enc.Encode(msg)
	if err != nil {
		fmt.Println("Error encoding join lobby success: ", err)
	}
}

func sendJoinLobbyFailure(conn net.Conn, lobbyID string) {
	var msg JoinLobbyAck
	msg.LobbyID = lobbyID
	msg.Success = false
	msg.Message = "Failed to join lobby " + lobbyID

	enc := gob.NewEncoder(conn)
	err := enc.Encode(msg)
	if err != nil {
		fmt.Println("Error encoding join lobby failure: ", err)
	}
}

func sendId(conn net.Conn, playerID string) {
	_, err := conn.Write([]byte(playerID))
	if err != nil {
		fmt.Println("Error sending player ID to client: ", err)
	}
}
