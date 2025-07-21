package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/meshShinrai/quick-notes/internal/models"
)

type NoteHandler struct {
	DB *sql.DB
}

func NewNoteHandler(db *sql.DB) *NoteHandler {
	return &NoteHandler{DB: db}
}

// GET /notes
func (h *NoteHandler) GetAllNotes(c *gin.Context) {
	rows, err := h.DB.Query("SELECT id, content, created_at, updated_at FROM notes")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch notes"})
		return
	}
	defer rows.Close()

	var notes []models.Note
	for rows.Next() {
		var note models.Note
		if err := rows.Scan(&note.ID, &note.Content, &note.CreatedAt, &note.UpdatedAt); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan note"})
			return
		}
		notes = append(notes, note)
	}

	c.JSON(http.StatusOK, notes)
}

// POST /notes
func (h *NoteHandler) CreateNote(c *gin.Context) {
	var input struct {
		Content string `json:"content" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Content is required"})
		return
	}

	query := "INSERT INTO notes (content) VALUES (?)"
	result, err := h.DB.Exec(query, input.Content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert note"})
		return
	}

	id, _ := result.LastInsertId()
	c.JSON(http.StatusCreated, gin.H{"message": "Note created", "id": id})
}

// PUT /notes/:id
func (h *NoteHandler) UpdateNote(c *gin.Context) {
	id := c.Param("id")

	var input struct {
		Content string `json:"content" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Content is required"})
		return
	}

	_, err := h.DB.Exec("UPDATE notes SET content = ? WHERE id = ?", input.Content, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update note"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Note updated"})
}

// DELETE /notes/:id
func (h *NoteHandler) DeleteNote(c *gin.Context) {
	id := c.Param("id")

	_, err := h.DB.Exec("DELETE FROM notes WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete note"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Note deleted"})
}
