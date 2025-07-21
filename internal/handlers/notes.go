package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/meshShinrai/quick-notes/internal/models"
)

func GetNotes(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := db.Query("SELECT id, content FROM notes ORDER BY id DESC")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch notes"})
			return
		}
		defer rows.Close()

		var notes []models.Note
		for rows.Next() {
			var note models.Note
			if err := rows.Scan(&note.ID, &note.Content); err != nil {
				continue
			}
			notes = append(notes, note)
		}

		c.JSON(http.StatusOK, notes)
	}
}

func CreateNote(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var note models.Note
		if err := c.BindJSON(&note); err != nil || note.Content == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid note content"})
			return
		}

		res, err := db.Exec("INSERT INTO notes (content) VALUES (?)", note.Content)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert note"})
			return
		}

		id, _ := res.LastInsertId()
		note.ID = int(id)
		c.JSON(http.StatusCreated, note)
	}
}

func DeleteNote(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		_, err := db.Exec("DELETE FROM notes WHERE id = ?", id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete note"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Note deleted"})
	}
}
