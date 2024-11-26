package main

import (
	"encoding/json"
	"net/http"

	"github.com/ericlagergren/decimal"
	"github.com/gofrs/uuid"
	"github.com/rs/zerolog/log"
)

func createAccountHandler(accountRepo *AccountRepository) http.HandlerFunc {
	return func(respWriter http.ResponseWriter, req *http.Request) {
		ctx := req.Context()

		// parse query
		initialBalance := decimal.New(10000, 0) //nolint: mnd // example

		// create account
		account, err := NewAccount(initialBalance)
		if err != nil {
			respWriter.WriteHeader(http.StatusInternalServerError)

			return
		}

		// save it to database
		if err := accountRepo.Save(ctx, account); err != nil {
			respWriter.WriteHeader(http.StatusInternalServerError)

			return
		}

		// write response
		respWriter.WriteHeader(http.StatusCreated)

		if err := json.NewEncoder(respWriter).Encode(account); err != nil {
			log.Warn().Err(err).
				Str("account_id", account.ID.String()).
				Msg("failed to write response")
		}
	}
}

func getAccountHandler(accountRepo *AccountRepository) http.HandlerFunc {
	return func(respWriter http.ResponseWriter, req *http.Request) {
		ctx := req.Context()

		// parse query
		idString := req.URL.Query().Get("id")
		if idString == "" {
			respWriter.WriteHeader(http.StatusBadRequest)

			return
		}

		accountID, err := uuid.FromString(idString)
		if err != nil {
			respWriter.WriteHeader(http.StatusBadRequest)

			return
		}

		// load it from database
		account, err := accountRepo.Load(ctx, accountID)
		if err != nil {
			respWriter.WriteHeader(http.StatusNotFound)

			return
		}

		// write response
		respWriter.WriteHeader(http.StatusCreated)

		if err := json.NewEncoder(respWriter).Encode(account); err != nil {
			log.Warn().Err(err).
				Str("account_id", account.ID.String()).
				Msg("failed to write response")
		}
	}
}
