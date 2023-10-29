package repository

import (
	"gorm.io/gorm"
)

type PostgresRepository struct {
	Conn *gorm.DB
}

func NewRepository(db *gorm.DB) *PostgresRepository {
	return &PostgresRepository{
		Conn: db,
	}
}
