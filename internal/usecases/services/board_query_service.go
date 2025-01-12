package services

import (
	"context"

	"task-management-api/internal/domain/models"
	"task-management-api/internal/domain/repositories"
)

// Interface para el servicio
type BoardQueryServiceI interface {
	GetAll(ctx context.Context) ([]*models.Board, error)
	GetById(ctx context.Context, id string) (*models.Board, error)
}

// Implementacion concreta del servicio
type BoardQueryServiceImpl struct {
	repo repositories.BoardQueryRepositoryI
}

// Constructor del servicio
func NewBoardQueryService(repo repositories.BoardQueryRepositoryI) BoardQueryServiceI {
	return &BoardQueryServiceImpl{
		repo: repo,
	}
}

// Metodo encargado del retornar el Board desde la base de datos a traves del ID
func (service *BoardQueryServiceImpl) GetById(ctx context.Context, id string) (*models.Board, error) {
	data, err := service.repo.GetById(ctx, id)

	if err != nil {
		return nil, err
	}
	return data, nil
}

func (service *BoardQueryServiceImpl) GetAll(ctx context.Context) ([]*models.Board, error) {
	data, err := service.repo.GetAll(ctx)

	if err != nil {
		return nil, err
	}

	return data, nil
}
