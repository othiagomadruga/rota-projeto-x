package main

import (
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	InitDB() // Conectar ao banco

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Use 8080 como padrão localmente
	}

	r := gin.Default()
	r.GET("/computadores/:device_id", getComputadorByDeviceID)
	r.DELETE("/computadores/:device_id", deleteComputadorByDeviceID)
	r.Run(":" + port)
}

type Computador struct {
	DeviceID string  `json:"device_id"`
	Nome     string  `json:"nome"`
	Preco    float64 `json:"preco"`
}
