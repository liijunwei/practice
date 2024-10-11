package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

// PGPASSWORD= go run .
func main() {
	connStr := fmt.Sprintf("postgres://postgres:%s@localhost/analysis?sslmode=disable", os.Getenv("PGPASSWORD"))
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM pg_timezone_names;")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	result := make([]map[string]any, 0, 100)

	for rows.Next() {
		var name string
		var abbrev string
		var utc_offset string
		var is_dst bool

		err = rows.Scan(&name, &abbrev, &utc_offset, &is_dst)
		if err != nil {
			log.Fatal(err)
		}

		row := make(map[string]any)
		row["name"] = name
		row["abbrev"] = abbrev
		row["utc_offset"] = utc_offset
		row["is_dst"] = is_dst

		result = append(result, row)
	}

	jsonresult, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonresult))
}
