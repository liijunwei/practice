package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"golang-practices/bank1/sqlcdb"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"

	_ "github.com/mattn/go-sqlite3" // SQLite driver
)

var once sync.Once

// simple bank app with sqlite3+sqlc
//
// features:
// 1. deposit
// 2. withdraw
// 3. transfer
//
// go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
// guide: https://dev.to/techschoolguru/series/7172
//
// rm /tmp/bank-app.db; sqlc generate && go run .
func main() {
	db := initDB()
	defer db.Close()

	queries := sqlcdb.New(db)
	service := NewAccountService(db)

	http.HandleFunc("GET /", indexHandler(queries))
	http.HandleFunc("POST /register", registerUserHandler(service))
	http.HandleFunc("POST /transfer", internalTransferHandler(service))

	// go func() {
	// 	for {
	// 		time.Sleep(1000 * time.Millisecond)
	// 		log.Println("gorutine", runtime.NumGoroutine())
	// 	}
	// }()

	go func() {
		log.Println("start pprof server", "http://localhost:6060/debug/pprof")
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	log.Println("start bank server", "http://localhost:8080")
	boom(http.ListenAndServe(":8080", nil))
}

func initDB() *sql.DB {
	db, err := sql.Open("sqlite3", "/tmp/bank-app.db")
	boom(err, "failed to open database")

	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)
	db.SetConnMaxLifetime(time.Minute * 5)

	_, err = db.Exec("PRAGMA journal_mode=WAL;")
	boom(err, "failed to set WAL mode")

	schema, err := os.ReadFile(filepath.Join(baseDir(), "schema.sql"))
	boom(err, "failed to read schema.sql")

	_, err = db.Exec(string(schema))
	boom(err, "failed to execute schema")

	seed, err := os.ReadFile(filepath.Join(baseDir(), "seed.sql"))
	boom(err, "failed to read seed.sql")

	_, err = db.Exec(string(seed))
	boom(err, "failed to seed data")

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

func indexHandler(db *sqlcdb.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		accounts, err := db.GetAllAccounts(ctx)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to list accounts: %v", err), http.StatusInternalServerError)
			return
		}

		html := `<h1>XXX Bank</h1>
		<h2>All Accounts</h2>
		<table border="1" style="border-collapse: collapse; width: 80%;">
			<thead>
				<tr>
					<th>Account ID</th>
					<th>User Email</th>
					<th>Currency</th>
					<th>Available Balance</th>
					<th>Created At</th>
				</tr>
			</thead>
			<tbody>`

		for _, account := range accounts {
			html += fmt.Sprintf(`
				<tr>
					<td>%s</td>
					<td>%s</td>
					<td>%s</td>
					<td>%.2f</td>
					<td>%s</td>
				</tr>`,
				account.AccountID,
				account.Email,
				account.Currency,
				account.Available,
				account.CreatedAt.Format(time.RFC3339),
			)
		}

		html += `
			</tbody>
		</table>`

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte(html))
	}
}

func registerUserHandler(service *AccountService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		type RegisterRequest struct {
			Username   string   `json:"username"`
			Password   string   `json:"password"`
			Email      string   `json:"email"`
			Currencies []string `json:"currencies"`
		}

		var req RegisterRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		if req.Username == "" || req.Password == "" || req.Email == "" {
			http.Error(w, "Missing required fields", http.StatusBadRequest)
			return
		}

		ctx := r.Context()

		user, err := service.Register(ctx, req.Username, req.Password, req.Email, req.Currencies, 100)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to register user: %v", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		data, err := json.Marshal(user)
		if err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}

		w.Write(data)
	}
}

func internalTransferHandler(service *AccountService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			log.Println("Method not allowed")

			return
		}

		type TransferRequest struct {
			FromAccountID string  `json:"from_account_id"`
			ToAccountID   string  `json:"to_account_id"`
			Amount        float64 `json:"amount"`
		}

		var req TransferRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			log.Println("invalid request")

			return
		}

		if req.Amount <= 0 {
			http.Error(w, "invalid transfer amount", http.StatusBadRequest)
			log.Println("invalid transfer amount")

			return
		}

		ctx := r.Context()

		errR := retry(5, 100*time.Millisecond, func() error {
			return service.Transfer(ctx, req.FromAccountID, req.ToAccountID, req.Amount)
		})
		if errR != nil {
			http.Error(w, fmt.Sprintf("Failed to transfer: %v", errR), http.StatusInternalServerError)
			log.Println("Failed to transfer after retry:", errR.Error())

			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		data, err := json.Marshal(req)
		if err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			log.Println("Failed to encode response")

			return
		}

		// simulate some processing delay
		time.Sleep(200 * time.Millisecond)

		w.Write(data)
	}
}

func boom(e error, msg ...string) {
	if e != nil {
		fmt.Println(strings.Join(msg, " "), ":", e)
		panic(e)
	}
}

func assert(condition bool, msg ...string) {
	if !condition {
		fmt.Println(strings.Join(msg, " "), "assert failed")
		panic("assert failed")
	}
}

func retry(attempts int, sleep time.Duration, f func() error) (err error) {
	for i := range attempts {
		if i > 0 {
			log.Println("retrying after error:", err)
			time.Sleep(sleep)
			sleep *= 2
		}

		if err = f(); err == nil {
			return nil
		}

		if !errors.Is(err, ErrStaleObject) {
			return err
		}
	}

	return fmt.Errorf("after %d attempts, last error: %s", attempts, err)
}
