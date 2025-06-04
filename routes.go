package main

import (
	"net/http" // Necessário para http.StatusNotFound, http.StatusOK, etc.

	"github.com/gin-gonic/gin"
)

// getComputadorByDeviceID lida com a requisição GET para um computador específico
func getComputadorByDeviceID(c *gin.Context) {
	deviceID := c.Param("device_id")

	var computador Computador
	// db é acessível aqui porque é uma variável global no pacote 'main' (definida em database.go)
	err := db.QueryRow("SELECT device_id, nome, preco FROM computadores WHERE device_id = ?", deviceID).Scan(&computador.DeviceID, &computador.Nome, &computador.Preco)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Computador não encontrado"})
		return
	}

	c.JSON(http.StatusOK, computador)
}

// deleteComputadorByDeviceID lida com a requisição DELETE para um computador específico
func deleteComputadorByDeviceID(c *gin.Context) {
	deviceID := c.Param("device_id")

	// db é acessível aqui porque é uma variável global no pacote 'main' (definida em database.go)
	_, err := db.Exec("DELETE FROM computadores WHERE device_id = ?", deviceID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar o computador"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Computador deletado com sucesso"})
}
