package eventsourcing

import (
	"context"
)

// AggregateSaver save a aggregate.
//
// Usually it stores aggregates to a projected table.
type AggregateSaver interface {
	Save(ctx context.Context, aggregate Aggregate) error
}
