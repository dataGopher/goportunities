package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Initialize() {
	// Inicia uma rota utilizando a configuracao default do gin
	router := gin.Default()
	// Define uma rota GET para o path /ping
	router.GET("/ping", func(c *gin.Context) {
		// Retorna uma resposta JSON com o status 200
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	// Roda o servidor na porta 8080
	router.Run("8080")
}
