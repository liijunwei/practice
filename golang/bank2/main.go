package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"golang-practices/bank2/sqlcdb"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

var once sync.Once

// simple bank app with pg+sqlc
// rm /tmp/bank2-app.db; sqlc generate && go run .
func main() {
	dsn := "postgres://postgres:123@192.168.31.50:5432/demo?sslmode=disable&application_name=bank2"

	db := initDB(dsn)
	defer db.Close()

	queries := sqlcdb.New(db)
	service := NewAccountService(db)

	http.HandleFunc("GET /", indexHandler(queries))
	http.HandleFunc("POST /register", registerUserHandler(service))
	http.HandleFunc("POST /transfer", internalTransferHandler(service))

	go func() {
		log.Println("start pprof server", "http://localhost:6060/debug/pprof")
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	log.Println("start bank server", "http://localhost:8080")
	boom(http.ListenAndServe(":8080", nil))
}

var _ sqlcdb.DBTX = &sql.DB{}

func initDB(dsn string) *sql.DB {
	db, err := sql.Open("postgres", dsn)
	boom(err, "failed to open db connection")

	db.SetMaxOpenConns(10)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	boom(db.PingContext(ctx))

	_, err = db.Exec("drop table if exists accounts, account_events, transactions, users;")
	boom(err, "failed to drop tables")

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
			FromAccountID uuid.UUID `json:"from_account_id"`
			ToAccountID   uuid.UUID `json:"to_account_id"`
			Amount        float64   `json:"amount"`
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

		errR := service.Transfer(ctx, req.FromAccountID, req.ToAccountID, req.Amount)
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
