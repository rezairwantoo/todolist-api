package subtask

import (
	"context"
	"reza/todolist-api/model"
	"reza/todolist-api/model/constant"

	zlog "github.com/rs/zerolog/log"
)

func (u *Usecase) Update(ctx context.Context, req model.UpdateSubTaskRequest) (model.UpdateResponse, error) {
	var (
		subtask   *model.Task
		err       error
		resp      model.UpdateResponse
		detailReq model.DetailSubTaskRequest
	)

	detailReq.SubTaskID = req.ID
	detailReq.TaskID = req.TaskID
	if subtask, err = u.CheckDetailSubtask(ctx, detailReq); err != nil {
		return resp, err
	}

	subtask.Title = req.Title
	subtask.Description = req.Description
	subtask.File = req.File
	if err = u.subtaskRepo.UpdateTask(ctx, *subtask); err != nil {
		zlog.Info().Interface("error", err.Error()).Msg("Failed update products")
		resp.Message = constant.ErrUpdate
		return resp, err
	}

	resp.Message = constant.SuccessUpdate
	resp.Data.Status = true
	return resp, nil
}
