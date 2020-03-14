package main

import "fmt"
import (
	"database/sql"

	_ "github.com/lib/pq"
)


func main() {
	connStr := ""
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	age := 21
	rows, err := db.Query("SELECT name FROM users WHERE age = $1", age)
	
}
