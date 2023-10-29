package model

type CreateTaskRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	File        string `json:"file"`
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
	ID          string `json:"id" validate:"required"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	File        string `json:"file"`
}

type DeleteRequest struct {
	ProductID int64 `json:"id" validate:"required"`
}
