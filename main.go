package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, welcome to the Go API!",
		})
	})

	router.GET("/api/data", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": "This is some sample data from the API.",
		})
	})

	router.Run(":8080")
}
