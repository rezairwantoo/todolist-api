package subtask

import (
	repository "reza/todolist-api/repository/gorm"
)

// Usecase contain all methods for usecase
type Usecase struct {
	subtaskRepo repository.TaskRepository
}

func NewSubTaskUsecase(
	repositoryPostgres repository.TaskRepository,
) SubTaskUsecase {
	return &Usecase{
		subtaskRepo: repositoryPostgres,
	}
}
