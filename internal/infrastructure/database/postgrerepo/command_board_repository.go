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

// Esta es la implementacion concreta del contrato
type PostgreSQLBoardCommandRepository struct {
	db         *sql.DB
	eventStore repositories.EventStore
	logger     repositories.Logger
}

// Importante: Las implementaciones concretas siempre deben retornar los contratos(interfaces)
func NewPostgreSQLBoardCommandRepository(db *sql.DB, eventStore repositories.EventStore, logger repositories.Logger) repositories.BoardCommandRepositoryI {
	return &PostgreSQLBoardCommandRepository{
		db:         db,
		eventStore: eventStore,
		logger:     logger,
	}
}

func (r *PostgreSQLBoardCommandRepository) Create(ctx context.Context, board *models.Board) (*models.Board, error) {
	query := `
	INSERT INTO tableros (nombre, descripcion)
	VALUES ($1, $2)
	RETURNING id
		`

	tx, err := r.db.BeginTx(ctx, nil)

	if err != nil {
		return nil, fmt.Errorf("error creating transaction: %v", err)
	}

	var boardID uuid.UUID

	err = tx.QueryRowContext(
		ctx,
		query,
		board.Name,
		board.Description,
	).Scan(&boardID)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// Asigno el ID obtenido al board
	board.ID = boardID

	// Inicializo el evento
	event := &events.BoardCreatedEvent{
		BaseEvent: events.BaseEvent{
			ID:        uuid.New(),
			Timestamp: time.Now(),
			Type:      "BoardCreated",
		},
		Name:        board.Name,
		Description: board.Description,
		BoardID:     boardID,
	}

	// guardo el evento en su respectiva tabla
	err = r.eventStore.SaveEvent(ctx, event)

	if err != nil {
		tx.Rollback()
		r.logger.Error("Error al guardar evento", zap.Error(err))
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return board, nil
}

func (r *PostgreSQLBoardCommandRepository) Update(ctx context.Context, board *models.Board) error {

	fmt.Println("id a modificar", board.ID)
	query := `
        UPDATE tableros 
        SET nombre = $1,
            descripcion = $2
        WHERE id = $3
    `

	tx, err := r.db.BeginTx(ctx, nil)

	if err != nil {
		return fmt.Errorf("error creating transaction: %v", err)
	}

	result, err := tx.ExecContext(
		ctx,
		query,
		board.Name,
		board.Description,
		board.ID,
	)

	if err != nil {
		tx.Rollback()
		return err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return fmt.Errorf("error getting rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("board with id %s not found", board.ID)
	}

	// Inicializo el evento
	event := &events.BoardUpdatedEvent{
		BaseEvent: events.BaseEvent{
			ID:        uuid.New(),
			Timestamp: time.Now(),
			Type:      "BoardUpdated",
		},
		BoardID:     board.ID,
		Name:        board.Name,
		Description: board.Description,
	}

	// guardo el evento en su respectiva tabla
	err = r.eventStore.SaveEvent(ctx, event)

	if err != nil {
		tx.Rollback()
		r.logger.Error("Error al guardar evento", zap.Error(err))
		return err
	}

	return tx.Commit()
}

func (r *PostgreSQLBoardCommandRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM tableros WHERE id = $1`

	tx, err := r.db.BeginTx(ctx, nil)

	if err != nil {
		return fmt.Errorf("error creating transaction: %v", err)
	}

	result, err := tx.ExecContext(ctx, query, id)

	if err != nil {
		return fmt.Errorf("error deleting board: %v", err)
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return fmt.Errorf("error getting rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("board with id %s not found", id)
	}

	// Inicializo el evento
	boardUUID, err := uuid.Parse(id)

	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error parsing UUID: %v", err)
	}

	event := &events.BoardDeletedEvent{
		BaseEvent: events.BaseEvent{
			ID:        uuid.New(),
			Timestamp: time.Now(),
			Type:      "BoardDeleted",
		},
		BoardID: boardUUID,
	}

	// guardo el evento en su respectiva tabla
	err = r.eventStore.SaveEvent(ctx, event)

	if err != nil {
		tx.Rollback()
		r.logger.Error("Error al guardar evento", zap.Error(err))
		return err
	}

	return tx.Commit()
}
