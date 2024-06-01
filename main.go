package main

import (
	"gin-mongo-api/configs"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": "Hello from Gin-gonic & mongoDB",
		})
	})
	configs.ConnectDB()

	router.Run("localhost:8080")
}
