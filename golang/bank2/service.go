package main

import (
	"context"
	"database/sql"
	"fmt"
	"golang-practices/bank1/sqlcdb"

	"strings"

	"github.com/google/uuid"
)

const TransactionKindTransfer = "internal_transfer"

type AccountService struct {
	queries *sqlcdb.Queries
	store   *sqlcdb.Store
}

func NewAccountService(db *sql.DB) *AccountService {
	return &AccountService{
		queries: sqlcdb.New(db),
		store:   sqlcdb.NewStore(db),
	}
}

func (s *AccountService) Register(
	ctx context.Context,
	username, password, email string,
	currencies []string,
	initialBalanceBalance float64, // TODO add deposit feature
) (*sqlcdb.User, error) {
	var user sqlcdb.User

	var err error

	errT := s.store.WithTx(ctx, func(q *sqlcdb.Queries) error {
		user, err = q.CreateUser(ctx, sqlcdb.CreateUserParams{
			ID:       uuid.New().String(),
			Username: username,
			Password: password,
			Email:    email,
		})
		if err != nil {
			return fmt.Errorf("failed to create user: %w", err)
		}

		for _, currency := range currencies {
			_, err := q.CreateAccount(ctx, sqlcdb.CreateAccountParams{
				ID:        uuid.New().String(),
				UserID:    user.ID,
				Currency:  currency,
				Available: initialBalanceBalance,
			})

			if err != nil {
				return fmt.Errorf("failed to create account: %w", err)
			}
		}

		return nil
	})
	if errT != nil {
		return nil, fmt.Errorf("failed to register within transaction: %w", errT)
	}

	return &user, nil
}

// TODO use decimal
func (s *AccountService) Transfer(
	ctx context.Context,
	fromAccountID, toAccountID string,
	amount float64,
) error {
	if amount <= 0 {
		return fmt.Errorf("amount should be positive: %f", amount)
	}

	assert(amount > 0)

	fromAccount, err := s.queries.GetAccount(ctx, fromAccountID)
	if err != nil {
		return fmt.Errorf("failed to get from_account: %w", err)
	}

	toAccount, err := s.queries.GetAccount(ctx, toAccountID)
	if err != nil {
		return fmt.Errorf("failed to get to_account: %w", err)
	}

	if fromAccount.Currency != toAccount.Currency {
		return fmt.Errorf("currency mismatch: %s vs %s", fromAccount.Currency, toAccount.Currency)
	}

	if fromAccount.Available <= 0 {
		return fmt.Errorf("balance in from_account <= 0 (account_id: %s)", fromAccount.ID)
	}

	if toAccount.Available <= 0 {
		return fmt.Errorf("balance in to_account <= 0 (account_id: %s)", toAccount.ID)
	}

	assert(fromAccount.Available >= 0)
	assert(toAccount.Available >= 0)

	if fromAccount.Available < amount {
		return fmt.Errorf("insufficient available balance in from_account(%s): %f < %f", fromAccount.ID, fromAccount.Available, amount)
	}

	errT := s.store.WithTx(ctx, func(q *sqlcdb.Queries) error {
		_, err = q.DebitAccount(ctx, sqlcdb.DebitAccountParams{
			ID:          fromAccount.ID,
			Amount:      amount,
			LockVersion: fromAccount.LockVersion,
		})
		if err != nil {
			if isStaleObjectError(err) {
				return ErrStaleObject
			}

			return fmt.Errorf("failed to debit from_account: %w", err)
		}

		err = q.CreateAccountEvent(ctx, sqlcdb.CreateAccountEventParams{
			ID:        uuid.New().String(),
			AccountID: fromAccount.ID,
			Amount:    -amount,
		})
		if err != nil {
			return fmt.Errorf("failed to append account event: %w", err)
		}

		_, err = q.CreditAccount(ctx, sqlcdb.CreditAccountParams{
			ID:          toAccount.ID,
			Amount:      amount,
			LockVersion: toAccount.LockVersion,
		})
		if err != nil {
			if isStaleObjectError(err) {
				return ErrStaleObject
			}

			return fmt.Errorf("failed to credit to_account: %w", err)
		}

		err = q.CreateAccountEvent(ctx, sqlcdb.CreateAccountEventParams{
			ID:        uuid.New().String(),
			AccountID: toAccount.ID,
			Amount:    amount,
		})
		if err != nil {
			return fmt.Errorf("failed to append account event: %w", err)
		}

		_, err = q.CreateTransaction(ctx, sqlcdb.CreateTransactionParams{
			ID:            uuid.New().String(),
			FromAccountID: fromAccount.ID,
			ToAccountID:   toAccount.ID,
			Amount:        amount,
			Description:   "transfer",
			Kind:          TransactionKindTransfer,
		})
		if err != nil {
			return fmt.Errorf("failed to create transaction: %w", err)
		}

		return nil
	})

	if errT != nil {
		return fmt.Errorf("failed to transfer within transaction: %w", errT)
	}

	return nil
}

func isStaleObjectError(err error) bool {
	if err == nil {
		return false
	}

	if strings.Contains(err.Error(), "no rows in result set") {
		return true
	}

	return false
}

var ErrStaleObject = fmt.Errorf("stale object error")
