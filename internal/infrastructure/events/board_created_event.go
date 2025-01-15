package events

import "github.com/google/uuid"

// Estructura para representar un evento cada vez que se cree un board nuevo
type BoardCreatedEvent struct {
	BaseEvent
	BoardID     uuid.UUID
	Name        string `json:"nombre"`
	Description string `json:"descripcion"`
}
