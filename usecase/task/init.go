package task

import (
	repository "reza/todolist-api/repository/gorm"
)

// Usecase contain all methods for usecase
type Usecase struct {
	postgreSQL repository.PostgresSQLRepository
}

func NewTaskUsecase(
	repositoryPostgres repository.PostgresSQLRepository,
) TaskUsecase {
	return &Usecase{
		postgreSQL: repositoryPostgres,
	}
}
