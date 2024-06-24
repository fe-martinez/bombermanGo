package utils

import (
	"bytes"
	"encoding/gob"
)

type ClientMessage struct {
	Action MessageAction
	Data   interface{}
	ID     string
}

type MessageAction int

const (
	ActionMove       MessageAction = 0
	ActionBomb       MessageAction = 1
	ActionLeave      MessageAction = 2
	ActionJoinGame   MessageAction = 3
	ActionCreateGame MessageAction = 4
	ActionStartGame  MessageAction = 5
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
