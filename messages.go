package main

import (
	"bytes"
	"encoding/gob"
)

type ClientMessage struct {
	Action string
	Data   interface{}
	ID     string
}

func encodeClientMessage(msg ClientMessage) ([]byte, error) {
	buf := new(bytes.Buffer)

	enc := gob.NewEncoder(buf)

	err := enc.Encode(msg)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func decodeClientMessage(encodedMsg []byte) (ClientMessage, error) {
	buf := bytes.NewBuffer(encodedMsg)

	var msg ClientMessage

	dec := gob.NewDecoder(buf)

	err := dec.Decode(&msg)
	if err != nil {
		return ClientMessage{}, err
	}

	return msg, nil
}
