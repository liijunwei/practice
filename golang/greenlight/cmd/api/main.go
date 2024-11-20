package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"expvar"
	"flag"
	"fmt"
	"greenlight/internal/approot"
	"greenlight/internal/assert"
	"greenlight/internal/common"
	"greenlight/internal/config"
	"greenlight/internal/data"
	"greenlight/internal/mailer"
	"greenlight/internal/rest"
	"greenlight/internal/rest/middleware"
	"greenlight/internal/sqlcdb"
	"greenlight/internal/validator"
	"io"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"runtime"
	"runtime/debug"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"

	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const version = "1.0.0"

func main() {
	cfg := appconfig()
	run(cfg)
}

func appconfig() config.Config {
	var cfg config.Config
	flag.StringVar(&cfg.Env, "env", "development", "environment(development|staging|production)")
	flag.StringVar(&cfg.IP, "ip", "localhost", "the server ip address")
	flag.IntVar(&cfg.Port, "port", 4000, "api server port")
	flag.StringVar(&cfg.DB.DSN, "db-dsn", os.Getenv("GREENLIGHT_DB_DSN"), "postgres Data Source Name (DSN)") // from sensitive-exports.sh
	flag.IntVar(&cfg.DB.MaxOpenConns, "db-max-open-conns", 25, "postgres max open connections")
	flag.IntVar(&cfg.DB.MaxIdleConns, "db-max-idle-conns", 25, "postgres max idle connections")
	flag.StringVar(&cfg.DB.MaxIdleTime, "db-max-idle-time", "15m", "postgres max connection idle time")
	flag.Float64Var(&cfg.Limiter.RPS, "limiter-rps", 2, "rate limiter maximum requests per second")
	flag.IntVar(&cfg.Limiter.Burst, "limiter-burst", 4, "rate limiter maximum burst")
	flag.BoolVar(&cfg.Limiter.Enabled, "limiter-enabled", true, "enable rate limiter")
	flag.BoolVar(&cfg.Debug, "debug-enabled", false, "verbose api response")
	flag.StringVar(&cfg.SMTP.Host, "smtp-host", "sandbox.smtp.mailtrap.io", "SMTP host")
	flag.IntVar(&cfg.SMTP.Port, "smtp-port", 2525, "SMTP port")
	flag.StringVar(&cfg.SMTP.Username, "smtp-username", os.Getenv("SMTP_USERNAME"), "SMTP username")
	flag.StringVar(&cfg.SMTP.Password, "smtp-password", os.Getenv("SMTP_PASSWORD"), "SMTP password")
	flag.StringVar(&cfg.SMTP.Sender, "smtp-sender", "greenlight-admin@example.com", "SMTP sender")
	flag.Func("cors-trusted-origins", "Trusted CORS origins(space separated)", func(val string) error {
		originsSet := make(map[string]struct{})
		origins := strings.Fields(val)
		for _, o := range origins {
			assert.Assert(o != "", "should not trusted empty origin")
			assert.Assert(o != "null", "should not trusted null origin")
			originsSet[o] = struct{}{}
		}

		cfg.CORS.TrustedOrigins = originsSet

		return nil
	})
	flag.StringVar(&cfg.LogConfig.Level, "log-level", "debug", "log level")

	flag.Parse()

	return cfg
}

func run(cfg config.Config) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	logger := log.Logger.With().Str("service", "greenlight-api").Logger()

	level, err := zerolog.ParseLevel(cfg.LogConfig.Level)
	assert.Assert(err == nil)
	assert.Assert(level != zerolog.NoLevel)

	zerolog.SetGlobalLevel(level)
	logger.Level(level)

	ctx := context.Background()
	ctx = logger.WithContext(ctx)

	db, err := openDB(cfg)
	assert.Assert(err == nil)

	zerolog.Ctx(ctx).Info().Msg("database connection pool established")

	defer db.Close()

	queries := sqlcdb.New(db)

	app := &application{
		config:  cfg,
		logger:  logger,
		models:  data.NewModels(queries),
		queries: queries,
		mailer:  mailer.New(cfg.SMTP.Host, cfg.SMTP.Port, cfg.SMTP.Username, cfg.SMTP.Password),
	}

	expvar.NewString("version").Set(version)
	expvar.Publish("goroutines", expvar.Func(func() any { return runtime.NumGoroutine() }))
	expvar.Publish("database", expvar.Func(func() any { return db.Stats() }))
	expvar.Publish("timestamp", expvar.Func(func() any { return time.Now().Unix() }))

	if err := app.serve(); err != nil {
		logger.Fatal().Err(err).Msg("unexpected server error")
	}
}

