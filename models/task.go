package models

import (
	"time"

	"github.com/google/uuid"
)

// Task structure
type Task struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Status    Status    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

type Status string

var (
	StatusDone    Status = "done"
	StatusPending Status = "pending"
)

func CreateTask(
	id, title string,
	status Status,
	createdAt time.Time) *Task {
	if id == "" {
		id = uuid.New().String()
	}

	if status != StatusDone && status != StatusPending {
		status = StatusPending
	}

	empty := time.Time{}
	if createdAt == empty {
		createdAt = time.Now()
	}

	return &Task{
		ID:        id,
		Title:     title,
		Status:    status,
		CreatedAt: createdAt,
	}
}
