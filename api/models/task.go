package models

import (
	"time"

	"github.com/google/uuid"
)

type TaskStatus string

const (
	Pending    TaskStatus = "PENDING"
	InProgress TaskStatus = "IN_PROGRESS"
	Completed  TaskStatus = "COMPLETED"
)

type Task struct {
	ID          uuid.UUID  `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	UserID      uuid.UUID  `json:"user_id"`
	CreatedAt   time.Time  `json:"created_at"  binding:"omitempty"`
	UpdatedAt   time.Time  `json:"updated_at" binding:"omitempty"`
}
