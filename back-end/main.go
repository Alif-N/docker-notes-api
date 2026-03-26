package main

import (
	"log"
	"net/http"
	"os"

	"notes-api/db"
	"notes-api/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect DB
	db.Connect()

	r := gin.Default()

	r.POST("/notes", handler.CreateNote)
	r.GET("/notes", handler.GetNotes)
	r.PUT("/notes/:id", handler.UpdateNote)
	r.DELETE("/notes/:id", handler.DeleteNote)
	r.GET("/notes/:id", handler.GetNoteByID)

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