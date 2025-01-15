package services

import (
	"context"
	"fmt"
	"task-management-api/internal/domain/models"
	"task-management-api/internal/domain/repositories"

	"go.uber.org/zap"
)

// interfaz para el servicio
type BoardCommandServiceI interface {
	Create(ctx context.Context, board *models.Board) (*models.Board, error)
	Update(ctx context.Context, board *models.Board) error
	Delete(ctx context.Context, id string) error
}

type boardCommandServiceImpl struct {
	repo   repositories.BoardCommandRepositoryI
	logger repositories.Logger
}

func NewBoardCommandService(repo repositories.BoardCommandRepositoryI, logger repositories.Logger) BoardCommandServiceI {
	return &boardCommandServiceImpl{
		repo:   repo,
		logger: logger,
	}
}

func (s *boardCommandServiceImpl) Create(ctx context.Context, board *models.Board) (*models.Board, error) {
	newBoard, err := s.repo.Create(ctx, board)

	if err != nil {
		s.logger.Error("Error en el service command Create()", zap.Error(err))
		return nil, err
	}

	return newBoard, nil
}

func (s *boardCommandServiceImpl) Update(ctx context.Context, board *models.Board) error {
	err := s.repo.Update(ctx, board)

	if err != nil {
		s.logger.Error("Error en el service command Update()", zap.Error(err))
		return err
	}

	return nil
}

func (s *boardCommandServiceImpl) Delete(ctx context.Context, id string) error {
	err := s.repo.Delete(ctx, id)

	if err != nil {
		fmt.Println("error en el servicio postgresql ", err)
		return err
	}

	return nil
}
