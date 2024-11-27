package main

import (
	"context"
	"encoding/json"
	"net/http"

	"greenlight/internal/eventsourcing/db"

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
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		ctx := r.Context()

		// parse json payload
		var data createDebitHoldRequest
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			w.Write([]byte(err.Error()))
			w.Write([]byte("\n"))

			return
		}

		var debitHold *DebitHold

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
			debitHold, err = NewDebitHold(account.ID, data.Amount)
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

			// err = accountRepo.Save(ctx, account)

			return nil
		})

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			w.Write([]byte("\n"))
			return
		}

		resp, err := json.Marshal(debitHold)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			w.Write([]byte("\n"))
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write(resp)
		w.Write([]byte("\n"))
	}
}
