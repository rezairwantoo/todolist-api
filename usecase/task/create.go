package task

import (
	"context"
	"reza/todolist-api/model"
	"reza/todolist-api/model/constant"

	zlog "github.com/rs/zerolog/log"
)

func (u *Usecase) Create(ctx context.Context, req model.CreateTaskRequest) (model.CreateResponse, error) {
	var (
		CreateRespData model.ResponseDataCreate
		err            error
		resp           model.CreateResponse
	)

	if err = u.postgreSQL.CreateTask(ctx, req); err != nil {
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
