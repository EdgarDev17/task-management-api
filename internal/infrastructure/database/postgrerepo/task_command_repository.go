package postgrerepo

import (
	"context"
	"database/sql"
	"fmt"
	"task-management-api/internal/domain/models"
	"task-management-api/internal/domain/repositories"
	"task-management-api/internal/infrastructure/events"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type taskCommandRepository struct {
	db         *sql.DB
	eventStore repositories.EventStore
	logger     repositories.Logger
}

func NewTaskCommandRepository(db *sql.DB, eventStore repositories.EventStore, logger repositories.Logger) repositories.TaskCommandRepository {
	return &taskCommandRepository{
		db:         db,
		eventStore: eventStore,
		logger:     logger,
	}
}

func (r *taskCommandRepository) Create(ctx context.Context, task *models.Task) error {
	query := `
	INSERT INTO tareas (id_tablero, titulo, descripcion, estado, fecha_creacion)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id`

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("error creating transaction: %v", err)
	}

	taskID := uuid.New()
	task.ID = taskID
	createdAt := time.Now()

	err = tx.QueryRowContext(
		ctx,
		query,
		task.BoardID,
		task.Title,
		task.Description,
		task.State,
		createdAt,
	).Scan(&taskID)

	if err != nil {
		tx.Rollback()
		return err
	}

	event := &events.TaskCreatedEvent{
		BaseEvent: events.BaseEvent{
			ID:        uuid.New(),
			Timestamp: time.Now(),
			Type:      "TaskCreated",
		},
		TaskID:  taskID,
		BoardID: task.BoardID,
		Title:   task.Title,
		State:   task.State,
	}

	err = r.eventStore.SaveEvent(ctx, event)
	if err != nil {
		tx.Rollback()
		r.logger.Error("Error al guardar evento", zap.Error(err))
		return err
	}

	return tx.Commit()
}

func (r *taskCommandRepository) Update(ctx context.Context, task *models.Task) error {
	query := `
        UPDATE tasks 
        SET title = $1,
            description = $2,
            state = $3
        WHERE id = $4 AND board_id = $5`

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("error creating transaction: %v", err)
	}

	result, err := tx.ExecContext(
		ctx,
		query,
		task.Title,
		task.Description,
		task.State,
		task.ID,
		task.BoardID,
	)

	if err != nil {
		tx.Rollback()
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error getting rows affected: %v", err)
	}

	if rowsAffected == 0 {
		tx.Rollback()
		return fmt.Errorf("task with id %s not found", task.ID)
	}

	event := &events.TaskUpdatedEvent{
		BaseEvent: events.BaseEvent{
			ID:        uuid.New(),
			Timestamp: time.Now(),
			Type:      "TaskUpdated",
		},
		TaskID:  task.ID,
		BoardID: task.BoardID,
		Title:   task.Title,
		State:   task.State,
	}

	err = r.eventStore.SaveEvent(ctx, event)
	if err != nil {
		tx.Rollback()
		r.logger.Error("Error al guardar evento", zap.Error(err))
		return err
	}

	return tx.Commit()
}

func (r *taskCommandRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM tasks WHERE id = $1`

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("error creating transaction: %v", err)
	}

	result, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error deleting task: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error getting rows affected: %v", err)
	}

	if rowsAffected == 0 {
		tx.Rollback()
		return fmt.Errorf("task with id %s not found", id)
	}

	taskUUID, err := uuid.Parse(id)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error parsing UUID: %v", err)
	}

	event := &events.TaskDeletedEvent{
		BaseEvent: events.BaseEvent{
			ID:        uuid.New(),
			Timestamp: time.Now(),
			Type:      "TaskDeleted",
		},
		TaskID: taskUUID,
	}

	err = r.eventStore.SaveEvent(ctx, event)
	if err != nil {
		tx.Rollback()
		r.logger.Error("Error al guardar evento", zap.Error(err))
		return err
	}

	return tx.Commit()
}
