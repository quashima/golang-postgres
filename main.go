package main

import (
	"database/sql"
	"fmt"

	usertable "./domain"
	dbconnection "./infrastructure"

	_ "github.com/lib/pq"
)

type ConfigList struct {
	Driver  string
	User    string
	DbName  string
	SslMode string
}

var Config ConfigList

func arrayMap(rows *sql.Rows) map[int]*usertable.Column {
	var usersMap = map[int]*usertable.Column{}
	var count = 0
	for rows.Next() {
		var user usertable.Column
		rows.Scan(&user.Id)

		usersMap[count] = &usertable.Column{Id: user.Id}
		count++
	}
	return usersMap
}

func main() {
	rows := dbconnection.Connection()
	resultsMap := arrayMap(rows)

	fmt.Printf("%s", resultsMap[0].Id)
}