type application struct {
	config  config.Config
	logger  zerolog.Logger
	models  data.Models
	queries *sqlcdb.Queries
	mailer  mailer.Mailer
	wg      sync.WaitGroup
}

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /v1/healthcheck", rest.HealthcheckHandler(app.config.Env, version))
	mux.HandleFunc("POST /v1/movies", app.requirePermission("movies:write", app.createMovieHandler))
	mux.HandleFunc("GET /v1/movies/{id}", app.requirePermission("movies:read", app.showMovieHandler))
	mux.HandleFunc("PUT /v1/movies/{id}", app.requirePermission("movies:write", app.updateMovieHandler))
	mux.HandleFunc("DELETE /v1/movies/{id}", app.requirePermission("movies:write", app.DeleteMovieHandler))
	mux.HandleFunc("GET /v1/movies", app.requirePermission("movies:read", app.listMovieHandler))
	mux.HandleFunc("POST /v1/users", app.registerUserHandler)
	mux.HandleFunc("PUT /v1/users/activated", app.activateUserHandler)
	mux.HandleFunc("POST /v1/tokens/authentication", app.createAuthenticationTokenHandler)
	mux.Handle("GET /debug/vars", expvar.Handler())

	return mux
}

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title   string   `json:"title"`
		Year    int32    `json:"year"`
		Runtime int32    `json:"runtime"`
		Genres  []string `json:"genres"`
	}

	if err := app.readJSON(w, r, &input); err != nil {
		common.RenderBadRequest(w, r, err)
		return
	}

	movie := &data.Movie{
		Title:   input.Title,
		Year:    input.Year,
		Runtime: data.Runtime(input.Runtime),
		Genres:  input.Genres,
	}

	v := validator.New()
	if data.ValidateMovie(v, movie); !v.Valid() {
		common.RenderFailedValidation(w, r, v.Errors)
		return
	}

	if err := app.models.Movies.Create(r.Context(), movie); err != nil {
		common.RenderInternalServerError(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/movies/%d", movie.ID))

	if err := app.writeJSON(w, http.StatusCreated, envelope{"movie": movie}, headers); err != nil {
		common.RenderInternalServerError(w, r, err)
	}
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	movieID, err := app.readIDParam(r)
	if err != nil {
		common.RenderNotFound(w, r)
		return
	}

	movie, err := app.models.Movies.Get(r.Context(), movieID)
	if err != nil {
		notFoundOrUnknownError(app, err, w, r)
		return
	}

	if err := app.writeJSON(w, http.StatusOK, envelope{"movies": movie}, nil); err != nil {
		common.RenderInternalServerError(w, r, err)
	}
}

func (app *application) updateMovieHandler(w http.ResponseWriter, r *http.Request) {
	movieID, err := app.readIDParam(r)
	if err != nil {
		common.RenderNotFound(w, r)
		return
	}

	movie, err := app.models.Movies.Get(r.Context(), movieID)
	if err != nil {
		notFoundOrUnknownError(app, err, w, r)
		return
	}

	var input struct {
		Title   string       `json:"title"`
		Year    int32        `json:"year"`
		Runtime data.Runtime `json:"runtime"`
		Genres  []string     `json:"genres"`
	}

	if err := app.readJSON(w, r, &input); err != nil {
		common.RenderBadRequest(w, r, err)
		return
	}

	movie.Title = input.Title
	movie.Year = input.Year
	movie.Runtime = input.Runtime
	movie.Genres = input.Genres

	v := validator.New()
	if data.ValidateMovie(v, movie); !v.Valid() {
		common.RenderInternalServerError(w, r, err)
		return
	}

	if err := app.models.Movies.Update(r.Context(), movie); err != nil {
		switch {
		case errors.Is(err, data.ErrStaleObject):
			common.RenderEditStaleRecord(w, r)
		default:
			common.RenderInternalServerError(w, r, err)
		}
		return
	}

	if err := app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil); err != nil {
		common.RenderInternalServerError(w, r, err)
	}
}

