package repository

import (
	"gorm.io/gorm"
)

type Repository struct {
	Conn *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Conn: db,
	}
}
