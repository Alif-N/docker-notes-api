package main

import (
	"log"
	"net/http"
	"os"

	"notes-api/db"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect DB
	db.Connect()
	db.InitSchema()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "API is running 🚀",
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server running on port", port)
	r.Run(":" + port)
}