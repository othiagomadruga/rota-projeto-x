package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDB() {
	var err error
	// Substitua a string de conexão com as informações fornecidas
	db, err = sql.Open("mysql", "avnadmin:AVNS_C_yFDsZiKGn0m6uUeAH@tcp(projeto-x-python-projeto-x-python.i.aivencloud.com:11484)/defaultdb?tls=true&tls_skip_verify=true")
	if err != nil {
		panic(err)
	}

	// Testar a conexão com o banco de dados
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Conectado ao banco de dados MySQL com sucesso!")
}

// Função para buscar um computador pelo device_id
func getComputadorByDeviceID(c *gin.Context) {
	deviceID := c.Param("device_id")
	var id int
	var deviceIDDB string
	var nome string // Adicione outros campos conforme sua tabela

	query := "SELECT id, device_id, nome FROM computadores WHERE device_id = ?" // Ajuste a query e os campos
	err := db.QueryRow(query, deviceID).Scan(&id, &deviceIDDB, &nome)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Computador não encontrado"})
			return
		}
		log.Printf("Erro ao buscar computador: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar computador"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id, "device_id": deviceIDDB, "nome": nome}) // Ajuste os campos retornados
}

// Função para deletar um computador pelo device_id
func deleteComputadorByDeviceID(c *gin.Context) {
	deviceID := c.Param("device_id")

	_, err := db.Exec("DELETE FROM computadores WHERE device_id = ?", deviceID)
	if err != nil {
		log.Printf("Erro ao deletar computador: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar computador"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Computador com device_id '%s' deletado com sucesso", deviceID)})
}

func main() {
	InitDB() // Conectar ao banco

	r := gin.Default()
	r.GET("/computadores/:device_id", getComputadorByDeviceID)
	r.DELETE("/computadores/:device_id", deleteComputadorByDeviceID)
	r.Run(":8080")
}
