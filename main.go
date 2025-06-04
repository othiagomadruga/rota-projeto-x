package main

import (
	"log" // Necessário para log.Printf e log.Fatalf
	// "net/http", // <--- REMOVA OU COMENTE ESTA LINHA EM main.go
	"os"   // Necessário para os.Getenv

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql" // Driver MySQL
)

// main é a função de entrada da aplicação
func main() {
	// Inicializa a conexão com o banco de dados
	InitDB()

	// Obtém a porta da variável de ambiente fornecida pelo Render
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Padrão para desenvolvimento local
	}

	log.Printf("Attempting to start server on port: %s", port)

	// Inicializa o roteador Gin
	r := gin.Default()

	// Define as rotas da API
	r.GET("/computadores/:device_id", getComputadorByDeviceID)
	r.DELETE("/computadores/:device_id", deleteComputadorByDeviceID)

	// Inicia o servidor Gin
	err := r.Run(":" + port)
	if err != nil {
		log.Fatalf("Server failed to start on port %s: %v", port, err)
	}
	log.Printf("Server successfully started on port: %s", port)
}

// Computador representa a estrutura de dados de um computador
type Computador struct {
	DeviceID string  `json:"device_id"`
	Nome     string  `json:"nome"`
	Preco    float64 `json:"preco"`
}