func (app *application) DeleteMovieHandler(w http.ResponseWriter, r *http.Request) {
	movieID, err := app.readIDParam(r)
	if err != nil {
		common.RenderNotFound(w, r)
		return
	}

	if err := app.models.Movies.Delete(r.Context(), movieID); err != nil {
		notFoundOrUnknownError(app, err, w, r)
		return
	}

	if err := app.writeJSON(w, http.StatusOK, envelope{"movie": "movie deleted"}, nil); err != nil {
		common.RenderInternalServerError(w, r, err)
	}
}

func (app *application) listMovieHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title   string
		Genres  []string
		Filters data.Filters
	}

	v := validator.New()
	qs := r.URL.Query()

	input.Title = common.ReadString(qs, "title", "")
	input.Genres = app.readCSV(qs, "genres", []string{})
	input.Filters.Page = app.readInt(qs, "page", 1, v)
	input.Filters.PageSize = app.readInt(qs, "page_size", 20, v)
	input.Filters.Sort = common.ReadString(qs, "sort", "id")
	input.Filters.SortSafelist = []string{"id", "title", "year", "runtime", "-id", "-title", "-year", "-runtime"}

	if data.ValidateFilters(v, input.Filters); !v.Valid() {
		common.RenderFailedValidation(w, r, v.Errors)
		return
	}

	if !v.Valid() {
		common.RenderFailedValidation(w, r, v.Errors)
		return
	}

	movies, metadata, err := app.models.Movies.GetAll(r.Context(), input.Title, input.Genres, input.Filters)
	if err != nil {
		common.RenderInternalServerError(w, r, err)
		return
	}

	if err := app.writeJSON(w, http.StatusOK, envelope{"movies": movies, "metadata": metadata}, nil); err != nil {
		common.RenderInternalServerError(w, r, err)
	}
}

func notFoundOrUnknownError(app *application, err error, w http.ResponseWriter, r *http.Request) {
	switch {
	case errors.Is(err, data.ErrRecordNotFound):
		common.RenderNotFound(w, r)
	default:
		common.RenderInternalServerError(w, r, err)
	}
}

func (app *application) readIDParam(r *http.Request) (int64, error) {
	str := r.PathValue("id")
	id, err := strconv.ParseInt(str, 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id param")
	}

	return id, nil
}

func openDB(cfg config.Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.DB.DSN)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.DB.MaxOpenConns)
	db.SetMaxIdleConns(cfg.DB.MaxIdleConns)
	duration, err := time.ParseDuration(cfg.DB.MaxIdleTime)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxIdleTime(duration)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	return db, nil
}

type envelope map[string]any

func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
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

