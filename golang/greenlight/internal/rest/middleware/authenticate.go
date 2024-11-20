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
			r = common.SetUserToContext(r, data.AnonymousUser)
			next.ServeHTTP(w, r)
			return
		}

		headerParts := strings.Split(authorizationHeader, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			common.RenderInvalidAuthenticationToken(w, r)
			return
		}

		token := headerParts[1]
		v := validator.New()

		if data.ValidateTokenPlaintext(v, token); !v.Valid() {
			common.RenderInvalidAuthenticationToken(w, r)
			return
		}

		user, err := models.Users.GetByToken(r.Context(), data.ScopeAuthentication, token)
		if err != nil {
			switch {
			case errors.Is(err, data.ErrRecordNotFound):
				common.RenderNotFound(w, r)
			default:
				common.RenderInternalServerError(w, r, err)
			}

			// app.logger.Warn().Err(err).Msg("failed to find user by auth token")

			return
		}

		r = common.SetUserToContext(r, user)

		next.ServeHTTP(w, r)
	})
}
