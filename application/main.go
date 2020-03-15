package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"gopkg.in/ini.v1"
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

func init() {
	cfg, _ := ini.Load("../config.ini")
	Config = ConfigList{
		Driver:  cfg.Section("db").Key("driver").String(),
		User:    cfg.Section("db").Key("user").String(),
		DbName:  cfg.Section("db").Key("dbname").String(),
		SslMode: cfg.Section("db").Key("sslmode").String(),
	}
}

func main() {
	connStr := "user=" + Config.User + " dbname=" + Config.DbName + " sslmode=" + Config.SslMode
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("select * from user")
	var users []User
	for rows.Next() {
		var user User
		rows.Scan(&user.Id)
		users = append(users, user)
	}
	fmt.Printf("%v", users)
}
