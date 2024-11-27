package main

import (
	"context"
	"fmt"

	"greenlight/internal/eventsourcing/db"

	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

// DebitHoldRepository is just a wrapper around AggregateRepository.
//
// Since event DebitRequest has at-most two events, it's cheap to loading data from event store.
type DebitHoldRepository struct {
	repo *db.AggregateRepository
}

func NewDebitHoldRepository(dbPool *pgxpool.Pool) *DebitHoldRepository {
	return &DebitHoldRepository{
		repo: db.NewAggregateRepository(&DebitHold{}, dbPool, nil, nil),
	}
}

// Load loads a DebitRequest from event store.
//
// AggregateRepository's Load method return a Aggregate interface, this method converts
// it to actual DebitRequest.
func (dhr *DebitHoldRepository) Load(ctx context.Context, id uuid.UUID) (*DebitHold, error) {
	aggregate, err := dhr.repo.Load(ctx, id)
	if err != nil {
		return nil, err //nolint: wrapcheck // example
	}

	debitRequest, ok := aggregate.(*DebitHold)
	if !ok {
		return nil, &TypeMismatchError{got: aggregate, expect: &DebitHold{}}
	}

	return debitRequest, nil
}

// Save saves a DebitRequest.
func (dhr *DebitHoldRepository) Save(ctx context.Context, debitRequest *DebitHold) error {
	err := dhr.repo.Save(ctx, debitRequest)
	if err != nil {
		return fmt.Errorf("failed to save debit hold: %w", err)
	}

	return nil
}
