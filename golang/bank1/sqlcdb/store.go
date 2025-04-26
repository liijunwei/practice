package sqlcdb

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type Store struct {
	queries *Queries // TODO unclear yet
	db      *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		queries: New(db),
		db:      db,
	}
}

const maxTxTimeout = 5 * time.Second

func (s *Store) WithTx(ctx context.Context, txFunc func(*Queries) error) error {
	ctx, cancel := context.WithTimeout(ctx, maxTxTimeout)
	defer cancel()

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)

	if err := txFunc(q); err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("txFunc failed: %v, unable to rollback: %v", err, rbErr)
		}
		return fmt.Errorf("txFunc failed: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}
