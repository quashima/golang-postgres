package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type ConfigList struct {
	Driver  string
	User    string
	DbName  string
	SslMode string
}

var Config ConfigList

func init() {
	cfg, _ := ini.load("config.ini")
	Config = ConfigList{
		Driver:  cfg.Section("db").key("driver").String(),
		User:    cfg.Section("db").key("user").String(),
		DbName:  cfg.Section("db").key("dbname").String(),
		SslMode: cfg.Section("db").key("sslmode").String(),
	}
}

func main() {
	connStr := "user=" + Config.User + "dbname=" + Config.DbName + "sslmode=" + Config.SslMode
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("select * from user")
	fmt.Printf("%v", rows)
}