func (app *application) readJSON(_ http.ResponseWriter, r *http.Request, dst any) error {
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

func (app *application) logError(r *http.Request, err error) {
	app.logger.Error().Err(err).
		Str("request_method", r.Method).
		Str("request_url", r.URL.String()).
		Send()
}

// use `debug.PrintStack()` when necessary
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

func (app *application) readCSV(qs url.Values, key string, defaultVal []string) []string {
	if s := qs.Get(key); s != "" {
		return strings.Split(s, ",")
	}

	return defaultVal
}

func (app *application) readInt(qs url.Values, key string, defaultVal int, v *validator.Validator) int {
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

// Note: return http.HandlerFunc instead of http.Handler
// so we can wrap handler func directly(not just router)
func (app *application) requireAuthenticatedUser(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := common.GetUserFromContext(r)
		if user.IsAnonymous() {
			common.RenderAuthenticationRequired(w, r)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (app *application) requireActivatedUser(next http.HandlerFunc) http.HandlerFunc {
	fn := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := common.GetUserFromContext(r)

		if !user.Activated() {
			common.RenderInactiveAccount(w, r)
			return
		}

		next.ServeHTTP(w, r)
	})

	return app.requireAuthenticatedUser(fn)
}

func (app *application) requirePermission(code string, next http.HandlerFunc) http.HandlerFunc {
	fn := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := common.GetUserFromContext(r)

		permissions, err := app.models.Permissions.GetAllForUser(r.Context(), user.ID)
		if err != nil {
			common.RenderInternalServerError(w, r, err)
			return
		}

		if !permissions.Include(code) {
			common.RenderNotPermitted(w, r)
			return
		}

		next.ServeHTTP(w, r)
	})

	return app.requireActivatedUser(fn)
}

func (app *application) serve() error {
	mux := app.routes()

	var handler http.Handler = mux

	// the middlewares order matters
	handler = middleware.Authenticate(handler, app.models, app.config.Debug)
	handler = middleware.RateLimit(handler, app.config.Limiter.Enabled, app.config.Limiter.RPS, app.config.Limiter.Burst)
	handler = middleware.EnableCORS(handler, app.config.CORS.TrustedOrigins)
	handler = middleware.RecoverPanic(handler)
	handler = middleware.CollectMetrics(handler)

	srv := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", app.config.IP, app.config.Port),
		Handler:      handler,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	shutdownError := make(chan error)

	go func() {
		quit := make(chan os.Signal, 1)

		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		s := <-quit // block until signal received

		app.logger.Info().
			Str("signal", s.String()).
			Str("service_type", "api server").
			Msg("shutting down")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			shutdownError <- srv.Shutdown(ctx)
		}

		app.logger.Info().Str("address", srv.Addr).Msg("completing background tasks")

		app.wg.Wait()

		shutdownError <- nil
	}()

	app.logger.Info().
		Str("service_type", "api server").
		Str("env", app.config.Env).
		Str("address", srv.Addr).
		Str("approot", approot.Root).
		Interface("config", configStructToMap(app.config)).
		Msg("server started")

	if err := srv.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			return err
		}

		if err := <-shutdownError; err != nil {
			return err
		}

		app.logger.Info().Msg("server shutdown")
	}

	return nil
}

func (app *application) registerUserHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := app.readJSON(w, r, &input); err != nil {
		common.RenderBadRequest(w, r, err)
		return
	}

	user := &data.User{
		Name:   input.Name,
		Email:  input.Email,
		Status: "pending",
	}

	if err := user.Password.Set(input.Password); err != nil {
		common.RenderInternalServerError(w, r, err)
		return
	}

	v := validator.New()

	if data.ValidateUser(v, user); !v.Valid() {
		common.RenderFailedValidation(w, r, v.Errors)
		return
	}

	if err := app.models.Users.Create(r.Context(), user); err != nil {
		switch {
		case errors.Is(err, data.ErrDuplicatedEmail):
			v.AddError("email", "email already taken") // FIXME: take care of "preventing enumeration attack" when necessary
			common.RenderFailedValidation(w, r, v.Errors)
		default:
			common.RenderInternalServerError(w, r, err)
		}

		return
	}

	if err := app.models.Permissions.AddForUser(r.Context(), user.ID, "movies:read"); err != nil {
		common.RenderInternalServerError(w, r, err)
		return
	}

	token, err := app.models.Tokens.New(r.Context(), user.ID, 3*24*time.Hour, data.ScopeActivation)
	if err != nil {
		common.RenderInternalServerError(w, r, err)
		return
	}

	app.runInBackground(func() {
		app.sendEmail(user, token)
	})

	if err := app.writeJSON(w, http.StatusAccepted, envelope{"user": user}, nil); err != nil {
		common.RenderInternalServerError(w, r, err)
	}
}

