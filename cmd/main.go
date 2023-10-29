package main

import (
	"context"
	"errors"
	"log"
	"reza/todolist-api/config"

	"reza/todolist-api/endpoint"
	taskUc "reza/todolist-api/usecase/task"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	ctx := context.Background()
	config.LoadConfigFile(ctx)
	settings, err := config.NewSettings(ctx)
	if err != nil {
		errWrap := errors.New("initialize settings, err: " + err.Error())
		log.Fatalln("initialize settings error", errWrap)
	}

	settings.Load(settings.SetPostgresRepo(settings))
	usecaseTask := taskUc.NewTaskUsecase(settings.Gorm)
	e.POST("/task", endpoint.MakeCreateTaskEndpoint(usecaseTask))
	e.GET("/task/:id", endpoint.MakeDetailProductEndpoint(usecaseTask))
	e.GET("/task", endpoint.MakeListTaksEndpoint(usecaseTask))
	// e.PUT("/product/:id", endpoint.MakeUpdateProductEndpoint(usecaseProducts))
	// e.DELETE("/product/:id", endpoint.MakeDeleteProductEndpoint(usecaseProducts))
	e.Logger.Fatal(e.Start(":1323"))
}
