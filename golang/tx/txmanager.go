package tx

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Manager is an interface that defines a method for executing a function within a db transaction context.
type Manager interface {
	WithinTransaction(ctx context.Context, fn func(tx *sql.Tx) error) error
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

// https://go.dev/doc/database/execute-transactions
func (m *PGManager) WithinTransaction(ctx context.Context, txFunc func(tx *sql.Tx) error) error {
	tx, err := m.conn.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	defer tx.Rollback()

	if err := txFunc(tx); err != nil {
		return fmt.Errorf("failed to execute function within transaction: %w", err)
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}
