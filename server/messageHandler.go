package server

import (
	"bombman/model"
	"bombman/utils"
	"fmt"
	"net"
)

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
	encodedGame, err := model.EncodeGame(*game)
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

func sendId(conn net.Conn, playerID string) {
	n, err := conn.Write([]byte(playerID))
	fmt.Println(n)
	if err != nil {
		fmt.Println("Error sending player ID to client: ", err)
	}
}
