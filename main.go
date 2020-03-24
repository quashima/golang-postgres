package main

import (
	"fmt"

	dbconnection "./infrastructure"
	row "./interface/controller"

	_ "github.com/lib/pq"
)

func main() {
	rows := dbconnection.Connection()
	resultsMap := row.RowToArrayMap(rows)

	fmt.Printf("%s", resultsMap[0].Id)
}
