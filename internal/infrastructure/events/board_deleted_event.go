package events

import "github.com/google/uuid"

// Estructura para representar un evento cada vez que se cree un board nuevo
type BoardDeletedEvent struct {
	BaseEvent
	BoardID uuid.UUID
}
