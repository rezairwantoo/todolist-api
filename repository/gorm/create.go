package repository

import (
	"context"
	"reza/todolist-api/helpers"
	"reza/todolist-api/model"
)

func (p *Repository) CreateTask(ctx context.Context, req model.CreateTaskRequest) error {
	now, _ := helpers.GetNow()
	task := &model.Task{
		Title:       req.Title,
		Description: req.Description,
		File:        req.File,
		CreatedAt:   &now,
		UpdatedAt:   &now,
	}

	return p.Conn.Create(task).Error
}
