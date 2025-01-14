package models

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID          uuid.UUID `json:"id"`
	BoardID     uuid.UUID `json:"boardid"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	State       string    `json:"state"`
	CreatedAt   time.Time `json:"createdat"`
}

type TaskQuery struct {
	ID          string    `bson:"_id"`
	Title       string    `bson:"titulo"`
	Description string    `bson:"descripcion"`
	State       string    `bson:"estado"`
	CreatedAt   time.Time `bson:"fecha_creacion"`
	BoardID     string    `bson:"id_tablero"`
}
