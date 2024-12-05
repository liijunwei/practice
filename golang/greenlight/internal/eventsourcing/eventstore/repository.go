package eventstore

import (
	"context"
	"reflect"

	"greenlight/internal/eventsourcing"

	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AggregateRepository struct {
	aggregateType   reflect.Type
	eventStore      EventStore
	aggregateLoader eventsourcing.AggregateLoader
	aggregateSaver  eventsourcing.AggregateSaver
}

func NewAggregateRepository(
	aggregate eventsourcing.Aggregate,
	dbPool *pgxpool.Pool,
	aggregateLoader eventsourcing.AggregateLoader,
	aggregateSaver eventsourcing.AggregateSaver,
) *AggregateRepository {
	eventRegistry := eventsourcing.NewEventRegistryFromStateMachine(aggregate)
	eventStore := newEventStore(eventRegistry, dbPool)

	if err := eventStore.migrate(); err != nil {
		panic(err)
	}

	repo := &AggregateRepository{
		aggregateType:   reflect.TypeOf(aggregate).Elem(),
		eventStore:      eventStore,
		aggregateLoader: aggregateLoader,
		aggregateSaver:  aggregateSaver,
	}

	return repo
}

// TODO support load from event store
// TODO support load from snapshot(view)
func (ar *AggregateRepository) Load(
	ctx context.Context,
	aggregateID uuid.UUID,
) (eventsourcing.Aggregate, error) {
	if ar.aggregateLoader != nil {
		return ar.loadFromAggregateLoader(ctx, aggregateID)
	}

	return ar.loadFromEventStore(ctx, aggregateID)
}

func (ar *AggregateRepository) loadFromAggregateLoader(
	ctx context.Context,
	aggregateID uuid.UUID,
) (eventsourcing.Aggregate, error) {
	aggregate, err := ar.aggregateLoader.Load(ctx, aggregateID)
	if err != nil {
		return nil, &AggregateLoaderError{
			err:         err,
			aggregateID: aggregateID,
		}
	}

	return aggregate, nil
}

func (ar *AggregateRepository) loadFromEventStore(
	ctx context.Context,
	aggregateID uuid.UUID,
) (eventsourcing.Aggregate, error) {
	aggregate := ar.NewEmptyAggregate()
	aggregate.SetAggregateID(aggregateID)

	events, err := ar.eventStore.Load(ctx, aggregateID, 0)
	if err != nil {
		return aggregate, err
	}

	if len(events) == 0 {
		return aggregate, &AggregateNotFoundError{
			err:         err,
			aggregateID: aggregateID,
		}
	}

	for _, event := range events {
		if err = aggregate.Apply(event); err != nil {
			return aggregate, &ApplyEventError{err: err, aggregateID: aggregateID}
		}
	}

	return aggregate, nil
}

func (ar *AggregateRepository) Save(ctx context.Context, aggregate eventsourcing.Aggregate) error {
	return Transaction(ctx, ar.eventStore.dbPool, func(ctx context.Context, _ pgx.Tx) error {
		changes := aggregate.GetChanges()

		if err := ar.eventStore.Append(ctx, aggregate.DomainName(), changes); err != nil {
			return err
		}

		if ar.aggregateSaver != nil {
			if err := ar.aggregateSaver.Save(ctx, aggregate); err != nil {
				return AggregateSaverError{
					err:       err,
					aggregate: aggregate,
				}
			}
		}

		return nil
	})
}

func (ar *AggregateRepository) NewEmptyAggregate() eventsourcing.Aggregate {
	aggregate, _ := reflect.New(ar.aggregateType).Interface().(eventsourcing.Aggregate)

	return aggregate
}
