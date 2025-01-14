package services

import (
	"context"
	"task-management-api/internal/domain/models"
	"task-management-api/internal/domain/repositories"

	"go.uber.org/zap"
)

type TaskCommandServiceI interface {
	Create(ctx context.Context, task *models.Task) error
	Update(ctx context.Context, task *models.Task) error
	Delete(ctx context.Context, id string) error
}

type taskCommandServiceImpl struct {
	repo   repositories.TaskCommandRepository
	logger repositories.Logger
}

func NewTaskCommandService(repo repositories.TaskCommandRepository, logger repositories.Logger) TaskCommandServiceI {
	return &taskCommandServiceImpl{
		repo:   repo,
		logger: logger,
	}
}

func (s *taskCommandServiceImpl) Create(ctx context.Context, task *models.Task) error {
	err := s.repo.Create(ctx, task)

	if err != nil {
		s.logger.Error("Error en el service command Create()", zap.Error(err))
		return err
	}

	return nil
}

func (s *taskCommandServiceImpl) Update(ctx context.Context, task *models.Task) error {
	err := s.repo.Update(ctx, task)

	if err != nil {
		s.logger.Error("Error en el service command Create()", zap.Error(err))
		return err
	}

	return nil
}

func (s *taskCommandServiceImpl) Delete(ctx context.Context, id string) error {
	err := s.repo.Delete(ctx, id)

	if err != nil {
		s.logger.Error("Error en el service command Create()", zap.Error(err))
		return err
	}

	return nil
}
