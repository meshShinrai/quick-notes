package db

import (
	"database/sql"
	"log"
)

func Migrate(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS notes (
		id INT AUTO_INCREMENT PRIMARY KEY,
		content TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
	);`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	log.Println("Migration completed successfully.")
}
