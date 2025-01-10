package models

import "github.com/google/uuid"

type Task struct {
	Id          uuid.UUID `json:"id"`
	BoardId     uuid.UUID `json:"boardid"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	State       string    `json:"state"`
	CreatedAt   string    `json:"createdat"`
}
