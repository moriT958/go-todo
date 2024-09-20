package models

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type Todo struct {
	TodoID      int       `json:"todo_id"`
	Task        string    `json:"task"`
	Done        bool      `json:"done"`
	CreatedAt   time.Time `json:"created_at"`
	CompletedAt time.Time `json:"completed_at"`
}

func InitDB(db *sql.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS todos (
			id SERIAL PRIMARY KEY,
			task VARCHAR(20) UNIQUE NOT NULL,
			done BOOLEAN NOT NULL,
			created_at TIMESTAMP NOT NULL,
			completed_at TIMESTAMP NOT NULL 
		) 
	`
	if _, err := db.Exec(query); err != nil {
		return err
	}

	log.Println("Successfly created table!")
	return nil
}
