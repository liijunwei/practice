package main

import (
	"context"
	"fmt"
	"greenlight/internal/ericlagergren"
	"net/http"
	"os"

	"github.com/rs/zerolog"

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

	accountRepo := NewAccountRepository(connPool)
	debitHoldRepo := NewDebitHoldRepository(connPool)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/account", createAccountHandler(accountRepo))
	mux.HandleFunc("GET /api/account", getAccountHandler(accountRepo))
	mux.HandleFunc("POST /api/debit-hold", createDebitHoldHandler(connPool, accountRepo, debitHoldRepo))

	fmt.Println("server started")

	if err := http.ListenAndServe("127.0.0.1:3000", mux); err != nil {
		return err
	}

	return nil
}
