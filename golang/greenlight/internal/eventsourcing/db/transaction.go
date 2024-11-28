package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type contextTxKey struct{}

func Transaction(ctx context.Context, dbPool *pgxpool.Pool, txFunc func(ctx context.Context, tx pgx.Tx) error) error {
	var err error

	pgtx, ok := GetTx(ctx)

	if ok { // already inside a transaction, start nested transaction
		tx, err := pgtx.Begin(ctx)
		if err != nil {
			return &TransactionError{err: err}
		}

		defer tx.Rollback(ctx)

		if err := txFunc(WithTx(ctx, tx), tx); err != nil {
			return fmt.Errorf("tx func error: %w", err)
		}

		return tx.Commit(ctx)
	}

	// start a new transaction
	err = pgx.BeginTxFunc(ctx, dbPool, pgx.TxOptions{}, func(tx pgx.Tx) error {
		ctx = WithTx(ctx, tx)

		return txFunc(ctx, tx)
	})
	if err != nil {
		return &TransactionError{err: err}
	}

	return nil
}

// GetTx retrevies a pgx.tx from context
func GetTx(ctx context.Context) (pgx.Tx, bool) { //nolint: ireturn // pgx.tx is a interface
	tx, ok := ctx.Value(contextTxKey{}).(pgx.Tx)

	return tx, ok
}

// WithTx adds a pgx.tx from context
func WithTx(ctx context.Context, tx pgx.Tx) context.Context {
	return context.WithValue(ctx, contextTxKey{}, tx)
}
