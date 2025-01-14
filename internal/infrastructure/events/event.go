package events

import (
	"time"

	"github.com/google/uuid"
)

// Cada evento debe heredar esta estructura, esta es la base
type BaseEvent struct {
	ID          uuid.UUID `json:"id"`
	AggregateID uuid.UUID `json:"aggregate_id"`
	Timestamp   time.Time `json:"timestamp"`
	Type        string    `json:"type"`
	Data        string    `json:"data"`
}

// Implementando los metodos

func (e BaseEvent) GetID() uuid.UUID          { return e.ID }
func (e BaseEvent) GetTimestamp() time.Time   { return e.Timestamp }
func (e BaseEvent) GetType() string           { return e.Type }
func (e BaseEvent) GetAggregateID() uuid.UUID { return e.AggregateID }
