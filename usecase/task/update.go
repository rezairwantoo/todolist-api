package task

import (
	"context"
	"reza/todolist-api/model"
	"reza/todolist-api/model/constant"

	zlog "github.com/rs/zerolog/log"
)

func (u *Usecase) UpdateTask(ctx context.Context, req model.UpdateTaskRequest) (model.UpdateResponse, error) {
	var (
		task      *model.Task
		err       error
		resp      model.UpdateResponse
		detailReq model.DetailTaskRequest
	)

	detailReq.TaskID = req.ID
	if task, err = u.taskRepo.DetailTask(ctx, detailReq); err != nil {
		zlog.Info().Interface("error", err.Error()).Msg("Failed Get products")
		resp.Message = constant.ErrMsgNotFoundDefault
		return resp, err
	}

	task.Title = req.Title
	task.Description = req.Description
	task.File = req.File
	if err = u.taskRepo.UpdateTask(ctx, *task); err != nil {
		zlog.Info().Interface("error", err.Error()).Msg("Failed update products")
		resp.Message = constant.ErrUpdate
		return resp, err
	}

	resp.Message = constant.SuccessUpdate
	resp.Data.Status = true
	return resp, nil
}
