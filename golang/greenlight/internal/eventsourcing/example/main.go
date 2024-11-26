package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"greenlight/internal/config"
	"greenlight/internal/postgres"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
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

	cfg, err := config.LoadWithEnv(ctx, configPath)
	if err != nil {
		log.Fatal().Err(err).Msg("load general config failed")
	}

	pgxPool, err := postgres.NewConnPool(ctx, &cfg.DB)
	if err != nil {
		return fmt.Errorf("connect to postgres failed: %w", err)
	}

	// Create repositories
	accountRepo := NewAccountRepository(pgxPool)
	debitHoldRepo := NewDebitHoldRepository(pgxPool)

	// router
	router := chi.NewRouter()

	// create account
	router.Post("/api/account", createAccountHandler(accountRepo))
	// get account
	router.Get("/api/account", getAccountHandler(accountRepo))

	// create debit hold
	router.Post("/api/debit-hold", createDebitHoldHandler(pgxPool, accountRepo, debitHoldRepo))

	if err := http.ListenAndServe("127.0.0.1:3000", router); err != nil { //nolint:gosec // this is just an example
		return fmt.Errorf("http server error: %w", err)
	}

	return nil
}
