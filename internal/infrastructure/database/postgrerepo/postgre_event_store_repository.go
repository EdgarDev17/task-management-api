package postgrerepo

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"task-management-api/internal/domain/models"
	"task-management-api/internal/domain/repositories"
	"task-management-api/internal/infrastructure/events"
	"time"

	"github.com/google/uuid"
)

type PostgresEventStore struct {
	db *sql.DB
}

func NewPostgresEventStore(db *sql.DB) (repositories.EventStore, error) {
	return &PostgresEventStore{db: db}, nil
}

// insertamos el evento en la base de datos posgre
func (s *PostgresEventStore) SaveEvent(ctx context.Context, event repositories.Event) error {
	data, err := json.Marshal(event)

	if err != nil {
		return err
	}

	query := `
        INSERT INTO events (id, aggregate_id, event_type, timestamp, data)
        VALUES ($1, $2, $3, $4, $5)`

	_, err = s.db.ExecContext(ctx, query,
		event.GetID(),
		event.GetID(),
		event.GetType(),
		event.GetTimestamp(),
		data,
	)
	return err
}

func (s *PostgresEventStore) GetEvents(ctx context.Context, state models.ProcessingState) ([]repositories.Event, error) {

	var query string
	var args []interface{}

	if state.LastProcessedID == nil {
		// Primera ejecución: solo usamos timestamp
		query = `
            SELECT id, aggregate_id, event_type, data, timestamp
            FROM events 
            WHERE timestamp >= $1 
            ORDER BY timestamp ASC, id ASC
            LIMIT 100`
		args = []interface{}{state.LastProcessedTimestamp}
	} else {
		// Ejecuciones subsiguientes: usamos ID y timestamp
		query = `
            SELECT id, aggregate_id, event_type, data, timestamp
            FROM events 
            WHERE (timestamp > $1) OR 
                  (timestamp = $1 AND id > $2)
            ORDER BY timestamp ASC, id ASC
            LIMIT 100`
		args = []interface{}{state.LastProcessedTimestamp, *state.LastProcessedID}
	}

	rows, err := s.db.QueryContext(ctx, query, args...)

	if err != nil {
		return nil, fmt.Errorf("error querying events: %w", err)
	}
	defer rows.Close()

	var eventsList []repositories.Event

	for rows.Next() {
		var (
			id          string
			id_agregate string
			eventType   string
			data        []byte
			timestamp   time.Time
		)

		if err := rows.Scan(&id, &id_agregate, &eventType, &data, &timestamp); err != nil {
			return nil, fmt.Errorf("error scanning event row: %w", err)
		}

		// Crear el evento específico según su tipo
		var event repositories.Event
		switch eventType {
		case "BoardCreated":
			var e events.BoardCreatedEvent
			if err := json.Unmarshal(data, &e); err != nil {
				return nil, fmt.Errorf("error unmarshaling ProductCreatedEvent: %w", err)
			}

			// Convertir string a UUID
			idUUID, err := uuid.Parse(id)
			if err != nil {
				return nil, fmt.Errorf("error parsing UUID: %w", err)
			}
			e.ID = idUUID
			e.Timestamp = timestamp
			e.Type = eventType
			event = &e

		case "BoardUpdated":
			var e events.BoardUpdatedEvent
			if err := json.Unmarshal(data, &e); err != nil {
				return nil, fmt.Errorf("error unmarshaling ProductUpdatedEvent: %w", err)
			}
			idUUID, err := uuid.Parse(id)
			if err != nil {
				return nil, fmt.Errorf("error parsing UUID: %w", err)
			}
			e.ID = idUUID
			e.Timestamp = timestamp
			e.Type = eventType
			event = &e

		case "BoardDeleted":
			var e events.BoardDeletedEvent
			if err := json.Unmarshal(data, &e); err != nil {
				return nil, fmt.Errorf("error unmarshaling BoardDeletedEvent: %w", err)
			}
			idUUID, err := uuid.Parse(id)
			if err != nil {
				return nil, fmt.Errorf("error parsing UUID: %w", err)
			}
			e.ID = idUUID
			e.Timestamp = timestamp
			e.Type = eventType
			event = &e

		// Agregar más casos según los tipos de eventos que manejes

		default:
			return nil, fmt.Errorf("unknown event type: %s", eventType)
		}

		eventsList = append(eventsList, event)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating event rows: %w", err)
	}

	return eventsList, nil
}

func (s *PostgresEventStore) GetLastProcessingState(ctx context.Context) (models.ProcessingState, error) {
	query := `
        SELECT last_processed_id, last_processed_timestamp
        FROM event_processing
        ORDER BY id DESC
        LIMIT 1`

	var state models.ProcessingState
	var idNullable sql.NullString

	err := s.db.QueryRowContext(ctx, query).Scan(&idNullable, &state.LastProcessedTimestamp)

	if err == sql.ErrNoRows {
		// Primera ejecución: comenzar desde el inicio del tiempo
		return models.ProcessingState{
			LastProcessedID:        nil,
			LastProcessedTimestamp: time.Unix(0, 0),
		}, nil
	}
	if err != nil {
		return models.ProcessingState{}, fmt.Errorf("error getting last processed state: %w", err)
	}

	if idNullable.Valid {
		id, err := uuid.Parse(idNullable.String)
		if err != nil {
			return models.ProcessingState{}, fmt.Errorf("invalid UUID in database: %w", err)
		}
		state.LastProcessedID = &id
	}

	return state, nil
}

func (s *PostgresEventStore) InsertProcessingState(ctx context.Context, state *models.ProcessingState) error {
	query := `
        INSERT INTO event_processing (last_processed_id, last_processed_timestamp, updated_at)
        VALUES ($1, $2, NOW())`

	_, err := s.db.ExecContext(ctx, query, state.LastProcessedID, state.LastProcessedTimestamp)
	return err
}
