package client

import (
	"bombman/utils"
	"fmt"
	"net"
)

type EventType int

const (
	EventCreateGame EventType = EventType(utils.ActionCreateGame)
	EventJoinGame   EventType = EventType(utils.ActionJoinGame)
	EventStartGame  EventType = EventType(utils.ActionStartGame)
	EventMove       EventType = EventType(utils.ActionMove)
	EventPlaceBomb  EventType = EventType(utils.ActionBomb)
	EventLeave      EventType = EventType(utils.ActionLeave)
	EventMainMenu   EventType = EventType(utils.ActionMainMenu)
)

type GameEvent struct {
	Type     EventType
	PlayerID string
	Payload  interface{}
}

type EventEmitter struct {
	connection net.Conn
}

func (e *EventEmitter) Emit(event GameEvent) error {
	encodedEvent, err := utils.EncodeClientMessage(utils.ClientMessage{
		Action: utils.Action(event.Type),
		ID:     event.PlayerID,
		Data:   event.Payload,
	})

	if err != nil {
		return fmt.Errorf("failed to encode event: %v", err)
	}

	_, err = e.connection.Write(encodedEvent)
	if err != nil {
		return fmt.Errorf("failed to send event: %v", err)
	}
	return nil
}
