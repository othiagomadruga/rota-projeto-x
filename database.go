package main

import (
	"database/sql"
	"fmt"
	"os" // Importe o pacote "os"
	"log" // Mantenha a importação de "log" por precaução, mesmo que não esteja explicitamente usado agora
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDB() {
	dbURL := os.Getenv("DATABASE_URL") // Usando a variável de ambiente para o Heroku (ou Render)
	var err error
	db, err = sql.Open("mysql", dbURL)
	if err != nil {
		log.Fatal(err) // Use log.Fatal para encerrar em caso de erro crítico na inicialização
		return
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err) // Use log.Fatal para encerrar em caso de erro crítico na inicialização
		return
	}
	fmt.Println("Conectado ao banco de dados MySQL com sucesso!")
}
