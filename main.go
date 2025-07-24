//go:build !glibc

package main

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/meshShinrai/quick-notes/internal/db"
	"github.com/meshShinrai/quick-notes/internal/handlers"
)

func main() {
	// Load environment variables from .env if present
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found. Using environment variables")
	}

	database, err := db.Connect()
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}
	defer database.Close()

	db.Migrate(database)

	router := gin.Default()

	// Allow CORS for local frontend
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		AllowCredentials: true,
	}))

	noteHandler := handlers.NewNoteHandler(database)

	router.GET("/notes", noteHandler.GetAllNotes)
	router.POST("/notes", noteHandler.CreateNote)
	router.PUT("/notes/:id", noteHandler.UpdateNote)
	router.DELETE("/notes/:id", noteHandler.DeleteNote)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	log.Printf("Server is running on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
