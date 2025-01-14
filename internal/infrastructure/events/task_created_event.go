package events

import "github.com/google/uuid"

type TaskCreatedEvent struct {
	BaseEvent
	TaskID      uuid.UUID
	BoardID     uuid.UUID
	Title       string `json:"titulo"`
	State       string `json:"estado"`
	Description string `json:"descripcion"`
}
