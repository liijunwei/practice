package common

import (
	"context"
	"encoding/json"
	"fmt"
	"greenlight/internal/approot"
	"greenlight/internal/assert"
	"greenlight/internal/data"
	"net/http"
	"runtime/debug"
	"strings"
)

type Envelope map[string]any

func WriteResponseJSON(w http.ResponseWriter, status int, data Envelope, headers http.Header) error {
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}
	for key, value := range headers {
		w.Header()[key] = value
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
	w.Write([]byte("\n"))

	return nil
}

func RenderNotFound(w http.ResponseWriter, r *http.Request) {
	message := "resource not found"
	RenderError(w, r, http.StatusNotFound, message)
}

func RenderError(w http.ResponseWriter, r *http.Request, status int, message any) {
	env := Envelope{"error": message}
	if err := WriteResponseJSON(w, status, env, nil); err != nil {
		// app.logError(r, err)
		w.WriteHeader(500)
	}
}

func RenderInternalServerError(w http.ResponseWriter, r *http.Request, err error, debug bool) {
	// app.logError(r, err)
	message := "server failed unexpectedly"

	if debug {
		details := map[string]any{
			"message": message,
			"error":   err.Error(),
			"traces":  sanitizedDebugTraces(),
		}

		RenderError(w, r, http.StatusServiceUnavailable, details)
		return
	}

	RenderError(w, r, http.StatusServiceUnavailable, message)
}

func sanitizedDebugTraces() []string {
	rawTraces := strings.Split(string(debug.Stack()), "\n")
	result := make([]string, 0, len(rawTraces))
	prefixToTrim := fmt.Sprintf("%s/", approot.Root)

	for _, trace := range rawTraces {
		if strings.Contains(trace, approot.Root) {
			fields := strings.Fields(trace) // ["/approot/foo.go:10", "+0x2d4"]
			assert.Assert(len(fields) == 2)
			result = append(result, strings.TrimPrefix(fields[0], prefixToTrim))
		}
	}

	return result
}

type contextKey string

const userContextKey = contextKey("user")

func SetUserToContext(r *http.Request, user *data.User) *http.Request {
	ctx := context.WithValue(r.Context(), userContextKey, user)
	return r.WithContext(ctx)
}

func GetUserFromContext(r *http.Request) *data.User {
	user, ok := r.Context().Value(userContextKey).(*data.User)
	assert.Assert(ok, "user must present in request context")

	return user
}

func RenderInvalidAuthenticationToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("WWW-Authenticate", "Bearer")
	message := "invalid or missing authentication token"
	RenderError(w, r, http.StatusUnauthorized, message)
}

func RenderRateLimitExceeded(w http.ResponseWriter, r *http.Request) {
	message := "rate limit exceeded"
	RenderError(w, r, http.StatusTooManyRequests, message)
}
