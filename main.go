package main

import (
	log "fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/yobdc/jobs/models"
	"net/http"
)

func main() {
	log.Println("hello")
	server := gin.Default()
	server.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	server.Run(":8080")
}
