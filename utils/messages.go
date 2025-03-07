package utils

import (
	"bytes"
	"encoding/gob"
)

type ClientMessage struct {
	Action Action
	Data   interface{}
	ID     string
}

type Action int

const (
	ActionMove Action = iota
	ActionBomb
	ActionLeave
	ActionJoinGame
	ActionCreateGame
	ActionStartGame
	ActionMainMenu
)

func EncodeClientMessage(msg ClientMessage) ([]byte, error) {
	buf := new(bytes.Buffer)

	enc := gob.NewEncoder(buf)

	err := enc.Encode(msg)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func DecodeClientMessage(encodedMsg []byte) (ClientMessage, error) {
	buf := bytes.NewBuffer(encodedMsg)

	var msg ClientMessage

	dec := gob.NewDecoder(buf)

	err := dec.Decode(&msg)
	if err != nil {
		return ClientMessage{}, err
	}

	return msg, nil
}
