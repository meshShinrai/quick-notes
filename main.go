package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/meshShinrai/quick-notes/internal/db"
)

var database *sql.DB

func init() {
	// Load environment variables from .env if present
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found or failed to load. Using environment variables.")
	}
}

func main() {
	secretName := os.Getenv("SECRET_NAME")
	if secretName == "" {
		log.Fatal("SECRET_NAME not set in environment")
	}

	// Connect to the DB
	var err error
	database, err = db.Connect(secretName)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.Close()

	// Create Gin router
	router := gin.Default()

	// Health check route
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	// Placeholder for future notes routes
	// e.g., router.GET("/notes", handlers.GetNotes)

	log.Println("Server running on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
