package repository

import (
	"context"
	"reza/todolist-api/model"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestCreateProduct(t *testing.T) {
	mockDb, _, _ := sqlmock.New()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})

	ctx := context.Background()
	repoPost := NewRepository(db)
	repoPost.CreateTask(ctx, model.CreateTaskRequest{
		Title:       "testmockrepo",
		Description: "testmockrepodesc",
		File:        "testmockrepo.txt",
	})
}
