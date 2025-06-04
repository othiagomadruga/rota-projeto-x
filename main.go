package main

import (
	"log"    // <-- Importe o pacote "log" se ainda não estiver
	"net/http" // <-- Mantenha se estiver usando constantes http.StatusOK, etc. em outras funcoes no main.go
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// db é declarado aqui para ser acessível pelas funções de rota
var db *sql.DB // Certifique-se de que esta linha está presente e única em todo o projeto.
              // Ela deve estar em 'database.go' ou 'main.go', mas não em ambos.
              // Se 'db' for declarado em 'database.go', remova esta linha.
              // Se 'db' for declarado aqui, remova de 'database.go'.
              // No contexto das nossas correções, 'db' está declarado e inicializado em 'database.go'
              // e é acessível no mesmo pacote 'main'.

// Importante: As funções InitDB, getComputadorByDeviceID, deleteComputadorByDeviceID
// DEVEM estar definidas em 'database.go' e 'routes.go' respectivamente,
// e NÃO devem ser duplicadas aqui em 'main.go'.
// A chamada a elas é feita aqui.

func main() {
	InitDB() // Conectar ao banco

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Use 8080 como padrão localmente
	}

	log.Printf("Attempting to start server on port: %s", port) // Novo log de depuração
	
	r := gin.Default()
	
	// Suas rotas
	r.GET("/computadores/:device_id", getComputadorByDeviceID)
	r.DELETE("/computadores/:device_id", deleteComputadorByDeviceID)
	
	// Captura o erro do r.Run
	err := r.Run(":" + port)
	if err != nil {
		log.Fatalf("Server failed to start on port %s: %v", port, err) // Log fatal se o servidor não iniciar
	}
	log.Printf("Server successfully started on port: %s", port) // Novo log de sucesso
}

// A struct Computador deve estar definida em um único lugar no seu projeto,
// por exemplo, em 'main.go' ou em um arquivo 'models.go' separado.
type Computador struct {
	DeviceID string  `json:"device_id"`
	Nome     string  `json:"nome"`
	Preco    float64 `json:"preco"`
}
