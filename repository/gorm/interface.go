package repository

import (
	"context"
	"reza/todolist-api/model"
)

type PostgresSQLRepository interface {
	CreateTask(ctx context.Context, req model.CreateTaskRequest) error
}
