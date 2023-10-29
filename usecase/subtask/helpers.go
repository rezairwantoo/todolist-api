package subtask

import (
	"context"
	"reza/todolist-api/model"
	"reza/todolist-api/model/constant"

	zlog "github.com/rs/zerolog/log"
)

func (u *Usecase) CheckDetailSubtask(ctx context.Context, req model.DetailSubTaskRequest) (*model.Task, error) {
	var (
		task    *model.Task
		subtask *model.Task
		err     error
		resp    model.DetailResponse
	)

	detailReq := model.DetailTaskRequest{
		TaskID: req.TaskID,
	}

	if task, err = u.subtaskRepo.DetailTask(ctx, detailReq); err != nil {
		zlog.Info().Interface("error", err.Error()).Msg("Failed Get products")
		resp.Message = constant.ErrMsgNotFoundDefault
		return nil, err
	}

	detailReq.TaskID = req.SubTaskID
	if subtask, err = u.subtaskRepo.DetailTask(ctx, detailReq); err != nil {
		zlog.Info().Interface("error", err.Error()).Msg("Failed Get products")
		resp.Message = constant.ErrMsgNotFoundDefault
		return nil, err
	}

	if subtask.TodoID != task.ID {
		resp.Message = constant.ErrDetailInvalidTaskID
		return nil, err
	}

	return subtask, nil
}
