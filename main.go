package main

import (
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/meshShinrai/quick-notes/internal/db"
	"github.com/meshShinrai/quick-notes/internal/handlers"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Initialize database
	database, err := db.Connect()
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}
	defer database.Close()

	// Run migrations
	if err := db.Migrate(database); err != nil {
		log.Fatalf("Database migration failed: %v", err)
	}

	// Configure Gin
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()

	// Enhanced CORS configuration
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:5173",
			"https://*.amplifyapp.com",
			"http://quick-notes-env.eba-3qq8dwtj.eu-north-1.elasticbeanstalk.com",
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "healthy"})
	})

	// API routes
	api := router.Group("/api")
	{
		noteHandler := handlers.NewNoteHandler(database)
		api.GET("/notes", noteHandler.GetAllNotes)
		api.POST("/notes", noteHandler.CreateNote)
		api.PUT("/notes/:id", noteHandler.UpdateNote)
		api.DELETE("/notes/:id", noteHandler.DeleteNote)
	}

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	log.Printf("Server starting on port %s", port)
	log.Fatal(router.Run(":" + port))
}
