package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"greenlight/internal/assert"
	"greenlight/internal/data"
	"greenlight/internal/jsonlog"
	"greenlight/internal/mailer"
	approot "greenlight/internal/projectroot"
	"greenlight/internal/validator"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"runtime/debug"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"

	_ "github.com/lib/pq"
	"golang.org/x/time/rate"
)

const version = "1.0.0"

func main() {
	var cfg config
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
	flag.StringVar(&cfg.SMTP.Username, "smtp-username", "placeholder", "SMTP username")
	flag.StringVar(&cfg.SMTP.Password, "smtp-password", "placeholder", "SMTP password")
	flag.StringVar(&cfg.SMTP.Sender, "smtp-sender", "greenlight-admin@example.com", "SMTP sender")
	flag.Parse()

	run(cfg)
}

func run(cfg config) {
	logger := jsonlog.New(os.Stdout, jsonlog.LevelInfo)

	db, err := openDB(cfg)
	if err != nil {
		logger.PrintFatal(err, nil)
	}

	logger.PrintInfo("database connection pool established", nil)

	defer db.Close()

	app := &application{
		config: cfg,
		logger: logger,
		models: data.NewModels(db),
		mailer: mailer.New(cfg.SMTP.Host, cfg.SMTP.Port, cfg.SMTP.Username, cfg.SMTP.Password, cfg.SMTP.Sender),
	}

	if err := app.serve(); err != nil {
		logger.PrintFatal(err, nil)
	}
}

type config struct {
	IP      string        `json:"ip"`
	Port    int           `json:"port"`
	Env     string        `json:"env"`
	DB      dbConfig      `json:"db"`
	Limiter limiterConfig `json:"limiter"`
	Debug   bool          `json:"debug"`
	SMTP    smtp          `json:"smtp"`
}

type dbConfig struct {
	DSN          string `json:"-"` // Data Source Name (DSN)
	MaxOpenConns int    `json:"max_open_conns"`
	MaxIdleConns int    `json:"max_idle_conns"`
	MaxIdleTime  string `json:"max_idle_time"`
}

type limiterConfig struct {
	RPS     float64 `json:"rps"`
	Burst   int     `json:"burst"`
	Enabled bool    `json:"enabled"`
}

