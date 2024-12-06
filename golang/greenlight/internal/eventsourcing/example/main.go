package main

import (
	"context"
	"encoding/json"
	"greenlight/internal/ericlagergren"
	"greenlight/internal/eventsourcing/eventstore"
	"greenlight/internal/eventsourcing/example/domain"
	"net/http"
	"os"

	"github.com/ericlagergren/decimal"
	"github.com/gofrs/uuid"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	ctx := context.Background()
	logger := zerolog.New(os.Stdout)
	ctx = logger.With().Str("project", "event-sourcing-example").Logger().WithContext(ctx)

	conf, err := pgxpool.ParseConfig(os.Getenv("GREENLIGHT_DB_DSN"))
	if err != nil {
		return err
	}

	conf.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		ericlagergren.Register(conn.TypeMap())

		return nil
	}

	connPool, err := pgxpool.NewWithConfig(ctx, conf)
	if err := connPool.Ping(ctx); err != nil {
		return err
	}

	accountRepo := domain.NewAccountRepository(connPool)
	debitHoldRepo := domain.NewDebitHoldRepository(connPool)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/account", createAccountHandler(accountRepo))
	mux.HandleFunc("GET /api/account", getAccountHandler(accountRepo))
	mux.HandleFunc("POST /api/debit-hold", createDebitHoldHandler(connPool, accountRepo, debitHoldRepo))

	if err := http.ListenAndServe("127.0.0.1:3000", mux); err != nil {
		return err
	}

	return nil
}

func createAccountHandler(accountRepo *domain.AccountRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		ctx := r.Context()

		initialBalance := decimal.New(10000, 0) //nolint: mnd // example
		account, err := domain.NewAccount(initialBalance)
		if err != nil {
			zerolog.Ctx(ctx).Error().Err(err).Msg("NewAccount error")

			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			w.Write([]byte("\n"))

			return
		}

		// save it to database
		if err := accountRepo.Save(ctx, account); err != nil {
			zerolog.Ctx(ctx).Error().Err(err).Msg("save account error")

			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			w.Write([]byte("\n"))

			return
		}

		// write response
		w.WriteHeader(http.StatusCreated)

		if err := json.NewEncoder(w).Encode(account); err != nil {
			log.Warn().Err(err).
				Str("account_id", account.ID.String()).
				Msg("failed to write response")
		}
	}
}

func getAccountHandler(accountRepo *domain.AccountRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		ctx := r.Context()

		idString := r.URL.Query().Get("id")
		if idString == "" {
			zerolog.Ctx(ctx).Error().Msg("no id provided")
			w.WriteHeader(http.StatusBadRequest)

			return
		}

		accountID, err := uuid.FromString(idString)
		if err != nil {
			zerolog.Ctx(ctx).Error().Err(err).Msg("not valid id")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			w.Write([]byte("\n"))

			return
		}

		account, err := accountRepo.Load(ctx, accountID)
		zerolog.Ctx(ctx).Error().Err(err).Msg("account not found")

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
			w.Write([]byte("\n"))

			return
		}

		w.WriteHeader(http.StatusCreated)

		data, err := json.Marshal(account)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			w.Write([]byte("\n"))

			return
		}

		w.Write(data)
		w.Write([]byte("\n"))
	}
}

type createDebitHoldRequest struct {
	AccountID uuid.UUID
	Amount    *decimal.Big
}

func createDebitHoldHandler(
	dbPool *pgxpool.Pool, accountRepo *domain.AccountRepository, debitHoldRepo *domain.DebitHoldRepository,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		ctx := r.Context()

		var data createDebitHoldRequest
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			w.Write([]byte(err.Error()))
			w.Write([]byte("\n"))

			return
		}

		var debitHold *domain.DebitHold

		err := eventstore.Transaction(ctx, dbPool, func(ctx context.Context, _ pgx.Tx) error {
			account, err := accountRepo.LoadLocked(ctx, data.AccountID)
			if err != nil {
				return err
			}

			if account.Available.Cmp(data.Amount) <= 0 {
				err := &domain.NotEnoughBalanceError{}

				return err
			}

			debitHold, err = domain.NewDebitHold(account.ID, data.Amount)
			if err != nil {
				return err
			}

			err = debitHoldRepo.Save(ctx, debitHold)
			if err != nil {
				return err
			}

			err = account.MoveAvailableToPending(data.Amount)
			if err != nil {
				return err
			}

			err = accountRepo.Save(ctx, account)
			if err != nil {
				return err
			}

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
