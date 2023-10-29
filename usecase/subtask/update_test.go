package subtask

import (
	"context"
	"errors"
	"reza/todolist-api/mocks"
	"reza/todolist-api/model"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Get Detail Task", func() {
	var (
		mockTaskRepo *mocks.MockTaskRepository
		mockUsecase  SubTaskUsecase
		ctx          context.Context
		task         *model.Task
		err          error
	)

	BeforeEach(func() {
		mockCtrl := gomock.NewController(GinkgoT())
		mockCtrl.Finish()
		ctx = context.Background()
		mockTaskRepo = mocks.NewMockTaskRepository(mockCtrl)
		mockUsecase = NewSubTaskUsecase(mockTaskRepo)
		err = errors.New("errors")
	})

	Describe("Get Detail Usecase", func() {
		Context("When request for get detail", func() {
			It("success", func() {
				req := model.UpdateSubTaskRequest{
					ID:          2,
					TaskID:      1,
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
					TaskID: 1,
				}
				mockTaskRepo.EXPECT().DetailTask(ctx, reqDetail).Return(task, nil)

				subtask := &model.Task{
					ID:          2,
					TodoID:      1,
					Title:       "Title 1",
					Description: "Description Title 1",
					File:        "filetitle1.txt",
				}
				reqDetail.TaskID = 2
				mockTaskRepo.EXPECT().DetailTask(ctx, reqDetail).Return(subtask, nil)
				subtask.File = req.File
				subtask.Title = req.Title
				subtask.Description = req.Description
				mockTaskRepo.EXPECT().UpdateTask(ctx, *subtask).Return(nil)
				_, err := mockUsecase.Update(ctx, req)
				Expect(err).To(BeNil())
			})
			It("return error when Get detail task", func() {
				req := model.UpdateSubTaskRequest{
					ID:          2,
					TaskID:      1,
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
					TaskID: 1,
				}
				mockTaskRepo.EXPECT().DetailTask(ctx, reqDetail).Return(task, err)
				_, err := mockUsecase.Update(ctx, req)
				Expect(err).ToNot(BeNil())
			})
			It("return error when Get detail sub task", func() {
				req := model.UpdateSubTaskRequest{
					ID:          2,
					TaskID:      1,
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
					TaskID: 1,
				}
				mockTaskRepo.EXPECT().DetailTask(ctx, reqDetail).Return(task, nil)

				subtask := &model.Task{
					ID:          2,
					TodoID:      1,
					Title:       "Title 1",
					Description: "Description Title 1",
					File:        "filetitle1.txt",
				}
				reqDetail.TaskID = 2
				mockTaskRepo.EXPECT().DetailTask(ctx, reqDetail).Return(subtask, err)
				_, err := mockUsecase.Update(ctx, req)
				Expect(err).ToNot(BeNil())
			})
			It("return error when sub task not child of task", func() {
				req := model.UpdateSubTaskRequest{
					ID:          2,
					TaskID:      1,
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
					TaskID: 1,
				}
				mockTaskRepo.EXPECT().DetailTask(ctx, reqDetail).Return(task, nil)

				subtask := &model.Task{
					ID:          2,
					TodoID:      3,
					Title:       "Title 1",
					Description: "Description Title 1",
					File:        "filetitle1.txt",
				}
				reqDetail.TaskID = 2
				mockTaskRepo.EXPECT().DetailTask(ctx, reqDetail).Return(subtask, nil)
				_, err := mockUsecase.Update(ctx, req)
				Expect(err).ToNot(BeNil())
			})
			It("error when updating data", func() {
				req := model.UpdateSubTaskRequest{
					ID:          2,
					TaskID:      1,
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
					TaskID: 1,
				}
				mockTaskRepo.EXPECT().DetailTask(ctx, reqDetail).Return(task, nil)

				subtask := &model.Task{
					ID:          2,
					TodoID:      1,
					Title:       "Title 1",
					Description: "Description Title 1",
					File:        "filetitle1.txt",
				}
				reqDetail.TaskID = 2
				mockTaskRepo.EXPECT().DetailTask(ctx, reqDetail).Return(subtask, nil)
				subtask.File = req.File
				subtask.Title = req.Title
				subtask.Description = req.Description
				mockTaskRepo.EXPECT().UpdateTask(ctx, *subtask).Return(err)
				_, err := mockUsecase.Update(ctx, req)
				Expect(err).ToNot(BeNil())
			})
		})
	})
})
