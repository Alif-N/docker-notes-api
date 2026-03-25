package handler

import (
	"notes-api/service"
	"notes-api/model"

	"net/http"
	"fmt"
	
	"github.com/gin-gonic/gin"
)

func CreateNote(c *gin.Context) {
	var note model.Note
	if err := c.ShouldBindJSON(&note); err != nil {
		Error(c, http.StatusBadRequest, "Invalid request: "+err.Error())
		return
	}

	if err := service.CreateNote(&note); err != nil {
		Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	
	Success(c, note)
}

func GetNotes(c *gin.Context) {
	// default pagination values
	page := 1
	limit := 10
	search := c.Query("search")

	// parse pagination query params
	if p := c.Query("page"); p != "" {
		fmt.Sscanf(p, "%d", &page)
	}
	if l := c.Query("limit"); l != "" {
		fmt.Sscanf(l, "%d", &limit)
	}

	notes, total, err := service.GetNotes(page, limit, search)
	if err != nil {
		Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	
	Success(c, gin.H{
		"data": notes,
		"total": total,
		"page": page,
		"limit": limit,
	})
}

func GetNoteByID(c *gin.Context) {
	id := c.Param("id")

	var note model.Note
	notePtr, err := service.GetNoteByID(id)
	if err != nil {
		Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	if notePtr == nil {
		Error(c, http.StatusNotFound, "Note not found")
		return
	}
	note = *notePtr

	Success(c, note)
}

func UpdateNote(c *gin.Context) {
	id := c.Param("id")

	var note model.Note
	if err := c.ShouldBindJSON(&note); err != nil {
		Error(c, http.StatusBadRequest, "Invalid request: "+err.Error())
		return
	}

	if err := service.UpdateNote(id, &note); err != nil {
		Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	Success(c, note)
}

func DeleteNote(c *gin.Context) {
	id := c.Param("id")

	rowsAffected, err := service.DeleteNote(id)
	if err != nil {
		if rowsAffected == 0 {
			Error(c, http.StatusNotFound, "Note not found")
			return
		}
		Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	Success(c, gin.H{
		"message":      "Note deleted successfully",
		"rowsAffected": rowsAffected,
	})
}