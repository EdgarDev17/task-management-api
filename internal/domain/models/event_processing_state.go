package models

import (
	"time"

	"github.com/google/uuid"
)

type ProcessingState struct {
	ID                     int64      `db:"id"`
	LastProcessedID        *uuid.UUID `db:"last_processed_id"`
	LastProcessedTimestamp time.Time  `db:"last_processed_timestamp"`
	UpdatedAt              time.Time  `db:"updated_at"`
}
