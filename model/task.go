package model

import "time"

type Task struct {
	ID          int64      `gorm:"column:id;primary_key" json:"id"`
	TodoID      int64      `gorm:"column:todo_id" json:"todo_id"`
	Title       string     `gorm:"column:title" json:"title"`
	Description string     `gorm:"column:description" json:"description"`
	File        string     `gorm:"column:file" json:"file"`
	CreatedAt   *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   *time.Time `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt   *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}
