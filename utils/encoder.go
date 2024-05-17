package utils

import (
	"bombman/model"
	"bytes"
	"encoding/gob"
)

func EncodeGame(game model.Game) ([]byte, error) {
	buf := new(bytes.Buffer)

	enc := gob.NewEncoder(buf)

	err := enc.Encode(game)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func DecodeGame(encodedGame []byte) (*model.Game, error) {
	buf := bytes.NewBuffer(encodedGame)

	dec := gob.NewDecoder(buf)

	var game model.Game

	err := dec.Decode(&game)
	if err != nil {
		return nil, err
	}

	return &game, nil
}
