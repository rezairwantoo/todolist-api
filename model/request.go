package model

type CreateTaskRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	File        string `json:"file"`
}

type DetailRequest struct {
	ProductID int64 `json:"id" validate:"required"`
}

type ListRequest struct {
	Limit  int32  `json:"limit"`
	Page   int32  `json:"page"`
	Search string `json:"search"`
	Offset int32  `json:"offset"`
}

type UpdateRequest struct {
	Item      string `json:"item" validate:"required"`
	Price     int64  `json:"price" validate:"required"`
	ProductID int64  `json:"id"`
	IsActive  bool   `json:"is_active" validate:"required"`
}

type DeleteRequest struct {
	ProductID int64 `json:"id" validate:"required"`
}
