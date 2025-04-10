package main

import (
	"context"
	"database/sql"
	"errors"
	"expvar"
	"flag"
	"fmt"
	"greenlight/internal/approot"
	"greenlight/internal/assert"
	"greenlight/internal/common"
	"greenlight/internal/config"
	"greenlight/internal/data"
	"greenlight/internal/dbconn"
	"greenlight/internal/mailer"
	"greenlight/internal/rest"
	"greenlight/internal/rest/middleware"
	"greenlight/internal/rest/userfacing/moviesapi"
	"greenlight/internal/rest/userfacing/tokensapi"
	"greenlight/internal/rest/userfacing/usersapi"
	"greenlight/internal/sse"
	"net/http"
	"os"
	"os/signal"
	"runtime"
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
	logger := log.Logger.With().Str("service", "greenlight-api").Logger()

	level, err := zerolog.ParseLevel(cfg.LogConfig.Level)
	assert.NoError(err)
	assert.Assert(level != zerolog.NoLevel)

	zerolog.SetGlobalLevel(level)
	logger.Level(level)

	ctx := context.Background()
	ctx = logger.WithContext(ctx)

	db, err := dbconn.NewDB(cfg)
	assert.NoError(err, "foo", "bar")
	defer db.Close()

	zerolog.Ctx(ctx).Info().Msg("database connection pool established")

	models, err := data.SetupModels(cfg, db)
	assert.NoError(err)

	app := &application{
		config: cfg,
		models: models,
		mailer: mailer.SetupMailer(cfg.SMTP),
	}

	expvar.NewString("version").Set(version)
	expvar.Publish("goroutines", expvar.Func(func() any { return runtime.NumGoroutine() }))
	expvar.Publish("database", expvar.Func(func() any { return db.Stats() }))
	expvar.Publish("timestamp", expvar.Func(func() any { return time.Now().Unix() }))

	if err := app.serve(ctx, db); err != nil {
		logger.Fatal().Err(err).Msg("unexpected server error")
	}
}

type application struct {
	config config.Config
	models data.Models
	mailer mailer.Mailer
	wg     sync.WaitGroup
}

func (app *application) routes(db *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /v1/healthcheck", rest.HealthcheckHandler(app.config.Env, version))
	mux.HandleFunc("GET /v1/dbstat", sse.Wrap(rest.StreamDatabaseStatHandler(db)))
	mux.HandleFunc("POST /v1/movies", requirePermission(app.models, "movies:write", moviesapi.CreateMovieHandler(app.models)))
	mux.HandleFunc("GET /v1/movies/{id}", requirePermission(app.models, "movies:read", moviesapi.GetMovieDetailHandler(app.models)))
	mux.HandleFunc("PUT /v1/movies/{id}", requirePermission(app.models, "movies:write", moviesapi.UpdateMovieDetailHandler(app.models)))
	mux.HandleFunc("DELETE /v1/movies/{id}", requirePermission(app.models, "movies:write", moviesapi.DeleteMovieHandler(app.models)))
	mux.HandleFunc("GET /v1/movies", requirePermission(app.models, "movies:read", moviesapi.GetMovieList(app.models)))
	mux.HandleFunc("POST /v1/users", usersapi.CreateUserHandler(app.models, app.mailer, app.config.SMTP.Sender))
	mux.HandleFunc("PUT /v1/users/activated", usersapi.ActivateUserHandler(app.models))
	mux.HandleFunc("POST /v1/tokens/authentication", tokensapi.CreateAuthenticationTokenHandler(app.models))
	mux.Handle("GET /debug/vars", expvar.Handler())

	return mux
}

// TODO maybe use chirouter is more obvious
//
// Note: return http.HandlerFunc instead of http.Handler
// so we can wrap handler func directly(not just router)
func requireAuthenticatedUser(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := common.GetUserFromContext(r)
		if user.IsAnonymous() {
			common.RenderAuthenticationRequired(w, r)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func requireActivatedUser(next http.HandlerFunc) http.HandlerFunc {
	fn := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := common.GetUserFromContext(r)

		if !user.Activated() {
			common.RenderInactiveAccount(w, r)
			return
		}

		next.ServeHTTP(w, r)
	})

	return requireAuthenticatedUser(fn)
}

func requirePermission(models data.Models, code string, next http.HandlerFunc) http.HandlerFunc {
	fn := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := common.GetUserFromContext(r)

		permissions, err := models.Permissions.GetAllForUser(r.Context(), user.ID)
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

	return requireActivatedUser(fn)
}

func (app *application) serve(ctx context.Context, db *sql.DB) error {
	mux := app.routes(db)

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

		zerolog.Ctx(ctx).Info().
			Str("signal", s.String()).
			Str("service_type", "api server").
			Msg("shutting down")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			shutdownError <- srv.Shutdown(ctx)
		}

		zerolog.Ctx(ctx).Info().Str("address", srv.Addr).Msg("completing background tasks")

		app.wg.Wait()

		shutdownError <- nil
	}()

	zerolog.Ctx(ctx).Info().
		Str("service_type", "api server").
		Str("env", app.config.Env).
		Str("address", srv.Addr).
		Str("approot", approot.Root).
		Interface("appconfig", app.config).
		Msg("server started")

	if err := srv.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			return err
		}

		if err := <-shutdownError; err != nil {
			return err
		}

		zerolog.Ctx(ctx).Info().Msg("server shutdown")
	}

	return nil
}
