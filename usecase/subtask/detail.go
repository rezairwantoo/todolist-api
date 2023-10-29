package subtask

import (
	"context"
	"reza/todolist-api/model"
	"reza/todolist-api/model/constant"
)

func (u *Usecase) Detail(ctx context.Context, req model.DetailSubTaskRequest) (model.DetailResponse, error) {
	var (
		subtask *model.Task
		err     error
		resp    model.DetailResponse
	)

	if subtask, err = u.CheckDetailSubtask(ctx, req); err != nil {
		return resp, err
	}

	return model.DetailResponse{
		Message: constant.SuccessDetail,
		Data:    subtask,
	}, nil
}
