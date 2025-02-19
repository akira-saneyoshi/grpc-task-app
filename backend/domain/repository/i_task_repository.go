package repository

import (
	"context"

	"github.com/akira-saneyoshi/task-app/domain/object/entity"
)

// Persist TaskEntity
type ITaskRepository interface {
	FindTaskByID(ctx context.Context, id string) (*entity.Task, error)
	FindTasksByUserID(ctx context.Context, userID string) ([]*entity.Task, error)
	CreateTask(ctx context.Context, arg *entity.Task) (string, error)
	UpdateTask(ctx context.Context, arg *entity.Task) error
	DeleteTask(ctx context.Context, id string) error
}
