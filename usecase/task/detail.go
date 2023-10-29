package task

import (
	"context"
	"reza/todolist-api/model"
	"reza/todolist-api/model/constant"

	zlog "github.com/rs/zerolog/log"
)

func (u *Usecase) Detail(ctx context.Context, req model.DetailTaskRequest) (model.DetailResponse, error) {
	var (
		task *model.Task
		err  error
		resp model.DetailResponse
	)

	if task, err = u.taskRepo.DetailTask(ctx, req); err != nil {
		zlog.Info().Interface("error", err.Error()).Msg("Failed Get products")
		resp.Message = constant.ErrMsgNotFoundDefault
		return resp, err
	}

	return model.DetailResponse{
		Message: constant.SuccessDetail,
		Data:    task,
	}, nil
}
