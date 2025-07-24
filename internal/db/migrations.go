package db

import (
	"database/sql"
	"log"
)

func Migrate(db *sql.DB) error {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS notes (
			id INT AUTO_INCREMENT PRIMARY KEY,
			content TEXT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP 
				ON UPDATE CURRENT_TIMESTAMP
		)`,
		`CREATE INDEX IF NOT EXISTS idx_notes_created_at ON notes(created_at)`,
	}

	for _, query := range queries {
		if _, err := db.Exec(query); err != nil {
			return err
		}
	}

	log.Println("Migrations completed successfully")
	return nil
}
