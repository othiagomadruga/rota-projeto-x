package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getComputadorByDeviceID(c *gin.Context) {
	deviceID := c.Param("device_id")

	var computador Computador
	err := db.QueryRow("SELECT device_id, nome, preco FROM computadores WHERE device_id = ?", deviceID).Scan(&computador.DeviceID, &computador.Nome, &computador.Preco)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Computador não encontrado"})
		return
	}

	c.JSON(http.StatusOK, computador)
}

func deleteComputadorByDeviceID(c *gin.Context) {
	deviceID := c.Param("device_id")

	_, err := db.Exec("DELETE FROM computadores WHERE device_id = ?", deviceID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar o computador"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Computador deletado com sucesso"})
}
