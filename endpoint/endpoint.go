package endpoint

import (
	"net/http"
	"reza/todolist-api/model"
	taskUc "reza/todolist-api/usecase/task"

	"github.com/labstack/echo/v4"
	zlog "github.com/rs/zerolog/log"
)

func MakeCreateTaskEndpoint(u taskUc.TaskUsecase) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			createRequest model.CreateTaskRequest
			err           error
			resp          model.CreateResponse
		)

		if err = c.Bind(&createRequest); err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}

		if err = c.Validate(createRequest); err != nil {
			zlog.Info().Interface("error", err).Msg("Validate Param Create")
			return err
		}

		if resp, err = u.Create(c.Request().Context(), createRequest); err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, resp)
	}
}
