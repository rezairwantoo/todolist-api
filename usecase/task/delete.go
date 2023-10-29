package task

import (
	"context"
	"reza/todolist-api/model"
	"reza/todolist-api/model/constant"

	zlog "github.com/rs/zerolog/log"
)

func (u *Usecase) DeleteTask(ctx context.Context, req model.DeleteTaskRequest) (model.DeleteResponse, error) {
	var (
		err       error
		resp      model.DeleteResponse
		detailReq model.DetailTaskRequest
	)

	detailReq.TaskID = req.TaskID
	if _, err = u.taskRepo.DetailTask(ctx, detailReq); err != nil {
		zlog.Info().Interface("error", err.Error()).Msg("Failed Get task")
		resp.Message = constant.ErrMsgNotFoundDefault
		return resp, err
	}

	if err = u.taskRepo.DeleteSubTask(ctx, req.TaskID); err != nil {
		zlog.Info().Interface("error", err.Error()).Msg("Failed delete task")
		resp.Message = constant.ErrDelete
		return resp, err
	}

	if err = u.taskRepo.DeleteTask(ctx, req.TaskID); err != nil {
		zlog.Info().Interface("error", err.Error()).Msg("Failed delete task")
		resp.Message = constant.ErrDelete
		return resp, err
	}

	resp.Message = constant.SuccessUpdate
	resp.Data.Status = true
	return resp, nil
}
