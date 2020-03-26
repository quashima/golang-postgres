package main

import (
	"fmt"

	dbconnection "../infrastructure"
	row "../interface/controller"
)

func main() {
	rows := dbconnection.Connection()
	resultsMap := row.RowToArrayMap(rows)

	fmt.Printf("%s", resultsMap[0].Id)
}
