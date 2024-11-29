package db

import (
	"context"
	"errors"
	"fmt"

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

const insertEvent = `
INSERT INTO events (aggregate_id, version, parent_id, event_type, payload, created_at) VALUES ($1, $2, $3, $4, $5, $6)
`

type EventStore struct {
	registry *eventsourcing.EventRegistry
	dbPool   *pgxpool.Pool
}

func NewEventStore(er *eventsourcing.EventRegistry, dbPool *pgxpool.Pool) EventStore {
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

func (es *EventStore) Append(ctx context.Context, events []eventsourcing.Event) error {
	return Transaction(ctx, es.dbPool, func(ctx context.Context, pgtx pgx.Tx) error {
		sql := insertEvent

		for _, event := range events {
			evModel, err := NewEventModelFromEvent(event)
			if err != nil {
				zerolog.Ctx(ctx).Warn().Err(err).Msg("failed to marshal event")

				return err
			}

			_, err = pgtx.Exec(ctx, sql,
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

				zerolog.Ctx(ctx).Warn().Err(err).
					Str("sql", sql).
					Str("aggregate_id", event.GetAggregateID().String()).
					Str("event_type", string(event.EventType())).
					Msg("insert event error")

				return fmt.Errorf("insert event<<< %t error %w", event, err)
			}
		}

		return nil
	})
}
