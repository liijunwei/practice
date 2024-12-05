package eventstore

import (
	"fmt"

	"greenlight/internal/eventsourcing"

	"github.com/gofrs/uuid"
	"github.com/rs/zerolog/log"
)

type AggregateNotFoundError struct {
	err         error
	aggregateID uuid.UUID
}

func (anfe AggregateNotFoundError) Error() string {
	return fmt.Sprintf("aggregate not found: %s (id %s)", anfe.err, anfe.aggregateID.String())
}

type LoadEventError struct {
	err error
}

func (lee LoadEventError) Error() string {
	return fmt.Sprintf("failed to load event: %s", lee.err)
}

type ApplyEventError struct {
	err         error
	aggregateID uuid.UUID
}

func (aee ApplyEventError) Error() string {
	return fmt.Sprintf("apply event error: %s (aggregate_id %s)", aee.err, aee.aggregateID.String())
}

type EventPublishError struct {
	err   error
	event eventsourcing.Event
}

func (epe EventPublishError) Error() string {
	return fmt.Sprintf("failed to publish event : %s (aggregate_id %s, event %s)",
		epe.err, epe.event.GetAggregateID().String(), epe.event.EventType())
}

type EventProjectorError struct {
	err   error
	event eventsourcing.Event
}

func (epe EventProjectorError) Error() string {
	return fmt.Sprintf("failed to project event : %s (aggregate_id %s, event %s)",
		epe.err, epe.event.GetAggregateID().String(), epe.event.EventType())
}

type AggregateSaverError struct {
	err       error
	aggregate eventsourcing.Aggregate
}

func (ale AggregateSaverError) Error() string {
	log.Error().
		Interface("changes", ale.aggregate.GetChanges()).
		Str("aggregate_id", ale.aggregate.GetAggregateID().String()).
		Msg("failed on AggregateSaverError")

	return fmt.Sprintf("failed to save aggregate, aggregate_id: %s, error: %s",
		ale.aggregate.GetAggregateID().String(), ale.err)
}

func (ale AggregateSaverError) Unwrap() error {
	return ale.err
}

type AggregateLoaderError struct {
	err         error
	aggregateID uuid.UUID
}

func (ale AggregateLoaderError) Unwrap() error {
	return ale.err
}

func (ale AggregateLoaderError) Error() string {
	return fmt.Sprintf("failed to load aggregate: %s (aggregate_id %s)",
		ale.err, ale.aggregateID)
}

type TransactionError struct {
	err error
}

func (te TransactionError) Error() string {
	return fmt.Sprintf("transaction error: %v", te.err)
}

func (te TransactionError) Unwrap() error {
	return te.err
}

type EventUnmarshalError struct {
	err        error
	eventModel *EventModel
}

func (eue EventUnmarshalError) Error() string {
	return fmt.Sprintf("failed to unmarshal event: %s (event type %s, aggregate_id %s)",
		eue.err, eue.eventModel.EventType, eue.eventModel.AggregateID)
}

type EventMarshalError struct {
	err   error
	event eventsourcing.Event
}

func (eue EventMarshalError) Error() string {
	return fmt.Sprintf("failed to marshal event: %s (event type %s, aggregate_id %s)",
		eue.err, eue.event.EventType(), eue.event.GetAggregateID())
}

type EventVersionConflictError struct {
	err   error
	event eventsourcing.Event
}

func (cee *EventVersionConflictError) Error() string {
	return fmt.Sprintf("event version conflicted: %s (event type %s, aggregate_id %s)",
		cee.err, cee.event.EventType(), cee.event.GetAggregateID())
}
