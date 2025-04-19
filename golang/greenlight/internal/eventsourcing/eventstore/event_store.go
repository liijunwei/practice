package eventstore

import (
	"context"
	"errors"
	"fmt"
	"time"

	"greenlight/internal/eventsourcing"

	"github.com/gofrs/uuid"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

const listEventsByAggregateIDAndVersion = `
SELECT aggregate_id,
       version,
       parent_id,
       event_type,
       payload,
       created_at
FROM events
WHERE aggregate_id = $1 and version >= $2
ORDER by version asc
`

type EventStore struct {
	registry *eventsourcing.EventRegistry
	dbPool   *pgxpool.Pool
}

func newEventStore(er *eventsourcing.EventRegistry, dbPool *pgxpool.Pool) EventStore {
	store := EventStore{}
	store.registry = er
	store.dbPool = dbPool

	return store
}

func (es *EventStore) Load(
	ctx context.Context, aggregateID uuid.UUID, startVersion int,
) ([]eventsourcing.Event, error) {
	events := []eventsourcing.Event{}

	err := Transaction(ctx, es.dbPool, func(ctx context.Context, tx pgx.Tx) error {
		sql := listEventsByAggregateIDAndVersion

		rows, err := tx.Query(ctx, sql, aggregateID, startVersion)
		if err != nil {
			zerolog.Ctx(ctx).Debug().Err(err).
				Str("aggregate_id", aggregateID.String()).
				Msg("load events")

			return fmt.Errorf("load events failed: %w", err)
		}
		defer rows.Close()

		for rows.Next() {
			var evModel EventModel

			if err := rows.Scan(&evModel.AggregateID,
				&evModel.Version,
				&evModel.ParentID,
				&evModel.EventType,
				&evModel.Payload,
				&evModel.CreatedAt); err != nil {
				zerolog.Ctx(ctx).Warn().Err(err).
					Str("aggregate_id", aggregateID.String()).
					Int("start_version", startVersion).
					Msg("scan model error")

				return fmt.Errorf("scan event model error: %w", err)
			}

			event, err := evModel.ToEvent(es.registry)
			if err != nil {
				zerolog.Ctx(ctx).Warn().Err(err).
					Str("event_type", evModel.EventType).
					Msg("failed converting event model to event")

				return err
			}

			events = append(events, event)
		}

		return nil
	})

	return events, err
}

const insertEvent = `
INSERT INTO events (aggregate_type, aggregate_id, version, parent_id, event_type, payload, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7)
`

func (es *EventStore) Append(ctx context.Context, domainName string, events []eventsourcing.Event) error {
	return Transaction(ctx, es.dbPool, func(ctx context.Context, pgtx pgx.Tx) error {

		for _, event := range events {
			evModel, err := buildEventModelFromEventPayload(event)
			if err != nil {
				zerolog.Ctx(ctx).Error().Err(err).Msg("failed to marshal event")

				return err
			}

			_, err = pgtx.Exec(ctx, insertEvent,
				domainName,
				evModel.AggregateID,
				evModel.Version,
				evModel.ParentID,
				evModel.EventType,
				evModel.Payload,
				evModel.CreatedAt)
			if err != nil {
				var pgErr *pgconn.PgError
				if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
					return &EventVersionConflictError{
						err:   err,
						event: event,
					}
				}

				zerolog.Ctx(ctx).Error().Err(err).
					Str("aggregate_id", event.GetAggregateID().String()).
					Str("event_type", string(event.EventType())).
					Interface("event", event).
					Msg("insert event error")

				return fmt.Errorf("failed to insert eventstore event error: %w", err)
			}
		}

		return nil
	})
}

func (es *EventStore) migrate() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	const sql = `
	CREATE TABLE IF NOT EXISTS events (
		aggregate_type VARCHAR (50) NOT NULL,
		aggregate_id uuid NOT NULL,
		version int NOT NULL,
		parent_id uuid NOT NULL,
		event_type VARCHAR (50) NOT NULL,
		payload jsonb NOT NULL,
		created_at timestamp without time zone NOT NULL,
		PRIMARY KEY (aggregate_id, version)
	);
	`

	if _, err := es.dbPool.Exec(ctx, sql); err != nil {
		return fmt.Errorf("failed to create event table: %w", err)
	}

	fmt.Println("event table db migration done")

	return nil
}
