package domain

import (
	"context"
	"fmt"

	"greenlight/internal/eventsourcing"
	"greenlight/internal/eventsourcing/eventstore"
	"greenlight/internal/eventsourcing/example/sqlcquery"

	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AccountRepository struct {
	dbPool  *pgxpool.Pool
	repo    *eventstore.AggregateRepository
	queries *sqlcquery.Queries
}

func NewAccountRepository(dbPool *pgxpool.Pool) *AccountRepository {
	loaderSaver := &AccountLoaderSaver{
		queries: sqlcquery.New(dbPool),
	}

	return &AccountRepository{
		dbPool:  dbPool,
		repo:    eventstore.NewAggregateRepository(&Account{}, dbPool, loaderSaver, loaderSaver),
		queries: sqlcquery.New(dbPool),
	}
}

func (ar *AccountRepository) Load(ctx context.Context, id uuid.UUID) (*Account, error) {
	aggregate, err := ar.repo.Load(ctx, id)
	if err != nil {
		return nil, err
	}

	account, ok := aggregate.(*Account)
	if !ok {
		return nil, &TypeMismatchError{expect: &Account{}, got: aggregate}
	}

	return account, nil
}

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

	trans, ok := eventstore.GetTx(ctx)
	if !ok {
		trans, err = ar.dbPool.Begin(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to start transaction %w", err)
		}
	}

	queries := ar.queries.WithTx(trans)

	account, err := queries.GetAccountByIDLocked(ctx, accountID)
	if err != nil {
		return nil, err
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

func (alc *AccountLoaderSaver) Load(ctx context.Context, accountID uuid.UUID) (eventsourcing.Aggregate, error) {
	queries := alc.queries

	if trans, ok := eventstore.GetTx(ctx); ok {
		queries = alc.queries.WithTx(trans)
	}

	account, err := queries.GetAccountByID(ctx, accountID)
	if err != nil {
		return nil, err
	}

	return fromSQLCAccount(account), nil
}

func (alc *AccountLoaderSaver) Save(ctx context.Context, aggregate eventsourcing.Aggregate) error {
	queries := alc.queries

	if trans, ok := eventstore.GetTx(ctx); ok {
		queries = alc.queries.WithTx(trans)
	}

	account, ok := aggregate.(*Account)
	if !ok {
		return &TypeMismatchError{
			expect: &Account{},
			got:    aggregate,
		}
	}

	if err := queries.UpsertAccount(ctx, &sqlcquery.UpsertAccountParams{
		ID:        account.ID,
		Balance:   account.Balance,
		Available: account.Available,
		Pending:   account.Pending,
		Version:   int32(account.Version),
		UpdatedAt: account.UpdatedAt,
		CreatedAt: account.CreatedAt,
	}); err != nil {
		return fmt.Errorf("failed to upsert account: %w", err)
	}

	return nil
}
