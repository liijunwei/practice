package domain

import (
	"context"
	"fmt"

	"greenlight/internal/eventsourcing/db"

	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DebitHoldRepository struct {
	repo *db.AggregateRepository
}

func NewDebitHoldRepository(dbPool *pgxpool.Pool) *DebitHoldRepository {
	return &DebitHoldRepository{
		repo: db.NewAggregateRepository(&DebitHold{}, dbPool, nil, nil),
	}
}

func (dhr *DebitHoldRepository) Load(ctx context.Context, id uuid.UUID) (*DebitHold, error) {
	aggregate, err := dhr.repo.Load(ctx, id)
	if err != nil {
		return nil, err
	}

	debitRequest, ok := aggregate.(*DebitHold)
	if !ok {
		return nil, &TypeMismatchError{got: aggregate, expect: &DebitHold{}}
	}

	return debitRequest, nil
}

func (dhr *DebitHoldRepository) Save(ctx context.Context, debitRequest *DebitHold) error {
	err := dhr.repo.Save(ctx, debitRequest)
	if err != nil {
		return fmt.Errorf("failed to save debit hold: %w", err)
	}

	return nil
}
