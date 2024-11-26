package main

import (
	"context"
	"encoding/json"
	"net/http"

	"knockout/internal/eventsourcing/db"

	"github.com/ericlagergren/decimal"
	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type createDebitHoldRequest struct {
	AccountID uuid.UUID
	Amount    *decimal.Big
}

func createDebitHoldHandler(
	dbPool *pgxpool.Pool, accountRepo *AccountRepository, debitHoldRepo *DebitHoldRepository,
) http.HandlerFunc {
	return func(respWriter http.ResponseWriter, req *http.Request) {
		ctx := req.Context()

		// parse json payload
		var data createDebitHoldRequest
		if err := json.NewDecoder(req.Body).Decode(&data); err != nil {
			respWriter.WriteHeader(http.StatusUnprocessableEntity)

			return
		}

		err := db.Transaction(ctx, dbPool, func(ctx context.Context, _ pgx.Tx) error {
			// load account, locked for update
			account, err := accountRepo.LoadLocked(ctx, data.AccountID)
			if err != nil {
				return err
			}

			// check balance
			if account.Available.Cmp(data.Amount) <= 0 {
				err := &NotEnoughBalanceError{}

				return err
			}

			// create a debit hold
			debitHold, err := NewDebitHold(account.ID, data.Amount)
			if err != nil {
				return err
			}

			err = debitHoldRepo.Save(ctx, debitHold)
			if err != nil {
				return err
			}

			// change account pending amount
			err = account.MoveAvailableToPending(data.Amount)
			if err != nil {
				return err
			}

			err = accountRepo.Save(ctx, account)

			return err
		})

		if err != nil {
			respWriter.WriteHeader(http.StatusBadRequest)
		} else {
			respWriter.WriteHeader(http.StatusCreated)
		}
	}
}
