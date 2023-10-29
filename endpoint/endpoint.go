package endpoint

import (
	"net/http"
	"reza/todolist-api/model"
	taskUc "reza/todolist-api/usecase/task"
	"strconv"

	b64 "encoding/base64"

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
			resp.Message = "invalid parameter"
			return c.JSON(http.StatusBadRequest, resp)
		}

		if err = c.Validate(createRequest); err != nil {
			zlog.Info().Interface("error", err).Msg("Validate Param Create")
			resp.Message = "invalid parameter"
			return c.JSON(http.StatusBadRequest, resp)
		}

		if resp, err = u.Create(c.Request().Context(), createRequest); err != nil {
			resp.Message = err.Error()
			return c.JSON(http.StatusInternalServerError, resp)
		}

		return c.JSON(http.StatusOK, resp)
	}
}

func MakeDetailProductEndpoint(u taskUc.TaskUsecase) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			detailRequest model.DetailTaskRequest
			err           error
			resp          model.DetailResponse
			sDec          []byte
			taskID        int
		)

		base64ID := c.Param("id")
		if sDec, err = b64.URLEncoding.DecodeString(base64ID); sDec == nil || err != nil {
			zlog.Info().Interface("error", err).Msg("Validate Param Detail")
			resp.Message = "invalid parameter"
			return c.JSON(http.StatusBadRequest, resp)
		}

		strDec := string(sDec)
		if taskID, err = strconv.Atoi(strDec); err != nil {
			zlog.Info().Interface("error", err).Msg("Failed conv str to int")
			resp.Message = "invalid parameter"
			return c.JSON(http.StatusBadRequest, resp)
		}

		detailRequest.TaskID = int64(taskID)
		if resp, err = u.Detail(c.Request().Context(), detailRequest); err != nil {
			resp.Message = err.Error()
			return c.JSON(http.StatusBadRequest, resp)
		}

		return c.JSON(http.StatusOK, resp)
	}
}

func MakeListTaksEndpoint(u taskUc.TaskUsecase) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			listRequest model.ListRequest
			err         error
			resp        *model.ListResponse
		)

		listRequest.Search = c.QueryParam("search")
		page, _ := strconv.Atoi(c.QueryParam("page"))
		listRequest.Page = int32(page)

		limit, _ := strconv.Atoi(c.QueryParam("limit"))
		listRequest.Limit = int32(limit)

		if resp, err = u.List(c.Request().Context(), listRequest); err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, resp)
	}
}

func MakeUpdateTaskEndpoint(u taskUc.TaskUsecase) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			updateRequest model.UpdateTaskRequest
			err           error
			resp          model.UpdateResponse
			sDec          []byte
			taskID        int
		)

		base64ID := c.Param("id")
		if sDec, err = b64.URLEncoding.DecodeString(base64ID); sDec == nil || err != nil {
			zlog.Info().Interface("error", err).Msg("Validate Param Detail")
			return c.String(http.StatusBadRequest, "invalid Parameter")
		}

		strDec := string(sDec)
		if taskID, err = strconv.Atoi(strDec); err != nil {
			zlog.Info().Interface("error", err).Msg("Failed conv str to int")
			return c.String(http.StatusBadRequest, "invalid Parameter")
		}

		updateRequest.ID = int64(taskID)
		if err = c.Bind(&updateRequest); err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}

		if err = c.Validate(updateRequest); err != nil {
			zlog.Info().Interface("error", err).Msg("Validate Param Create")
			return err
		}
		if resp, err = u.UpdateTask(c.Request().Context(), updateRequest); err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, resp)
	}
}
