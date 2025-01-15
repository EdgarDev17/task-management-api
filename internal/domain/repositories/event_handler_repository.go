package repositories

import "context"

type EventHandlerI interface {
	ProcessEvents(ctx context.Context) error
}
