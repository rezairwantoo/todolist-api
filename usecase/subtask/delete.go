package subtask

import (
	"context"
	"reza/todolist-api/model"
	"reza/todolist-api/model/constant"

	zlog "github.com/rs/zerolog/log"
)

func (u *Usecase) Delete(ctx context.Context, req model.DeleteSubTaskRequest) (model.DeleteResponse, error) {
	var (
		err       error
		resp      model.DeleteResponse
		detailReq model.DetailSubTaskRequest
	)

	detailReq.SubTaskID = req.ID
	detailReq.TaskID = req.TaskID
	if _, err = u.CheckDetailSubtask(ctx, detailReq); err != nil {
		return resp, err
	}

	if err = u.subtaskRepo.DeleteTask(ctx, req.ID); err != nil {
		zlog.Info().Interface("error", err.Error()).Msg("Failed delete task")
		resp.Message = constant.ErrDelete
		return resp, err
	}

	resp.Message = constant.SuccessUpdate
	resp.Data.Status = true
	return resp, nil
}
