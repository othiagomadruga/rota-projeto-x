package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDB() {
	dbURL := "root:Thiago123@tcp(127.0.0.1:3306)/crud_go" // Mantenha ou ajuste para a variável de ambiente no Heroku
	var err error
	db, err = sql.Open("mysql", dbURL)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Conectado ao banco de dados MySQL com sucesso!")
}
