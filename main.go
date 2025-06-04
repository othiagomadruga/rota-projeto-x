package main

import (
	"log" // Necessário para log.Printf e log.Fatalf
	"net/http" // Necessário para http.StatusOK, http.StatusNotFound, etc.
	"os"   // Necessário para os.Getenv

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql" // Driver MySQL (mantido aqui por precaução, mas pode ser movido para database.go se não houver outros usos no main.go)
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
	// As funções getComputadorByDeviceID e deleteComputadorByDeviceID
	// são definidas em 'routes.go' e acessam a variável 'db'
	// que é global no pacote 'main' (definida em 'database.go').
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
