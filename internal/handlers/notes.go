package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/meshShinrai/quick-notes/internal/models"
)

type NoteHandler struct {
	db *sql.DB
}

func NewNoteHandler(db *sql.DB) *NoteHandler {
	return &NoteHandler{db: db}
}

func (h *NoteHandler) GetAllNotes(c *gin.Context) {
	const query = `SELECT id, content, created_at, updated_at FROM notes ORDER BY created_at DESC`
	rows, err := h.db.Query(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
		return
	}
	defer rows.Close()

	var notes []models.Note
	for rows.Next() {
		var note models.Note
		if err := rows.Scan(&note.ID, &note.Content, &note.CreatedAt, &note.UpdatedAt); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "data parsing error"})
			return
		}
		notes = append(notes, note)
	}

	c.JSON(http.StatusOK, notes)
}

func (h *NoteHandler) CreateNote(c *gin.Context) {
	var input struct {
		Content string `json:"content" binding:"required,min=1,max=500"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	result, err := h.db.Exec("INSERT INTO notes (content) VALUES (?)", input.Content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create note"})
		return
	}

	id, _ := result.LastInsertId()
	c.JSON(http.StatusCreated, gin.H{
		"id":      id,
		"content": input.Content,
	})
}

func (h *NoteHandler) UpdateNote(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid note ID"})
		return
	}

	var input struct {
		Content string `json:"content" binding:"required,min=1,max=500"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	res, err := h.db.Exec("UPDATE notes SET content = ? WHERE id = ?", input.Content, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
		return
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "note not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "note updated"})
}

func (h *NoteHandler) DeleteNote(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid note ID"})
		return
	}

	res, err := h.db.Exec("DELETE FROM notes WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
		return
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "note not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "note deleted"})
}
