package handler

import (
	"notes-api/db"
	"notes-api/model"

	"net/http"
	
	"github.com/gin-gonic/gin"
)

func CreateNote(c *gin.Context) {
	var note model.Note
	if err := c.ShouldBindJSON(&note); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `INSERT INTO notes (title, content) VALUES ($1, $2) RETURNING id, created_at, updated_at`
	err := db.DB.QueryRow(query, note.Title, note.Content).Scan(&note.ID, &note.CreatedAt, &note.UpdatedAt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create note"})
		return
	}
	
	c.JSON(http.StatusCreated, note)
}

func GetNotes(c *gin.Context) {
	rows, err := db.DB.Query(`SELECT id, title, content, created_at, updated_at FROM notes ORDER BY created_at DESC`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch notes"})
		return
	}

	defer rows.Close()

	var notes []model.Note
	for rows.Next() {
		var note model.Note
		if err := rows.Scan(&note.ID, &note.Title, &note.Content, &note.CreatedAt, &note.UpdatedAt); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse notes"})
			return
		}
		notes = append(notes, note)
	}

	c.JSON(http.StatusOK, notes)
}