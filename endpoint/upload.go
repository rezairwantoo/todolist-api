package endpoint

import (
	"io"
	"net/http"
	"os"
	"reza/todolist-api/helpers"
	"reza/todolist-api/model"

	"github.com/labstack/echo/v4"
)

func MakeUploadFile() echo.HandlerFunc {
	return func(c echo.Context) error {
		file, err := c.FormFile("file")
		if err != nil {
			return err
		}
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		_, now := helpers.GetNow()
		// Destination
		destFile := "resources/upload/" + now + "-" + file.Filename
		dst, err := os.Create(destFile)
		if err != nil {
			return err
		}
		defer dst.Close()

		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			return err
		}

		resp := model.UploadResponse{
			Message: "Upload success",
			Data: model.UploadResponseData{
				Filename: destFile,
			},
		}
		return c.JSON(http.StatusOK, resp)
	}
}
