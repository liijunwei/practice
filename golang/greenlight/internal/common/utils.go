package common

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"greenlight/internal/approot"
	"greenlight/internal/assert"
	"greenlight/internal/data"
	"greenlight/internal/mailer"
	"greenlight/internal/validator"
	"io"
	"net/http"
	"net/url"
	"runtime/debug"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/rs/zerolog"
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

func ReadJSON(_ http.ResponseWriter, r *http.Request, dst any) error {
	if err := json.NewDecoder(r.Body).Decode(dst); err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var invalidUnmarshalError *json.InvalidUnmarshalError

		switch {
		case errors.As(err, &syntaxError):
			return fmt.Errorf("json SyntaxError at char %d", syntaxError.Offset)
		case errors.Is(err, io.ErrUnexpectedEOF):
			return fmt.Errorf("UnexpectedEOF")
		case errors.As(err, &unmarshalTypeError):
			return fmt.Errorf("body contains incorrect json type")
		case errors.Is(err, io.EOF):
			return fmt.Errorf("body must not be empty")
		case errors.As(err, &invalidUnmarshalError):
			panic(err)
		default:
			return fmt.Errorf("no error handler %w", err)
		}
	}

	return nil
}

func ReadString(qs url.Values, key string, defaultVal string) string {
	if s := qs.Get(key); s != "" {
		return s
	}

	return defaultVal
}

func ReadCSV(qs url.Values, key string, defaultVal []string) []string {
	if s := qs.Get(key); s != "" {
		return strings.Split(s, ",")
	}

	return defaultVal
}

func ReadInt(qs url.Values, key string, defaultVal int, v *validator.Validator) int {
	if s := qs.Get(key); s != "" {
		i, err := strconv.Atoi(s)
		if err != nil {
			v.AddError(key, "must be an integer value")
			return defaultVal
		}

		return i
	}

	return defaultVal
}

func ReadIDParam(r *http.Request) (int64, error) {
	str := r.PathValue("id")
	id, err := strconv.ParseInt(str, 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id param")
	}

	return id, nil
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

func RenderInternalServerError(w http.ResponseWriter, r *http.Request, err error) {
	// app.logError(r, err)
	message := "server failed unexpectedly"

	details := map[string]any{
		"message": message,
		"error":   err.Error(),
		"traces":  sanitizedDebugTraces(),
	}

	RenderError(w, r, http.StatusServiceUnavailable, details)
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

func RenderInvlidCredentials(w http.ResponseWriter, r *http.Request) {
	message := "invalid authentication credentials"
	RenderError(w, r, http.StatusUnauthorized, message)
}

func RenderAuthenticationRequired(w http.ResponseWriter, r *http.Request) {
	message := "you must be authenticated to access this resource"
	RenderError(w, r, http.StatusUnauthorized, message)
}

func RenderInactiveAccount(w http.ResponseWriter, r *http.Request) {
	message := "you user account must be activated to access this resource"
	RenderError(w, r, http.StatusForbidden, message)
}

func RenderNotPermitted(w http.ResponseWriter, r *http.Request) {
	message := "you user account doesn't have necessary permissions to access this resource"
	RenderError(w, r, http.StatusForbidden, message)
}

func RenderBadRequest(w http.ResponseWriter, r *http.Request, err error) {
	RenderError(w, r, http.StatusBadRequest, err.Error())
}

func RenderFailedValidation(w http.ResponseWriter, r *http.Request, errors map[string]string) {
	RenderError(w, r, http.StatusUnprocessableEntity, errors)
}

func RenderEditStaleRecord(w http.ResponseWriter, r *http.Request) {
	message := "unable to update a stale object, please try again"
	RenderError(w, r, http.StatusConflict, message)
}

func RenderNotFoundOrUnknownError(err error, w http.ResponseWriter, r *http.Request) {
	switch {
	case errors.Is(err, data.ErrRecordNotFound):
		RenderNotFound(w, r)
	default:
		RenderInternalServerError(w, r, err)
	}
}

func RunInBackground(ctx context.Context, wg *sync.WaitGroup, fn func()) {
	wg.Add(1)

	go func() {
		defer wg.Done()

		defer func() {
			if err := recover(); err != nil {
				zerolog.Ctx(ctx).Error().Err(fmt.Errorf("%s", err)).Msg("recover panic failed")
			}
		}()

		fn()
	}()
}

func SendEmail(ctx context.Context, user *data.User, token *data.Token, mailer mailer.Mailer, sender string) {
	logger := zerolog.Ctx(ctx).With().Int64("user_id", user.ID).Logger()

	data := map[string]any{
		"userID":          user.ID,
		"activationToken": token.Plaintext,
	}

	logger.Info().Msg("send email start")
	startTime := time.Now()

	if err := mailer.Send(sender, user.Email, "user_welcome.tmpl", data); err != nil {
		logger.Error().Err(err).Msg("send email failed")
		return
	}

	logger.Info().Str("duration", time.Since(startTime).String()).Msg("send email done")
}
