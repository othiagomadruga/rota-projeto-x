package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDB() {
	var err error
	db, err = sql.Open("mysql", "USER:PASSWORD@tcp(HOST:PORT)/DATABASE")
	if err != nil {
		panic(err)
	}
}
