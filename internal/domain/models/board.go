package models

import (
	"github.com/google/uuid"
)

type Board struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   string    `json:"createdat"`
	UpdatedAt   string    `json:"updatedat"`
}
