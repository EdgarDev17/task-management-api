package repositories

import (
	"context"
	"task-management-api/internal/domain/models"
)

type BoardQueryRepositoryI interface {
	GetAll(ctx context.Context) ([]*models.Board, error)
	GetById(ctx context.Context, id string) (*models.Board, error)
}
