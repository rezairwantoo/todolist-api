package subtask

import (
	"context"
	"reza/todolist-api/model"
	"reza/todolist-api/model/constant"

	zlog "github.com/rs/zerolog/log"
)

func (u *Usecase) Create(ctx context.Context, req model.CreateSubTaskRequest) (model.CreateResponse, error) {
	var (
		CreateRespData model.ResponseDataCreate
		err            error
		resp           model.CreateResponse
	)

	createTaskRequest := model.CreateTaskRequest{
		Title:       req.Title,
		Description: req.Description,
		File:        req.File,
		TaskID:      req.TaskID,
	}
	if err = u.subtaskRepo.CreateTask(ctx, createTaskRequest); err != nil {
		zlog.Info().Interface("error", err.Error()).Msg("Failed Create products")
		resp.Message = constant.ErrCreate
		return resp, err
	}

	CreateRespData.Status = true
	return model.CreateResponse{
		Message: constant.SuccessCreate,
		Data:    CreateRespData,
	}, nil
}
