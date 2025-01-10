package repositories

import (
	"task-management-api/internal/domain/models"

	"github.com/google/uuid"
)

type BoardRepository interface {
	Create(board *models.Board) error
	GetByID(id uuid.UUID) (*models.Board, error)
	GetAll() ([]*models.Board, error)
	Update(board *models.Board) error
	Delete(id uuid.UUID) error
}
