package repository

import (
	"context"
	"reza/todolist-api/model"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestCreateTask(t *testing.T) {
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

func TestDetailTask(t *testing.T) {
	mockDb, _, _ := sqlmock.New()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})

	ctx := context.Background()
	repoPost := NewRepository(db)
	repoPost.DetailTask(ctx, model.DetailTaskRequest{
		TaskID: 1,
	})
}

func TestListTask(t *testing.T) {
	mockDb, _, _ := sqlmock.New()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})
	repoPost := NewRepository(db)
	repoPost.ListTask(model.Pagination{
		Limit: 1,
		Page:  2,
	}, "test")

	repoPost.ListTask(model.Pagination{
		Limit: 0,
		Page:  0,
	}, "test")
}

func TestUpdateTask(t *testing.T) {
	mockDb, _, _ := sqlmock.New()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})

	ctx := context.Background()
	repoPost := NewRepository(db)
	repoPost.UpdateTask(ctx, model.Task{
		Title:       "testmockrepo",
		Description: "testmockrepodesc",
		File:        "testmockrepo.txt",
	})
}

func TestDeleteTask(t *testing.T) {
	mockDb, _, _ := sqlmock.New()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})

	ctx := context.Background()
	repoPost := NewRepository(db)
	repoPost.DeleteSubTask(ctx, int64(1))
	repoPost.DeleteTask(ctx, int64(1))
}

func TestListSubTask(t *testing.T) {
	mockDb, _, _ := sqlmock.New()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})
	repoPost := NewRepository(db)
	repoPost.ListSubTask(model.Pagination{
		Limit: 1,
		Page:  2,
	}, "test", int64(1))

	repoPost.ListSubTask(model.Pagination{
		Limit: 0,
		Page:  0,
	}, "test", int64(1))
}
