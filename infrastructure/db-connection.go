package dbconnection

import (
	"database/sql"
	"log"

	"gopkg.in/ini.v1"
)

type ConfigList struct {
	Driver  string
	User    string
	DbName  string
	SslMode string
}

var Config ConfigList

func init() {
	cfg, _ := ini.Load("./config.ini")
	Config = ConfigList{
		Driver:  cfg.Section("db").Key("driver").String(),
		User:    cfg.Section("db").Key("user").String(),
		DbName:  cfg.Section("db").Key("dbname").String(),
		SslMode: cfg.Section("db").Key("sslmode").String(),
	}
}

func Connection() *sql.Rows {
	connStr := "user=" + Config.User + " dbname=" + Config.DbName + " sslmode=" + Config.SslMode
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("select * from user")
	return rows
}
