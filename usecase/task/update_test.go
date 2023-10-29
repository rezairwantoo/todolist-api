package task

import (
	"context"
	"errors"
	"reza/todolist-api/mocks"
	"reza/todolist-api/model"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Update Task", func() {
	var (
		mockTaskRepo *mocks.MockTaskRepository
		mockUsecase  TaskUsecase
		ctx          context.Context
		err          error
		task         *model.Task
	)

	BeforeEach(func() {
		mockCtrl := gomock.NewController(GinkgoT())
		mockCtrl.Finish()
		ctx = context.Background()
		mockTaskRepo = mocks.NewMockTaskRepository(mockCtrl)
		mockUsecase = NewTaskUsecase(mockTaskRepo)
		err = errors.New("errors")
	})

	Describe("Update Usecase", func() {
		Context("When request for Update", func() {
			It("success", func() {
				req := model.UpdateTaskRequest{
					ID:          1,
					Title:       "test title 1",
					Description: "test title 1",
					File:        "testfile1.txt",
				}

				task = &model.Task{
					ID:          1,
					TodoID:      0,
					Title:       "Title 1",
					Description: "Description Title 1",
					File:        "filetitle1.txt",
				}

				reqDetail := model.DetailTaskRequest{
					TaskID: req.ID,
				}
				mockTaskRepo.EXPECT().DetailTask(ctx, reqDetail).Return(task, nil)
				task.File = req.File
				task.Title = req.Title
				task.Description = req.Description
				mockTaskRepo.EXPECT().UpdateTask(ctx, *task).Return(nil)
				_, err := mockUsecase.UpdateTask(ctx, req)
				Expect(err).To(BeNil())
			})
		})
		Context("When request for Update are failed", func() {
			It("error when get detail", func() {
				req := model.UpdateTaskRequest{
					ID:          1,
					Title:       "test title 1",
					Description: "test title 1",
					File:        "testfile1.txt",
				}

				task = &model.Task{
					ID:          1,
					TodoID:      0,
					Title:       "Title 1",
					Description: "Description Title 1",
					File:        "filetitle1.txt",
				}

				reqDetail := model.DetailTaskRequest{
					TaskID: req.ID,
				}
				mockTaskRepo.EXPECT().DetailTask(ctx, reqDetail).Return(task, err)
				_, err := mockUsecase.UpdateTask(ctx, req)
				Expect(err).ToNot(BeNil())
			})
			It("error when update task", func() {
				req := model.UpdateTaskRequest{
					ID:          1,
					Title:       "test title 1",
					Description: "test title 1",
					File:        "testfile1.txt",
				}

				task = &model.Task{
					ID:          1,
					TodoID:      0,
					Title:       "Title 1",
					Description: "Description Title 1",
					File:        "filetitle1.txt",
				}

				reqDetail := model.DetailTaskRequest{
					TaskID: req.ID,
				}
				mockTaskRepo.EXPECT().DetailTask(ctx, reqDetail).Return(task, nil)
				task.File = req.File
				task.Title = req.Title
				task.Description = req.Description
				mockTaskRepo.EXPECT().UpdateTask(ctx, *task).Return(err)
				_, err := mockUsecase.UpdateTask(ctx, req)
				Expect(err).ToNot(BeNil())
			})
		})
	})
})
