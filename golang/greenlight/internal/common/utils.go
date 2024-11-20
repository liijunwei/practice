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

func WriteJSON(w http.ResponseWriter, status int, data Envelope, headers http.Header) error {
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

func NotFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "resource not found"
	ErrorResponse(w, r, http.StatusNotFound, message)
}

func ErrorResponse(w http.ResponseWriter, r *http.Request, status int, message any) {
	env := Envelope{"error": message}
	if err := WriteJSON(w, status, env, nil); err != nil {
		// app.logError(r, err)
		w.WriteHeader(500)
	}
}

func ServerErrorResponse(w http.ResponseWriter, r *http.Request, err error, debug bool) {
	// app.logError(r, err)
	message := "server failed unexpectedly"

	if debug {
		details := map[string]any{
			"message": message,
			"error":   err.Error(),
			"traces":  sanitizedDebugTraces(),
		}

		ErrorResponse(w, r, http.StatusServiceUnavailable, details)
		return
	}

	ErrorResponse(w, r, http.StatusServiceUnavailable, message)
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

func ContextSetUser(r *http.Request, user *data.User) *http.Request {
	ctx := context.WithValue(r.Context(), userContextKey, user)
	return r.WithContext(ctx)
}

func ContextGetUser(r *http.Request) *data.User {
	user, ok := r.Context().Value(userContextKey).(*data.User)
	assert.Assert(ok, "user must present in request context")

	return user
}
