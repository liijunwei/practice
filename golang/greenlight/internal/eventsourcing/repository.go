package eventsourcing

import (
	"context"

	"github.com/google/uuid"
)

type AggregateRepository interface {
	Load(ctx context.Context, id uuid.UUID) (Aggregate, error)
	Save(ctx context.Context, aggregate Aggregate) error
}
