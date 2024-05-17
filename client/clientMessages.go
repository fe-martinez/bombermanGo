package client

import (
	"bombman/utils"
	"fmt"
	"net"
)

func SendMessage(msg utils.ClientMessage, conn net.Conn) error {
	encodedMsg, err := utils.EncodeClientMessage(msg)
	if err != nil {
		return fmt.Errorf("failed to encode message: %w", err)
	}

	_, err = conn.Write(encodedMsg)
	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}

	return nil
}

func SendMoveMessage(move string, conn net.Conn, playerID string) error {
	msg := utils.ClientMessage{
		Action: utils.ActionMove,
		Data:   move,
		ID:     playerID,
	}

	return SendMessage(msg, conn)
}

func SendBombMessage(conn net.Conn, playerID string) error {
	msg := utils.ClientMessage{
		Action: utils.ActionBomb,
		Data:   nil,
		ID:     playerID,
	}
	return SendMessage(msg, conn)
}

func SendUpdateMessage(conn net.Conn, playerID string) error {
	msg := utils.ClientMessage{
		Action: utils.ActionUpdate,
		Data:   nil,
		ID:     playerID,
	}
	return SendMessage(msg, conn)
}

func SendLeaveMessage(conn net.Conn, playerID string) error {
	msg := utils.ClientMessage{
		Action: utils.ActionLeave,
		ID:     playerID,
	}
	return SendMessage(msg, conn)
}
