package task

import (
	"context"
	"reza/todolist-api/model"
)

type TaskUsecase interface {
	Create(ctx context.Context, req model.CreateTaskRequest) (model.CreateResponse, error)
	Detail(ctx context.Context, req model.DetailTaskRequest) (model.DetailResponse, error)
	List(ctx context.Context, req model.ListRequest) (*model.ListResponse, error)
	UpdateTask(ctx context.Context, req model.UpdateTaskRequest) (model.UpdateResponse, error)
}
