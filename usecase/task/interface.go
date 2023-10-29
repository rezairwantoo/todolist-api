package task

import (
	"context"
	"reza/todolist-api/model"
)

type TaskUsecase interface {
	Create(ctx context.Context, req model.CreateTaskRequest) (model.CreateResponse, error)
}
