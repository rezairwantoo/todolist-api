package subtask

import (
	"context"
	"reza/todolist-api/model"
)

type SubTaskUsecase interface {
	Create(ctx context.Context, req model.CreateSubTaskRequest) (model.CreateResponse, error)
	Detail(ctx context.Context, req model.DetailSubTaskRequest) (model.DetailResponse, error)
	List(ctx context.Context, req model.ListSubTaskRequest) (*model.ListResponse, error)
	Update(ctx context.Context, req model.UpdateSubTaskRequest) (model.UpdateResponse, error)
	Delete(ctx context.Context, req model.DeleteSubTaskRequest) (model.DeleteResponse, error)
}
