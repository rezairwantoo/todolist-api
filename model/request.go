package model

type CreateTaskRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	File        string `json:"file"`
	TaskID      int64
}

type DetailTaskRequest struct {
	TaskID int64 `json:"id" validate:"required"`
}

type ListRequest struct {
	Limit  int32  `json:"limit"`
	Page   int32  `json:"page"`
	Search string `json:"search"`
	Offset int32  `json:"offset"`
}

type UpdateTaskRequest struct {
	ID          int64  `json:"id" validate:"required"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	File        string `json:"file"`
}

type DeleteTaskRequest struct {
	TaskID int64 `json:"id" validate:"required"`
}

type CreateSubTaskRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	File        string `json:"file"`
	TaskID      int64  `json:"task_id" validate:"required"`
}

type DetailSubTaskRequest struct {
	TaskID    int64 `json:"id" validate:"required"`
	SubTaskID int64 `json:"subtask_id" validate:"required"`
}

type ListSubTaskRequest struct {
	TaskID int64  `json:"task_id"`
	Limit  int32  `json:"limit"`
	Page   int32  `json:"page"`
	Search string `json:"search"`
	Offset int32  `json:"offset"`
}

type UpdateSubTaskRequest struct {
	TaskID      int64  `json:"task_id" validate:"required"`
	ID          int64  `json:"id" validate:"required"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	File        string `json:"file"`
}

type DeleteSubTaskRequest struct {
	ID     int64 `json:"id" validate:"required"`
	TaskID int64 `json:"task_id" validate:"required"`
}
