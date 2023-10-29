package subtask

import (
	"context"
	"reza/todolist-api/model"
)

type SubTaskUsecase interface {
	Create(ctx context.Context, req model.CreateSubTaskRequest) (model.CreateResponse, error)
	// Detail(ctx context.Context, req model.DetailTaskRequest) (model.DetailResponse, error)
	// List(ctx context.Context, req model.ListRequest) (*model.ListResponse, error)
	// UpdateTask(ctx context.Context, req model.UpdateTaskRequest) (model.UpdateResponse, error)
	// DeleteTask(ctx context.Context, req model.DeleteTaskRequest) (model.DeleteResponse, error)
}