func (app *application) sendEmail(user *data.User, token *data.Token) {
	logger := app.logger.With().Int64("user_id", user.ID).Logger()

	data := map[string]any{
		"userID":          user.ID,
		"activationToken": token.Plaintext,
	}

	logger.Info().Msg("send email start")
	startTime := time.Now()

	if err := app.mailer.Send(app.config.SMTP.Sender, user.Email, "user_welcome.tmpl", data); err != nil {
		logger.Error().Err(err).Msg("send email failed")
		return
	}

	logger.Info().Str("duration", time.Since(startTime).String()).Msg("send email done")
}

func (app *application) runInBackground(fn func()) {
	app.wg.Add(1)

	go func() {
		defer app.wg.Done()

		defer func() {
			if err := recover(); err != nil {
				app.logger.Error().Err(fmt.Errorf("%s", err)).Msg("recover panic failed")
			}
		}()

		fn()
	}()
}

// struct -> []byte -> map
func configStructToMap(cfg config.Config) map[string]any {
	configInfo, err := json.Marshal(cfg)
	assert.Assert(err == nil)

	configMap := make(map[string]any)
	err = json.Unmarshal(configInfo, &configMap)
	assert.Assert(err == nil)

	return configMap
}

func (app *application) activateUserHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		TokenPlaintext string `json:"token"`
	}

	if err := app.readJSON(w, r, &input); err != nil {
		common.RenderBadRequest(w, r, err)
		return
	}

	v := validator.New()

	if data.ValidateTokenPlaintext(v, input.TokenPlaintext); !v.Valid() {
		common.RenderFailedValidation(w, r, v.Errors)
		return
	}

	user, err := app.models.Users.GetByToken(r.Context(), data.ScopeActivation, input.TokenPlaintext)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			v.AddError("token", "invalid or expired activation token")
			common.RenderFailedValidation(w, r, v.Errors)
		default:
			common.RenderInternalServerError(w, r, err)
		}

		return
	}

	user.Status = "activated"

	if err := app.models.Users.Update(r.Context(), user); err != nil {
		switch {
		case errors.Is(err, data.ErrStaleObject):
			common.RenderEditStaleRecord(w, r)
		default:
			common.RenderInternalServerError(w, r, err)
		}

		return
	}

	if err := app.models.Tokens.DeleteAllForUser(r.Context(), data.ScopeActivation, user.ID); err != nil {
		common.RenderInternalServerError(w, r, err)
		return
	}

	if err := app.writeJSON(w, http.StatusOK, envelope{"user": user}, nil); err != nil {
		common.RenderInternalServerError(w, r, err)
	}
}

func (app *application) createAuthenticationTokenHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := app.readJSON(w, r, &input); err != nil {
		common.RenderBadRequest(w, r, err)
		return
	}

	v := validator.New()
	data.ValidateEmail(v, input.Email)
	data.ValidatePasswordPlaintext(v, input.Password)

	if !v.Valid() {
		common.RenderFailedValidation(w, r, v.Errors)
		return
	}

	ctx := r.Context()
	user, err := app.models.Users.GetByEmail(ctx, input.Email)
	if err != nil {
		notFoundOrUnknownError(app, err, w, r)
		return
	}

	match, err := user.Password.Matches(input.Password)
	if err != nil {
		common.RenderInternalServerError(w, r, err)
		return
	}

	if !match {
		common.RenderInvlidCredentials(w, r)
		return
	}

	token, err := app.models.Tokens.New(r.Context(), user.ID, 24*time.Hour, data.ScopeAuthentication)
	if err != nil {
		common.RenderInternalServerError(w, r, err)
		return
	}

	if err := app.writeJSON(w, http.StatusCreated, envelope{"authentication_token": token}, nil); err != nil {
		common.RenderInternalServerError(w, r, err)
	}
}
