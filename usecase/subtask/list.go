package subtask

import (
	"context"
	"reza/todolist-api/model"
	"reza/todolist-api/model/constant"

	zlog "github.com/rs/zerolog/log"
)

func (u *Usecase) List(ctx context.Context, req model.ListSubTaskRequest) (*model.ListResponse, error) {
	var (
		err error
	)
	resp := model.ListResponse{}
	paginate := u.preparePaginateList(req)
	if paginate, err = u.subtaskRepo.ListSubTask(*paginate, req.Search, req.TaskID); err != nil {
		zlog.Info().Interface("error", err.Error()).Msg("Failed Get products")
		resp.Message = constant.ErrInternal
		return &resp, err
	}

	resp.Data = paginate
	resp.Message = constant.SuccessDetail
	return &resp, nil
}

func (*Usecase) preparePaginateList(req model.ListSubTaskRequest) *model.Pagination {
	var paginate model.Pagination
	paginate.Limit = int(req.Limit)
	paginate.Page = int(req.Page)
	return &paginate
}
