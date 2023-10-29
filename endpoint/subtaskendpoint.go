package endpoint

import (
	b64 "encoding/base64"
	"net/http"
	"reza/todolist-api/model"
	subtaskkUc "reza/todolist-api/usecase/subtask"
	"strconv"

	"github.com/labstack/echo/v4"
	zlog "github.com/rs/zerolog/log"
)

func MakeCreateSubTaskEndpoint(u subtaskkUc.SubTaskUsecase) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			createRequest model.CreateSubTaskRequest
			err           error
			resp          model.CreateResponse
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

		createRequest.TaskID = int64(taskID)
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

func MakeDetailSubTaskEndpoint(u subtaskkUc.SubTaskUsecase) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			detailRequest model.DetailSubTaskRequest
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

		base64ID = c.Param("subid")
		if sDec, err = b64.URLEncoding.DecodeString(base64ID); sDec == nil || err != nil {
			zlog.Info().Interface("error", err).Msg("Validate Param Detail")
			resp.Message = "invalid parameter"
			return c.JSON(http.StatusBadRequest, resp)
		}

		strDec = string(sDec)
		if taskID, err = strconv.Atoi(strDec); err != nil {
			zlog.Info().Interface("error", err).Msg("Failed conv str to int")
			resp.Message = "invalid parameter"
			return c.JSON(http.StatusBadRequest, resp)
		}

		detailRequest.SubTaskID = int64(taskID)
		if resp, err = u.Detail(c.Request().Context(), detailRequest); err != nil {
			resp.Message = err.Error()
			return c.JSON(http.StatusBadRequest, resp)
		}

		return c.JSON(http.StatusOK, resp)
	}
}
