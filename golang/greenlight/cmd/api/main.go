package main

import (
	"flag"
	"fmt"
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
	fmt.Fprintf(w, "status: available\n")
	fmt.Fprintf(w, "environment: %s\n", app.config.env)
	fmt.Fprintf(w, "version: %s\n", version)
}

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/healthcheck", app.healthcheckHandler)

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
