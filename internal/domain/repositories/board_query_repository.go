package repositories

import (
	"context"
	"task-management-api/internal/domain/models"
)

type BoardQueryRepositoryI interface {
	GetAll(ctx context.Context) ([]*models.BoardQuery, error)
	GetById(ctx context.Context, id string) (*models.BoardQuery, error)
	Upsert(ctx context.Context, board *models.Board) error
	Delete(ctx context.Context, board *models.Board) error
}
