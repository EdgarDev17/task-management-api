package events

import (
	"context"
	"task-management-api/internal/domain/repositories"
	"time"

	"go.uber.org/zap"
)

type EventProcessor struct {
	handlers []repositories.EventHandlerI
	ticker   *time.Ticker
	done     chan bool
	logger   repositories.Logger
}

func NewEventProcessor(handlers []repositories.EventHandlerI, logger repositories.Logger) *EventProcessor {
	return &EventProcessor{
		handlers: handlers,
		ticker:   time.NewTicker(5 * time.Second),
		done:     make(chan bool),
		logger:   logger,
	}
}

func (ep *EventProcessor) Start() {
	go func() {
		for {
			select {
			case <-ep.done:
				return
			case <-ep.ticker.C:
				ctx := context.Background()
				for _, handler := range ep.handlers {
					if err := handler.ProcessEvents(ctx); err != nil {
						ep.logger.Error("Error processing events", zap.Error(err))
					}
				}
			}
		}
	}()
}
