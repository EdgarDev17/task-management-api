package services

import (
	"context"
	"task-management-api/internal/domain/models"
	"task-management-api/internal/domain/repositories"
)

type TaskQueryServiceI interface {
	GetAll(ctx context.Context) ([]*models.TaskQuery, error)
	GetById(ctx context.Context, id string) (*models.TaskQuery, error)
}

// Implementacion concreta del servicio
type taskQueryServiceImpl struct {
	repo   repositories.TaskQueryRepositoryI
	logger repositories.Logger
}

func NewTaskQueryServiceImpl(repo repositories.TaskQueryRepositoryI, logger repositories.Logger) TaskQueryServiceI {
	return &taskQueryServiceImpl{
		repo:   repo,
		logger: logger,
	}
}

// Metodo encargado del retornar el Board desde la base de datos a traves del ID
func (service *taskQueryServiceImpl) GetById(ctx context.Context, id string) (*models.TaskQuery, error) {
	data, err := service.repo.GetById(ctx, id)

	if err != nil {
		return nil, err
	}
	return data, nil
}

func (service *taskQueryServiceImpl) GetAll(ctx context.Context) ([]*models.TaskQuery, error) {
	data, err := service.repo.GetAll(ctx)

	if err != nil {
		return nil, err
	}

	return data, nil
}
