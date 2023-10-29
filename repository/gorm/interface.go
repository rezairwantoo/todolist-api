package repository

import (
	"context"
	"reza/todolist-api/model"
)

type TaskRepository interface {
	CreateTask(ctx context.Context, req model.CreateTaskRequest) error
	DetailTask(ctx context.Context, req model.DetailTaskRequest) (*model.Task, error)
	ListTask(pagination model.Pagination, search string) (*model.Pagination, error)
	UpdateTask(ctx context.Context, req model.Task) error
	DeleteTask(ctx context.Context, taskID int64) error
	DeleteSubTask(ctx context.Context, taskID int64) error
}
