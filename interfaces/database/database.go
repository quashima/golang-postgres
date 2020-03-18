package database

import "gopkg.in/ini.v1"

type ConfigList struct {
	Driver  string
	User    string
	DbName  string
	SslMode string
}

var Config ConfigList

func connection() {
	cfg, _ := ini.Load("../config.ini")
	Config = ConfigList{
		Driver:  cfg.Section("db").Key("driver").String(),
		User:    cfg.Section("db").Key("user").String(),
		DbName:  cfg.Section("db").Key("dbname").String(),
		SslMode: cfg.Section("db").Key("sslmode").String(),
	}
}
