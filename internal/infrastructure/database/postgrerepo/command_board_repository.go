package postgrerepo

import (
	"context"
	"database/sql"
	"fmt"
	"task-management-api/internal/domain/models"
	"task-management-api/internal/domain/repositories"
)

// Esta es la implementacion concreta del contrato
type PostgreSQLBoardCommandRepository struct {
	db *sql.DB
}

// Importante: Las implementaciones concretas siempre deben retornar los contratos(interfaces)
func NewPostgreSQLBoardCommandRepository(db *sql.DB) repositories.BoardCommandRepositoryI {
	return &PostgreSQLBoardCommandRepository{
		db: db,
	}
}

func (r *PostgreSQLBoardCommandRepository) Create(ctx context.Context, board *models.Board) error {
	query := `
        INSERT INTO tableros (nombre, descripcion)
        VALUES ($1, $2)
    `

	_, err := r.db.ExecContext(
		ctx,
		query,
		board.Name,
		board.Description,
	)

	if err != nil {
		return fmt.Errorf("error creating board: %v", err)
	}

	return nil
}

func (r *PostgreSQLBoardCommandRepository) Update(ctx context.Context, board *models.Board) error {
	query := `
        UPDATE tableros 
        SET name = $1,
            description = $2,
            updated_at = $3
        WHERE id = $4
    `

	result, err := r.db.ExecContext(
		ctx,
		query,
		board.Name,
		board.Description,
		board.ID,
	)

	if err != nil {
		return fmt.Errorf("error updating board: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("board with id %s not found", board.ID)
	}

	return nil
}

func (r *PostgreSQLBoardCommandRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM tableros WHERE id = $1`

	result, err := r.db.ExecContext(ctx, query, id)
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

	return nil
}
