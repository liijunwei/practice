package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"
	"text/template"

	_ "github.com/mattn/go-sqlite3" // SQLite driver
)

// simple todo app with sqlite3
// 1. create a todo
// 2. delete a todo
// 3. list all todos
func main() {
	db := initDB()
	defer db.Close()

	http.HandleFunc("/", indexHandler(db))
	http.HandleFunc("/create", createHandler(db))
	http.HandleFunc("/delete", deleteHandler(db))

	fmt.Println("Server is running at http://localhost:8080")
	boom(http.ListenAndServe(":8080", nil))
}

type Todo struct {
	ID    int
	Title string
}

func initDB() *sql.DB {
	db, err := sql.Open("sqlite3", "/tmp/todo-app.db")
	boom(err, "failed to open database")

	sql := `
	CREATE TABLE IF NOT EXISTS todos (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		title TEXT
	);`

	_, err = db.Exec(sql)
	boom(err, "failed to create table")

	return db
}

func indexHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT id, title FROM todos")
		boom(err, "failed to select todos")

		defer rows.Close()

		todos := []Todo{}

		for rows.Next() {
			var todo Todo

			err := rows.Scan(&todo.ID, &todo.Title)
			boom(err, "failed to scan row")

			todos = append(todos, todo)
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

func createHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			title := r.FormValue("title")
			assert(title != "", "title required")

			_, err := db.Exec("INSERT INTO todos(title) VALUES(?)", title)
			boom(err, "failed to create todo")

			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}
}

func deleteHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		assert(id != "", "id required")

		_, err := db.Exec("DELETE FROM todos WHERE id = ?", id)
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
