package eventsourcing

import (
	"context"

	"github.com/gofrs/uuid"
)

type EventStore[T Aggregate] interface {
	Load(ctx context.Context, aggregateID uuid.UUID, startVersion int) ([]Event, error)

	Append(ctx context.Context, events []Event) error
}
