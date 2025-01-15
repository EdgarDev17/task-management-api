package handlers

import (
	"context"
	"fmt"
	"log"
	"task-management-api/internal/domain/models"
	"task-management-api/internal/domain/repositories"
	"task-management-api/internal/infrastructure/events"
	"time"
)

type EventHandler struct {
	eventStore      repositories.EventStore
	logger          repositories.Logger
	queryRepository repositories.BoardQueryRepositoryI
	ticker          *time.Ticker
	done            chan bool
}

func NewEventHandler(
	eventStore repositories.EventStore,
	queryRepository repositories.BoardQueryRepositoryI,
	logger repositories.Logger,
) *EventHandler {
	return &EventHandler{
		eventStore:      eventStore,
		queryRepository: queryRepository,
		ticker:          time.NewTicker(5 * time.Second),
		done:            make(chan bool),
		logger:          logger,
	}
}

func (h *EventHandler) ProcessEvents(ctx context.Context) error {
	state, err := h.eventStore.GetLastProcessingState(ctx)

	if err != nil {
		log.Printf("Error getting last processing state: %v", err)
		return err
	}

	eventsList, err := h.eventStore.GetEvents(ctx, state)

	if err != nil {
		log.Printf("Error getting events: %v", err)
		return err
	}

	if len(eventsList) == 0 {
		return nil
	}

	for _, event := range eventsList {
		var err error

		// Procesamiento solo para creación de Board
		switch e := event.(type) {
		case *events.BoardCreatedEvent:
			readModel := &models.Board{
				ID:          e.BoardID,
				Name:        e.Name,
				Description: e.Description,
				CreatedAt:   e.GetTimestamp(),
			}

			err = h.queryRepository.Upsert(ctx, readModel)

		case *events.BoardUpdatedEvent:
			readModel := &models.Board{
				ID:          e.BoardID,
				Name:        e.Name,
				Description: e.Description,
			}
			err = h.queryRepository.Upsert(ctx, readModel)

		case *events.BoardDeletedEvent:
			readModel := &models.Board{
				ID: e.BoardID,
			}
			err = h.queryRepository.Delete(ctx, readModel)

		default:
			err = fmt.Errorf("unknown event type: %T", event)
		}

		if err != nil {
			log.Printf("Error processing event %s: %v", event.GetID(), err)
			continue
		}

		// Actualizar el estado con el último evento procesado exitosamente en la base de datos
		lastIDUUID := event.GetID()
		state.LastProcessedID = &lastIDUUID
		state.LastProcessedTimestamp = event.GetTimestamp()

		h.eventStore.InsertProcessingState(ctx, &state)
	}

	return nil
}
