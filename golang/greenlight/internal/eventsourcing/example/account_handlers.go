package main

import (
	"encoding/json"
	"net/http"

	"github.com/ericlagergren/decimal"
	"github.com/gofrs/uuid"
	"github.com/rs/zerolog/log"
)

func createAccountHandler(accountRepo *AccountRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		ctx := r.Context()

		// parse query
		initialBalance := decimal.New(10000, 0) //nolint: mnd // example

		// create account
		account, err := NewAccount(initialBalance)
		if err != nil {
			log.Error().Err(err).Msg("NewAccount error")

			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			w.Write([]byte("\n"))

			return
		}

		// save it to database
		if err := accountRepo.Save(ctx, account); err != nil {
			log.Error().Err(err).Msg("save account error")

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

func getAccountHandler(accountRepo *AccountRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		ctx := r.Context()

		idString := r.URL.Query().Get("id")
		if idString == "" {
			log.Error().Msg("no id provided")
			w.WriteHeader(http.StatusBadRequest)

			return
		}

		accountID, err := uuid.FromString(idString)
		if err != nil {
			log.Error().Err(err).Msg("not valid id")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			w.Write([]byte("\n"))

			return
		}

		account, err := accountRepo.Load(ctx, accountID)
		log.Error().Err(err).Msg("account not found")

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
