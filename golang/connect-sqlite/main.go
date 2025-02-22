package main

import (
	"database/sql"
	"fmt"
	"golang-practices/connect-sqlite/sqlcdb"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"text/template"

	_ "github.com/mattn/go-sqlite3" // SQLite driver
)

var once sync.Once

// simple todo app with sqlite3+sqlc
//
// features:
// 1. create a todo
// 2. delete a todo
// 3. list all todos
//
// go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
// guide: https://docs.sqlc.dev/en/latest/tutorials/getting-started-sqlite.html
//
// rm /tmp/todo-app.db && sqlc generate && go run main.go
func main() {
	db := initDB()
	defer db.Close()

	queries := sqlcdb.New(db)

	http.HandleFunc("GET /", indexHandler(queries))
	http.HandleFunc("POST /create", createHandler(queries))
	http.HandleFunc("GET /delete", deleteHandler(queries)) // ideally should be DELETE

	fmt.Println("Server is running at http://localhost:8080")
	boom(http.ListenAndServe(":8080", nil))
}

type todo struct {
	ID    int64
	Title string
}

func initDB() *sql.DB {
	db, err := sql.Open("sqlite3", "/tmp/todo-app.db")
	boom(err, "failed to open database")

	schema, err := os.ReadFile(filepath.Join(baseDir(), "schema.sql"))
	boom(err, "failed to read schema.sql")

	_, err = db.Exec(string(schema))
	boom(err, "failed to execute schema")

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
		entities, err := db.ListTODOs(r.Context())
		boom(err, "failed to list todos")

		todos := make([]todo, 0, len(entities))

		for _, entity := range entities {
			todos = append(todos, todo{
				ID:    entity.ID,
				Title: entity.Title,
			})
		}

		tmpl := template.Must(template.New("index").Parse(`
		<!DOCTYPE html>
		<html>
			<head>
				<title>Todo List</title>
			</head>
			<body>
				<h1>Todo List</h1>
				<form action="/create" method="POST">
					<input type="text" name="title" placeholder="New Todo" required>
					<button type="submit">Add</button>
				</form>
				<ul>
					{{range .}}
					<li>{{.Title}} <a href="/delete?id={{.ID}}">Delete</a></li>
					{{end}}
				</ul>
			</body>
		</html>
		`))

		tmpl.Execute(w, todos)
	}
}

func createHandler(db *sqlcdb.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		assert(r.Method == "POST", "invalid method")

		title := r.FormValue("title")
		assert(title != "", "title required")

		err := db.CreateTODO(r.Context(), title)
		boom(err, "failed to create todo")

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func deleteHandler(db *sqlcdb.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.URL.Query().Get("id")
		assert(idStr != "", "id required")

		id, err := strconv.Atoi(idStr)
		boom(err, "failed to convert id")

		err = db.DeleteTODO(r.Context(), int64(id))
		boom(err, "failed to delete todo")

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func boom(e error, msg ...string) {
	if e != nil {
		fmt.Println(strings.Join(msg, " "), e)
		panic(e)
	}
}

func assert(condition bool, msg ...string) {
	if !condition {
		fmt.Println(strings.Join(msg, " "), "assert failed")
		panic("assert failed")
	}
}
