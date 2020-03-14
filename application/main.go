package main

import (
	"gopkg.in/ini.v1"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var Config ConfigList

func init() {
	cfg, _ := ini.load("config.ini")
	Config = ConfigList{
		Url:    cfg.Section("db").key("url").MustString(),
		Driver: cfg.Section("db").key("driver").String(),
	}
}

func main() {
	connUrl := Config.Url
	db, err := sql.Open(Config.Driver, connUrl)
	if err != nil {
		log.Fatal(err)
	}

	age := 21
	rows, err := db.Query("SELECT name FROM users WHERE age = $1", age)
}
