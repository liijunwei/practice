package main

import (
	"crypto/sha1"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"golang-practices/shorturl/sqlcdb"
	"log"
	"math/big"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"

	_ "github.com/mattn/go-sqlite3" // SQLite driver
)

var once sync.Once

const base62Chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// simple shorturl app with sqlite3+sqlc
//
// features:
// 1. create short url
// 2. redirect to original url
// 3. analyse short url access
// 4. test single sqlite3 db max qps(HOW?)
//
// go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
// guide: https://docs.sqlc.dev/en/latest/tutorials/getting-started-sqlite.html
//
// rm /tmp/shorturl-app.db; sqlc generate && go run .
func main() {
	db := initDB()
	defer db.Close()

	queries := sqlcdb.New(db)

	// http.HandleFunc("GET /", indexHandler(queries, db))
	http.HandleFunc("POST /shorturl", createHandler(queries, db))
	http.HandleFunc("GET /shorturl/{code}", redirectHandler(queries, db))

	log.Println("Server is running at http://localhost:8080")
	boom(http.ListenAndServe(":8080", nil))
}

type ShortURL struct {
	ID        int64  `json:"id"`
	Original  string `json:"original"`
	Short     string `json:"short"`
	CreatedAt string `json:"created_at"`
}

func toShortURL(entity sqlcdb.Shorturl) ShortURL {
	return ShortURL{
		ID:        entity.ID,
		Original:  entity.Original,
		Short:     fmt.Sprintf("http://localhost:8080/shorturl/%s", entity.Shorturl),
		CreatedAt: entity.CreatedAt.Format(time.RFC3339),
	}
}

func initDB() *sql.DB {
	db, err := sql.Open("sqlite3", "/tmp/shorturl-app.db")
	boom(err, "failed to open database")

	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)

	_, err = db.Exec("PRAGMA journal_mode=WAL;")
	boom(err, "failed to set WAL mode")

	schema, err := os.ReadFile(filepath.Join(baseDir(), "schema.sql"))
	boom(err, "failed to read schema.sql")

	_, err = db.Exec(string(schema))
	boom(err, "failed to execute schema")

	boom(db.Ping())

	// go func() {
	// 	for range time.Tick(1 * time.Second) {
	// 		data, _ := json.Marshal(db.Stats())
	// 		log.Println("db stats:", string(data))
	// 	}
	// }()

	return db
}

func baseDir() string {
	var dir string

	once.Do(func() {
		_, f, _, _ := runtime.Caller(0)
		dir = filepath.Dir(f)
	})

	return dir
}

func indexHandler(db *sqlcdb.Queries, _ *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		entities, err := db.ListShorturls(r.Context())
		boom(err, "failed to list shorturls")

		urls := make([]ShortURL, 0, len(entities))

		for _, entity := range entities {
			urls = append(urls, toShortURL(entity))
		}

		WriteResponseJSON(w, http.StatusOK, Envelope{"shorturls": urls}, nil)
	}
}

func createHandler(db *sqlcdb.Queries, sqldb *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		assert(r.Method == "POST", "invalid method")

		var input struct {
			Original string `json:"original"`
		}

		err := json.NewDecoder(r.Body).Decode(&input)
		assert(err == nil, "failed to decode json")

		ctx := r.Context()

		tx, err := sqldb.BeginTx(ctx, nil)
		boom(err)

		defer func() {
			if err := tx.Rollback(); err != nil && !errors.Is(err, sql.ErrTxDone) {
				log.Println("failed to rollback tx:", err)
			}
		}()

		db = db.WithTx(tx)

		result, err := db.OriginalExists(ctx, input.Original)
		if err != nil {
			log.Println("failed to check original url exists:", err)
			WriteResponseJSON(w, http.StatusInternalServerError, Envelope{"error": err.Error()}, nil)
			return
		}

		if exists := result == 1; exists {
			log.Println("original url exists", input.Original)
			WriteResponseJSON(w, http.StatusBadRequest, Envelope{"error": "original url exists"}, nil)
			return
		}

		created, err := db.CreateShorturl(ctx, sqlcdb.CreateShorturlParams{
			Original: input.Original,
			Shorturl: genShorturl(input.Original),
		})

		if err != nil {
			log.Println("failed to create shorturl:", err)
			WriteResponseJSON(w, http.StatusInternalServerError, Envelope{"error": err.Error()}, nil)
			return
		}

		if err = tx.Commit(); err != nil {
			log.Println("failed to commit tx:", err)
			WriteResponseJSON(w, http.StatusInternalServerError, Envelope{"error": err.Error()}, nil)
			return
		}

		WriteResponseJSON(w, http.StatusCreated, Envelope{"shorturl": toShortURL(created)}, nil)
	}
}

func redirectHandler(db *sqlcdb.Queries, _ *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		assert(r.Method == "GET", "invalid method")

		code := r.PathValue("code")
		if code == "" {
			WriteResponseJSON(w, http.StatusBadRequest, Envelope{"error": "code is required"}, nil)
			return
		}

		ctx := r.Context()

		shortURL, err := db.GetOriginalByShort(ctx, code)
		if err != nil {
			log.Println("failed to query by short code:", err)
			WriteResponseJSON(w, http.StatusInternalServerError, Envelope{"error": "failed to query by short code"}, nil)
		}

		http.Redirect(w, r, shortURL.Original, http.StatusFound)
	}
}

func genShorturl(original string) string {
	h := sha1.New()
	h.Write([]byte(original))
	bs := h.Sum(nil)
	num := new(big.Int).SetBytes(bs)

	return base62Encode(num)[:8]
}

// TODO verify its correctness
func base62Encode(num *big.Int) string {
	base := big.NewInt(62)
	zero := big.NewInt(0)
	encoded := ""

	for num.Cmp(zero) > 0 {
		mod := new(big.Int)
		num.DivMod(num, base, mod)
		encoded = string(base62Chars[mod.Int64()]) + encoded
	}

	return encoded
}

func boom(e error, msg ...string) {
	if e != nil {
		_, filename, line, _ := runtime.Caller(1)
		log.Println(filename, line, strings.Join(msg, " "), e)

		panic(e)
	}
}

func assert(ok bool, msg ...string) {
	if !ok {
		log.Println(strings.Join(msg, " "), "assert failed")
		panic("assert failed")
	}
}

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
