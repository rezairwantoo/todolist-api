package task

import (
	repository "reza/todolist-api/repository/gorm"
)

// Usecase contain all methods for usecase
type Usecase struct {
	taskRepo repository.TaskRepository
}

func NewTaskUsecase(
	repositoryPostgres repository.TaskRepository,
) TaskUsecase {
	return &Usecase{
		taskRepo: repositoryPostgres,
	}
}
