package models

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	Id          uuid.UUID `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	IsCompleted bool      `json:"is_completed" db:"is_completed"`
}

type CreateTask struct {
	Title       string `json:"title" db:"title"`
	Description string `json:"description"`
}
