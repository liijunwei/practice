package eventsourcing

import (
	"context"

	"github.com/gofrs/uuid"
)

// AggregateLoader loads a aggregate by id.
//
// Usually it loads aggregates from a projected table.
type AggregateLoader interface {
	Load(ctx context.Context, id uuid.UUID) (Aggregate, error)
}
