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
		TodoID:      req.TaskID,
	}

	return p.Conn.Create(task).Error
}

func (p *Repository) DetailTask(ctx context.Context, req model.DetailTaskRequest) (*model.Task, error) {
	var err error
	task := &model.Task{}

	err = p.Conn.First(task, req.TaskID).Error
	return task, err
}

func (p *Repository) ListTask(pagination model.Pagination, search string) (*model.Pagination, error) {
	var task []*model.Task

	err := p.Conn.Scopes(paginate(task, &pagination, p.Conn)).Where("todo_id = 0 AND (title ilike ? OR description ilike ?)", "%"+search+"%", "%"+search+"%").Find(&task).Error
	pagination.Rows = task

	return &pagination, err
}

func (p *Repository) UpdateTask(ctx context.Context, req model.Task) error {
	now, _ := helpers.GetNow()
	req.UpdatedAt = &now
	return p.Conn.Updates(req).Where("id = ?", req.ID).Error
}

func (p *Repository) DeleteTask(ctx context.Context, taskID int64) error {
	return p.Conn.Delete(model.Task{}, taskID).Error
}

func (p *Repository) DeleteSubTask(ctx context.Context, taskID int64) error {
	var subTask []model.Task
	return p.Conn.Where("todo_id = ?", taskID).Delete(&subTask).Error
}
