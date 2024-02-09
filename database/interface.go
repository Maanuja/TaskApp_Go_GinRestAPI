package database

import (
	"context"

	"github.com/EarvinKayonga/tasks/models"
)

type TaskDB interface {
	CreateTask(ctx context.Context, task models.Task) (*models.Task, error)
	GetAllTasks(ctx context.Context) ([]models.Task, error)
	GetTaskByID(ctx context.Context, id string) (*models.Task, error)
	UpdateTask(ctx context.Context, task models.Task) (*models.Task, error)
	DeleteTaskByID(ctx context.Context, id string) (*models.Task, error)
}
