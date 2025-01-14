package repositories

import (
	"time"

	"github.com/google/uuid"
)

// Este es el contrato que todos los eventos deben seguir
type Event interface {
	GetID() uuid.UUID
	GetTimestamp() time.Time
	GetType() string
	GetAggregateID() uuid.UUID
}
