package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	if err := run(); err != nil {
		log.Error().Err(err).Msg("api failed with an error")

		os.Exit(1)
	}
}

func run() error {
	ctx := context.Background()
	logger := zerolog.New(os.Stdout)
	ctx = logger.With().Str("project", "event-sourcing-example").Logger().WithContext(ctx)

	connPool, err := pgxpool.New(ctx, os.Getenv("GREENLIGHT_DB_DSN"))
	if err != nil {
		return err
	}

	accountRepo := NewAccountRepository(connPool)
	debitHoldRepo := NewDebitHoldRepository(connPool)

	router := chi.NewRouter()
	router.Post("/api/account", createAccountHandler(accountRepo))
	router.Get("/api/account", getAccountHandler(accountRepo))
	router.Post("/api/debit-hold", createDebitHoldHandler(connPool, accountRepo, debitHoldRepo))

	if err := http.ListenAndServe("127.0.0.1:3000", router); err != nil {
		return fmt.Errorf("http server error: %w", err)
	}

	return nil
}
