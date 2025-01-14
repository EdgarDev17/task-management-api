package repositories

import (
	"context"
	"task-management-api/internal/domain/models"

	"github.com/google/uuid"
)

// Este contrato se encarga de definir los metodos que debe seguir
// cualquier base de datos que se utilize para escritura (command db)
type TaskCommandRepository interface {
	Create(ctx context.Context, task *models.Task) error
	Update(ctx context.Context, task *models.Task) error
	Delete(ctx context.Context, id string) error
	UpdateState(ctx context.Context, taskID uuid.UUID, boardID uuid.UUID, newState string) error
}
