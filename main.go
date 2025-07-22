package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/meshShinrai/quick-notes/internal/db"
	"github.com/meshShinrai/quick-notes/internal/handlers"
)

func main() {
	// Load environment variables from .env if present
	//_ = os.Setenv("AWS_REGION", "your-region")         // Optional: put in .env
	//_ = os.Setenv("AWS_SECRET_NAME", "your-secret-id") // Optional: put in .env

	database, err := db.Connect()
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}
	defer database.Close()

	db.Migrate(database)

	router := gin.Default()

	noteHandler := handlers.NewNoteHandler(database)

	router.GET("/notes", noteHandler.GetAllNotes)
	router.POST("/notes", noteHandler.CreateNote)
	router.PUT("/notes/:id", noteHandler.UpdateNote)
	router.DELETE("/notes/:id", noteHandler.DeleteNote)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server is running on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
