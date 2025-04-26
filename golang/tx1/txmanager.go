package tx

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// example ref: https://dev.to/techschoolguru/a-clean-way-to-implement-database-transaction-in-golang-2ba

// Manager is an interface that defines a method for executing a function within a db transaction context.
type Manager interface {
	ExecTx(ctx context.Context, fn func(ctx context.Context) error) error
}

var _ Manager = &PGManager{}

type PGManager struct {
	conn *sql.DB
}

func New(db *sql.DB) Manager {
	return &PGManager{
		conn: db,
	}
}

// Transaction key to store and retrieve the transaction from context
type txKey struct{}

// MustGetTx retrieves the transaction from the context
func MustGetTx(ctx context.Context) *sql.Tx {
	tx, ok := ctx.Value(txKey{}).(*sql.Tx)
	if !ok {
		panic("assert transaction in context")
	}
	return tx
}

// https://go.dev/doc/database/execute-transactions
func (m *PGManager) ExecTx(ctx context.Context, txFunc func(ctx context.Context) error) error {
	tx, err := m.conn.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	defer tx.Rollback()

	// Create a new context with the transaction
	txCtx := context.WithValue(ctx, txKey{}, tx)

	if err := txFunc(txCtx); err != nil {
		return fmt.Errorf("failed to execute function within transaction: %w", err)
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}
