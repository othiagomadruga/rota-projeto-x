package main

import "github.com/gin-gonic/gin"

func main() {
	InitDB() // Conectar ao banco

	r := gin.Default()
	r.GET("/computadores/:device_id", getComputadorByDeviceID)
	r.DELETE("/computadores/:device_id", deleteComputadorByDeviceID)
	r.Run(":8080")
}
