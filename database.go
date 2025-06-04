package main

import (
	"database/sql" // Necessário para *sql.DB e sql.Open
	"fmt"
	"log"    // Necessário para log.Fatal
	"os"     // Necessário para os.Getenv
	_ "github.com/go-sql-driver/mysql" // Driver MySQL
)

// db é declarado aqui para ser a única declaração global no pacote 'main'
var db *sql.DB

// InitDB inicializa a conexão com o banco de dados
func InitDB() {
	// Obtém a URL do banco de dados da variável de ambiente
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		// Fallback para desenvolvimento local se a variável de ambiente não estiver definida
		// IMPORTANTE: Para Render, DATABASE_URL DEVE estar configurada
		dbURL = "root:Thiago123@tcp(127.0.0.1:3306)/crud_go"
		log.Println("WARNING: DATABASE_URL environment variable not set. Using local fallback.")
	}

	var err error
	db, err = sql.Open("mysql", dbURL)
	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
	}

	// Tenta fazer ping no banco de dados para verificar a conexão
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	fmt.Println("Conectado ao banco de dados MySQL com sucesso!")
}
