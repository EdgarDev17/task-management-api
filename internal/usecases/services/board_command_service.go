package services

import (
	"context"
	"fmt"
	"task-management-api/internal/domain/models"
	"task-management-api/internal/domain/repositories"
)

// interfaz para el servicio
type BoardCommandServiceI interface {
	Create(ctx context.Context, board *models.Board) error
	Update(ctx context.Context, board *models.Board) error
	Delete(ctx context.Context, id string) error
}

type boardCommandServiceImpl struct {
	repo repositories.BoardCommandRepositoryI
}

func NewBoardCommandService(repo repositories.BoardCommandRepositoryI) BoardCommandServiceI {
	return &boardCommandServiceImpl{
		repo: repo,
	}
}

func (s *boardCommandServiceImpl) Create(ctx context.Context, board *models.Board) error {
	err := s.repo.Create(ctx, board)

	if err != nil {
		fmt.Println("error en el servicio postgresql ", err)
		return err
	}

	return nil
}

func (s *boardCommandServiceImpl) Update(ctx context.Context, board *models.Board) error {
	err := s.repo.Update(ctx, board)

	if err != nil {
		fmt.Println("error en el servicio postgresql ", err)
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
