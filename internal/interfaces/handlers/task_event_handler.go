package handlers

import (
	"context"
	"fmt"
	"log"
	"task-management-api/internal/domain/models"
	"task-management-api/internal/domain/repositories"
	"task-management-api/internal/infrastructure/events"
	"time"

	"go.uber.org/zap"
)

type taskEventHandler struct {
	eventStore      repositories.EventStore
	logger          repositories.Logger
	queryRepository repositories.TaskQueryRepositoryI
	ticker          *time.Ticker
	done            chan bool
}

func NewTaskEventHandler(
	eventStore repositories.EventStore,
	queryRepository repositories.TaskQueryRepositoryI,
	logger repositories.Logger,
) *taskEventHandler {
	return &taskEventHandler{
		eventStore:      eventStore,
		queryRepository: queryRepository,
		ticker:          time.NewTicker(5 * time.Second),
		done:            make(chan bool),
		logger:          logger,
	}
}

func (h *taskEventHandler) ProcessEvents(ctx context.Context) error {
	state, err := h.eventStore.GetLastProcessingState(ctx)

	if err != nil {
		h.logger.Error("Error al obtener el eltimo evento procesado", zap.Error(err))
		return err
	}

	eventsList, err := h.eventStore.GetEvents(ctx, state)

	if err != nil {
		h.logger.Error("Error al obtener eventos en Tasks", zap.Error(err))
		return err
	}

	if len(eventsList) == 0 {
		return nil
	}

	for _, event := range eventsList {
		var err error

		// Procesamiento solo para creación de Board
		switch e := event.(type) {
		case *events.TaskCreatedEvent:
			readModel := &models.Task{
				ID:        e.TaskID,
				Title:     e.Title,
				BoardID:   e.BoardID,
				State:     e.State,
				CreatedAt: e.GetTimestamp(),
			}

			err = h.queryRepository.Upsert(ctx, readModel)

		case *events.TaskUpdatedEvent:
			readModel := &models.Task{
				ID:          e.TaskID,
				Title:       e.Title,
				BoardID:     e.BoardID,
				State:       e.State,
				Description: e.Description,
			}
			err = h.queryRepository.Upsert(ctx, readModel)

		case *events.TaskDeletedEvent:
			readModel := &models.Task{
				ID: e.TaskID,
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
