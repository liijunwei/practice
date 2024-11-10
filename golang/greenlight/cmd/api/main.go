package main

import (
	"errors"
	"flag"
	"fmt"
	"greenlight/internal/data"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

const version = "1.0.0"

func main() {
	var cfg config
	flag.StringVar(&cfg.env, "env", "development", "environment(development|staging|production)")
	flag.StringVar(&cfg.ip, "ip", "localhost", "the server ip address")
	flag.IntVar(&cfg.port, "port", 4000, "api server port")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	app := &application{
		config: cfg,
		logger: *logger,
	}

	mux := app.routes()

	server := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", cfg.ip, cfg.port),
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("starting %s server on %s", cfg.env, server.Addr)
	err := server.ListenAndServe()
	logger.Fatal(err)
}

type config struct {
	ip   string
	port int
	env  string
}

type application struct {
	config config
	logger log.Logger
}

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	env := envelope{
		"status": "available",
		"system_info": map[string]string{
			"environment": app.config.env,
			"version":     version,
		},
	}

	if err := app.writeJSON(w, http.StatusOK, env, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.notFoundResponse)

	mux.HandleFunc("GET /v1/healthcheck", app.healthcheckHandler)
	mux.HandleFunc("POST /v1/movies", app.createMovieHandler)
	mux.HandleFunc("GET /v1/movies/{id}", app.showMovieHandler)

	// borrowed from: https://github.com/benhoyt/go-routing/blob/master/stdlib/route.go
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /contact", contact)
	mux.HandleFunc("GET /api/widgets", apiGetWidgets)
	mux.HandleFunc("POST /api/widgets", apiCreateWidget)
	mux.HandleFunc("POST /api/widgets/{slug}", apiUpdateWidget)
	mux.HandleFunc("POST /api/widgets/{slug}/parts", apiCreateWidgetPart)
	mux.HandleFunc("POST /api/widgets/{slug}/parts/{id}/update", apiUpdateWidgetPart)
	mux.HandleFunc("POST /api/widgets/{slug}/parts/{id}/delete", apiDeleteWidgetPart)
	mux.HandleFunc("GET /{slug}", widgetGet)
	mux.HandleFunc("GET /{slug}/admin", widgetAdmin)
	mux.HandleFunc("POST /{slug}/image", widgetImage)

	return mux
}

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "createMovieHandler\n")
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	movieID, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	movie := data.Movie{
		ID:        movieID,
		CreatedAt: time.Now(),
		Title:     "BOPE",
		Year:      2007,
		Runtime:   114,
		Genres:    []string{"crime", "war"},
		Version:   1,
	}

	if err := app.writeJSON(w, http.StatusOK, envelope{"movies": movie}, nil); err != nil {
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

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "home\n")
}

func contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "contact\n")
}

func apiGetWidgets(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "apiGetWidgets\n")
}

func apiCreateWidget(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "apiCreateWidget\n")
}

func apiUpdateWidget(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	fmt.Fprintf(w, "apiUpdateWidget %s\n", slug)
}

func apiCreateWidgetPart(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	fmt.Fprintf(w, "apiCreateWidgetPart %s\n", slug)
}

func apiUpdateWidgetPart(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "apiUpdateWidgetPart %s %d\n", slug, id)
}

func apiDeleteWidgetPart(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "apiDeleteWidgetPart %s %d\n", slug, id)
}

func widgetGet(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	fmt.Fprintf(w, "widget %s\n", slug)
}

func widgetAdmin(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	fmt.Fprintf(w, "widgetAdmin %s\n", slug)
}

func widgetImage(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	fmt.Fprintf(w, "widgetImage %s\n", slug)
}
