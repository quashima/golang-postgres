package main

import (
	"database/sql"
	"fmt"

	dbconnection "./infrastructure"

	_ "github.com/lib/pq"
)

type ConfigList struct {
	Driver  string
	User    string
	DbName  string
	SslMode string
}

type User struct {
	Id string
}

var Config ConfigList

func arrayMap(rows *sql.Rows) map[int]*User {
	var usersMap = map[int]*User{}
	var count = 0
	for rows.Next() {
		var user User
		rows.Scan(&user.Id)

		usersMap[count] = &User{Id: user.Id}
		count++
	}
	return usersMap
}

func main() {
	rows := dbconnection.Connection()
	resultsMap := arrayMap(rows)

	fmt.Printf("%s", resultsMap[0].Id)
}
