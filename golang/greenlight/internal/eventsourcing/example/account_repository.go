package main

import (
	"context"
	"fmt"

	"greenlight/internal/eventsourcing"
	"greenlight/internal/eventsourcing/db"
	"greenlight/internal/eventsourcing/example/sqlcquery"

	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// AccountRepository is just a wrapper around AggregateRepository.
type AccountRepository struct {
	dbPool  *pgxpool.Pool
	repo    *db.AggregateRepository
	queries *sqlcquery.Queries
}

func NewAccountRepository(dbPool *pgxpool.Pool) *AccountRepository {
	// aggregate loader && saver
	loaderSaver := &AccountLoaderSaver{
		queries: sqlcquery.New(dbPool),
	}

	return &AccountRepository{
		dbPool:  dbPool,
		repo:    db.NewAggregateRepository(&Account{}, dbPool, loaderSaver, loaderSaver),
		queries: sqlcquery.New(dbPool),
	}
}

// Load loads a Account from database.
func (ar *AccountRepository) Load(ctx context.Context, id uuid.UUID) (*Account, error) {
	aggregate, err := ar.repo.Load(ctx, id)
	if err != nil {
		return nil, err //nolint: wrapcheck // example
	}

	account, ok := aggregate.(*Account)
	if !ok {
		return nil, &TypeMismatchError{expect: &Account{}, got: aggregate}
	}

	return account, nil
}

// Save saves a Account.
func (ar *AccountRepository) Save(ctx context.Context, account *Account) error {
	err := ar.repo.Save(ctx, account)
	if err != nil {
		return fmt.Errorf("failed to save account: %w", err)
	}

	return nil
}

// LoadLocked loads a Account from database, it also locked the row, forbiding others
// to load it.
func (ar *AccountRepository) LoadLocked(ctx context.Context, accountID uuid.UUID) (*Account, error) {
	var (
		trans pgx.Tx
		err   error
	)

	// get existing transaction or start a new one
	trans, ok := db.GetTx(ctx)
	if !ok {
		trans, err = ar.dbPool.Begin(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to start transaction %w", err)
		}
	}

	queries := ar.queries.WithTx(trans)

	account, err := queries.GetAccountByIDLocked(ctx, accountID)
	if err != nil {
		return nil, err //nolint: wrapcheck // example
	}

	return fromSQLCAccount(account), nil
}

func fromSQLCAccount(sqlcAccount *sqlcquery.Account) *Account {
	account := &Account{
		Balance:   sqlcAccount.Balance,
		Available: sqlcAccount.Available,
		Pending:   sqlcAccount.Pending,
		CreatedAt: sqlcAccount.CreatedAt,
		UpdatedAt: sqlcAccount.UpdatedAt,
	}
	account.SetAggregateID(sqlcAccount.ID)
	account.SetVersion(int(sqlcAccount.Version))

	return account
}

type AccountLoaderSaver struct {
	queries *sqlcquery.Queries
}

func (alc *AccountLoaderSaver) Load( //nolint: ireturn // AggregateLoader interface requires it
	ctx context.Context, accountID uuid.UUID,
) (eventsourcing.Aggregate, error) {
	queries := alc.queries

	// if inside a transaction
	if trans, ok := db.GetTx(ctx); ok {
		queries = alc.queries.WithTx(trans)
	}

	account, err := queries.GetAccountByID(ctx, accountID)
	if err != nil {
		return nil, err //nolint: wrapcheck // example
	}

	return fromSQLCAccount(account), nil
}

func (alc *AccountLoaderSaver) Save(ctx context.Context, aggregate eventsourcing.Aggregate) error {
	queries := alc.queries

	// if inside a transaction
	if trans, ok := db.GetTx(ctx); ok {
		queries = alc.queries.WithTx(trans)
	}

	account, ok := aggregate.(*Account)
	if !ok {
		return &TypeMismatchError{
			expect: &Account{},
			got:    aggregate,
		}
	}

	err := queries.UpsertAccount(ctx, &sqlcquery.UpsertAccountParams{
		ID:        account.ID,
		Balance:   account.Balance,
		Available: account.Available,
		Pending:   account.Pending,
		Version:   int32(account.Version),
		UpdatedAt: account.UpdatedAt,
		CreatedAt: account.CreatedAt,
	})

	return err //nolint: wrapcheck // example
}
