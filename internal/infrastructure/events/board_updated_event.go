package events

import "github.com/google/uuid"

// Estructura para representar un evento cada vez que se cree un board nuevo
type BoardUpdatedEvent struct {
	BaseEvent
	BoardID     uuid.UUID
	Name        string `json:"nombre"`
	Description string `json:"descripcion"`
}
