package repositories

import (
	"context"
	"task-management-api/internal/domain/models"
)

type TaskQueryRepositoryI interface {
	GetAll(ctx context.Context) ([]*models.TaskQuery, error)
	GetById(ctx context.Context, id string) (*models.TaskQuery, error)
	Upsert(ctx context.Context, board *models.Task) error
	Delete(ctx context.Context, board *models.Task) error
}
