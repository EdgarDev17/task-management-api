package events

import "github.com/google/uuid"

type TaskStateUpdatedEvent struct {
	BaseEvent
	TaskID      uuid.UUID `json:"task_id"`
	BoardID     uuid.UUID `json:"board_id"`
	NewState    string    `json:"new_state"`
	Title       string    `json:"titulo"`
	Description string    `json:"descripcion"`
}