type smtp struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"-"`
	Password string `json:"-"`
	Sender   string `json:"sender"`
}

type application struct {
	config config
	logger *jsonlog.Logger
	models data.Models
	mailer mailer.Mailer
	wg     sync.WaitGroup
}

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.notFoundResponse)

	mux.HandleFunc("GET /v1/healthcheck", app.healthcheckHandler)
	mux.HandleFunc("POST /v1/movies", app.createMovieHandler)
	mux.HandleFunc("GET /v1/movies/{id}", app.showMovieHandler)
	mux.HandleFunc("PUT /v1/movies/{id}", app.updateMovieHandler)
	mux.HandleFunc("DELETE /v1/movies/{id}", app.DeleteMovieHandler)
	mux.HandleFunc("GET /v1/movies", app.listMovieHandler)
	mux.HandleFunc("POST /v1/users", app.registerUserHandler)

	return mux
}

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	env := envelope{
		"status": "available",
		"system_info": map[string]any{
			"environment": app.config.Env,
			"version":     version,
		},
	}

	if err := app.writeJSON(w, http.StatusOK, env, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title   string   `json:"title"`
		Year    int32    `json:"year"`
		Runtime int32    `json:"runtime"`
		Genres  []string `json:"genres"`
	}

	if err := app.readJSON(w, r, &input); err != nil {
		app.badRequestResponse(w, r, err)
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
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	if err := app.models.Movies.Insert(movie); err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/movies/%d", movie.ID))

	if err := app.writeJSON(w, http.StatusCreated, envelope{"movie": movie}, headers); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	movieID, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	movie, err := app.models.Movies.Get(movieID)
	if err != nil {
		notFoundOrUnknownError(app, err, w, r)
		return
	}

	if err := app.writeJSON(w, http.StatusOK, envelope{"movies": movie}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) updateMovieHandler(w http.ResponseWriter, r *http.Request) {
	movieID, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	movie, err := app.models.Movies.Get(movieID)
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
		app.badRequestResponse(w, r, err)
		return
	}

	movie.Title = input.Title
	movie.Year = input.Year
	movie.Runtime = input.Runtime
	movie.Genres = input.Genres

	v := validator.New()
	if data.ValidateMovie(v, movie); !v.Valid() {
		app.serverErrorResponse(w, r, err)
		return
	}

	if err := app.models.Movies.Update(movie); err != nil {
		switch {
		case errors.Is(err, data.ErrStaleObject):
			app.editStaleRecordResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	if err := app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) DeleteMovieHandler(w http.ResponseWriter, r *http.Request) {
	movieID, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	if err := app.models.Movies.Delete(movieID); err != nil {
		notFoundOrUnknownError(app, err, w, r)
		return
	}

	if err := app.writeJSON(w, http.StatusOK, envelope{"movie": "movie deleted"}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
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

	input.Title = app.readString(qs, "title", "")
	input.Genres = app.readCSV(qs, "genres", []string{})
	input.Filters.Page = app.readInt(qs, "page", 1, v)
	input.Filters.PageSize = app.readInt(qs, "page_size", 20, v)
	input.Filters.Sort = app.readString(qs, "sort", "id")
	input.Filters.SortSafelist = []string{"id", "title", "year", "runtime", "-id", "-title", "-year", "-runtime"}

	if data.ValidateFilters(v, input.Filters); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	movies, metadata, err := app.models.Movies.GetAll(input.Title, input.Genres, input.Filters)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	if err := app.writeJSON(w, http.StatusOK, envelope{"movies": movies, "metadata": metadata}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func notFoundOrUnknownError(app *application, err error, w http.ResponseWriter, r *http.Request) {
	switch {
	case errors.Is(err, data.ErrRecordNotFound):
		app.notFoundResponse(w, r)
	default:
		app.serverErrorResponse(w, r, err)
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

func openDB(cfg config) (*sql.DB, error) {
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
	app.logger.PrintError(err, map[string]any{
		"request_method": r.Method,
		"request_url":    r.URL.String(),
	})
}

func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message any) {
	env := envelope{"error": message}
	if err := app.writeJSON(w, status, env, nil); err != nil {
		app.logError(r, err)
		w.WriteHeader(500)
	}
}

func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logError(r, err)
	message := "server failed unexpectedly"

	if app.config.Debug {
		details := map[string]any{
			"message": message,
			"error":   err.Error(),
			"traces":  sanitizedDebugTraces(),
		}

		app.errorResponse(w, r, http.StatusServiceUnavailable, details)
		return
	}

	app.errorResponse(w, r, http.StatusServiceUnavailable, message)
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

func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "resource not found"
	app.errorResponse(w, r, http.StatusNotFound, message)
}

func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.errorResponse(w, r, http.StatusBadRequest, err.Error())
}

func (app *application) failedValidationResponse(w http.ResponseWriter, r *http.Request, errors map[string]string) {
	app.errorResponse(w, r, http.StatusUnprocessableEntity, errors)
}

func (app *application) editStaleRecordResponse(w http.ResponseWriter, r *http.Request) {
	message := "unable to update a stale object, please try again"
	app.errorResponse(w, r, http.StatusConflict, message)
}

func (app *application) rateLimitExceededResponse(w http.ResponseWriter, r *http.Request) {
	message := "rate limit exceeded"
	app.errorResponse(w, r, http.StatusTooManyRequests, message)
}

func (app *application) readString(qs url.Values, key string, defaultVal string) string {
	if s := qs.Get(key); s != "" {
		return s
	}

	return defaultVal
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

func (app *application) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()

		app.logger.PrintInfo("request start", map[string]any{
			"request_method": r.Method,
			"request_url":    r.URL.String(),
		})

		defer func() {
			app.logger.PrintInfo("request end", map[string]any{
				"duration": time.Since(now).String(),
			})
		}()

		defer func() {
			if err := recover(); err != nil {
				w.Header().Set("Connection", "close")
				app.serverErrorResponse(w, r, fmt.Errorf("%s", err))
			}
		}()

		next.ServeHTTP(w, r)
	})
}

func (app *application) rateLimit(next http.Handler) http.Handler {
	type client struct {
		limiter  *rate.Limiter
		lastSeen time.Time
	}

	var mu sync.Mutex
	var clients = make(map[string]*client) // works for single machine

	go func() {
		ticker := time.NewTicker(1 * time.Minute)

		for range ticker.C {
			mu.Lock()

			app.logger.PrintInfo("checking client map...", nil)

			for ip, client := range clients {
				if time.Since(client.lastSeen) > 3*time.Minute {
					delete(clients, ip)
				}
			}

			mu.Unlock()
		}
	}()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			app.serverErrorResponse(w, r, err)
			return
		}

		mu.Lock()

		if _, ok := clients[ip]; !ok {
			clients[ip] = &client{limiter: rate.NewLimiter(rate.Limit(app.config.Limiter.RPS), app.config.Limiter.Burst)}
		}

		clients[ip].lastSeen = time.Now()

		if !clients[ip].limiter.Allow() {
			mu.Unlock()

			app.logger.PrintInfo("rate_limit triggered", map[string]any{
				"request_method": r.Method,
				"request_url":    r.URL.String(),
				"user_agent":     r.UserAgent(),
			})

			app.rateLimitExceededResponse(w, r)

			return
		}

		mu.Unlock()

		next.ServeHTTP(w, r)
	})
}

func (app *application) serve() error {
	mux := app.routes()

	var handler http.Handler = mux

	// the middlewares order matters
	if app.config.Limiter.Enabled {
		handler = app.rateLimit(handler)
	}
	handler = app.recoverPanic(handler)

	srv := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", app.config.IP, app.config.Port),
		Handler:      handler,
		ErrorLog:     log.New(app.logger, "", 0),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	shutdownError := make(chan error)

	go func() {
		quit := make(chan os.Signal, 1)

		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		s := <-quit // block until signal received

		app.logger.PrintInfo("shutting down", map[string]any{
			"service_type": "api server",
			"signal":       s.String(),
		})

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			shutdownError <- srv.Shutdown(ctx)
		}

		app.logger.PrintInfo("completing background tasks", map[string]any{
			"addr": srv.Addr,
		})

		app.wg.Wait()

		shutdownError <- nil
	}()

	app.logger.PrintInfo("server started", map[string]any{
		"env":     app.config.Env,
		"address": srv.Addr,
		"config":  configStructToMap(app.config),
		"approot": approot.Root,
	})

	if err := srv.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			return err
		}

		if err := <-shutdownError; err != nil {
			return err
		}

		app.logger.PrintInfo("server shutdown", nil)
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
		app.badRequestResponse(w, r, err)
		return
	}

	user := &data.User{
		Name:   input.Name,
		Email:  input.Email,
		Status: "pending",
	}

	if err := user.Password.Set(input.Password); err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	v := validator.New()

	if data.ValidateUser(v, user); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	if err := app.models.Users.Insert(user); err != nil {
		switch {
		case errors.Is(err, data.ErrDuplicatedEmail):
			v.AddError("email", "email already taken") // FIXME: take care of "preventing enumeration attack" when necessary
			app.failedValidationResponse(w, r, v.Errors)
		default:
			app.serverErrorResponse(w, r, err)
		}

		return
	}

	token, err := app.models.Tokens.New(user.ID, 3*24*time.Hour, data.ScopeActivation)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.runInBackground(func() {
		sendEmail(app.logger, app.mailer, user, token)
	})

	if err := app.writeJSON(w, http.StatusAccepted, envelope{"user": user}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func sendEmail(logger *jsonlog.Logger, mailer mailer.Mailer, user *data.User, token *data.Token) {
	detail := map[string]any{
		"user_id": user.ID,
	}

	data := map[string]any{
		"userID":          user.ID,
		"activationToken": token.Plaintext,
	}

	logger.PrintInfo("send email: start", detail)
	startTime := time.Now()

	if err := mailer.Send(user.Email, "user_welcome.tmpl", data); err != nil {
		detail["error_message"] = "send email: fail"
		logger.PrintError(err, detail)
		return
	}

	detail["duration"] = time.Since(startTime).String()
	logger.PrintInfo("send email: done", detail)
}

func (app *application) runInBackground(fn func()) {
	app.wg.Add(1)

	go func() {
		defer app.wg.Done()

		defer func() {
			if err := recover(); err != nil {
				app.logger.PrintError(fmt.Errorf("%s", err), nil)
			}
		}()

		fn()
	}()
}

// struct -> []byte -> map
func configStructToMap(config config) map[string]any {
	configInfo, err := json.Marshal(config)
	assert.Assert(err == nil)

	configMap := make(map[string]any)
	err = json.Unmarshal(configInfo, &configMap)
	assert.Assert(err == nil)

	return configMap
}
