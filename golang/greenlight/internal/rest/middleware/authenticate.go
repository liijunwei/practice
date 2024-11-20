package middleware

import (
	"errors"
	"greenlight/internal/common"
	"greenlight/internal/data"
	"greenlight/internal/validator"
	"net/http"
	"strings"
)

func Authenticate(next http.Handler, models data.Models, debug bool) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO this is new to me, learn more
		// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Vary
		w.Header().Add("Vary", "Authorization")

		authorizationHeader := r.Header.Get("Authorization")
		if authorizationHeader == "" {
			r = common.ContextSetUser(r, data.AnonymousUser)
			next.ServeHTTP(w, r)
			return
		}

		headerParts := strings.Split(authorizationHeader, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			common.InvalidAuthenticationTokenResponse(w, r)
			return
		}

		token := headerParts[1]
		v := validator.New()

		if data.ValidateTokenPlaintext(v, token); !v.Valid() {
			common.InvalidAuthenticationTokenResponse(w, r)
			return
		}

		user, err := models.Users.GetByToken(r.Context(), data.ScopeAuthentication, token)
		if err != nil {
			switch {
			case errors.Is(err, data.ErrRecordNotFound):
				common.NotFoundResponse(w, r)
			default:
				common.ServerErrorResponse(w, r, err, debug)
			}

			// app.logger.Warn().Err(err).Msg("failed to find user by auth token")

			return
		}

		r = common.ContextSetUser(r, user)

		next.ServeHTTP(w, r)
	})
}
