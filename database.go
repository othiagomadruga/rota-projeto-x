package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDB() {
	dbURL := os.Getenv("DATABASE_URL") // Usando a variável de ambiente
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
