package models

import (
	"time"
)

type Todo struct {
	TodoID    int       `json:"todo_id"`
	Task      string    `json:"task"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"created_at"`
}
