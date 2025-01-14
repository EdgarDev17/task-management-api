package events

import "github.com/google/uuid"

type TaskDeletedEvent struct {
	BaseEvent
	TaskID uuid.UUID
}
