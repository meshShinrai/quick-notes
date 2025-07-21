package main

import (
	"github.com/gin-gonic/gin"
	"log"

	"github.com/meshShinrai/quick-notes/db"
	"github.com/meshShinrai/quick-notes/handlers"
)

func main() {
	database, err := db.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	db.Migrate(database)

	router := gin.Default()
	noteHandler := handlers.NewNoteHandler(database)

	router.GET("/notes", noteHandler.GetAllNotes)
	router.POST("/notes", noteHandler.CreateNote)
	router.PUT("/notes/:id", noteHandler.UpdateNote)
	router.DELETE("/notes/:id", noteHandler.DeleteNote)

	router.Run(":8080")
}
