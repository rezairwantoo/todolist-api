package repository

import (
	"context"
	"reza/todolist-api/model"
)

type TaskRepository interface {
	CreateTask(ctx context.Context, req model.CreateTaskRequest) error
}
