package repositories

import (
	"context"
	"task-management-api/internal/domain/models"
)

// Este es un contrato que define las operaciones necesarias que toda Store debe seguir
type EventStore interface {
	//Este guarda el evento en la base de datos
	//
	// @param event: El evento que se debe almacenar. Este evento debe cumplir con el contrato de la interfaz Event.
	SaveEvent(ctx context.Context, event Event) error

	// Retorna una lista de eventos
	//
	// @param eventState: El estado de procesamiento de los eventos que se desean recuperar.
	GetEvents(ctx context.Context, eventState models.ProcessingState) ([]Event, error)

	//Este guarda el evento en la base de datos
	GetLastProcessingState(ctx context.Context) (models.ProcessingState, error)

	//Actualiza el evento ya procesado
	InsertProcessingState(ctx context.Context, state *models.ProcessingState) error
}
