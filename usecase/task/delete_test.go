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

var _ = Describe("Delete Task", func() {
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

	Describe("Delete Usecase", func() {
		Context("When request for Delete", func() {
			It("success", func() {
				req := model.DeleteTaskRequest{
					TaskID: 1,
				}

				task = &model.Task{
					ID:          1,
					TodoID:      0,
					Title:       "Title 1",
					Description: "Description Title 1",
					File:        "filetitle1.txt",
				}

				reqDetail := model.DetailTaskRequest{
					TaskID: req.TaskID,
				}
				mockTaskRepo.EXPECT().DetailTask(ctx, reqDetail).Return(task, nil)
				mockTaskRepo.EXPECT().DeleteSubTask(ctx, req.TaskID).Return(nil)
				mockTaskRepo.EXPECT().DeleteTask(ctx, req.TaskID).Return(nil)
				_, err := mockUsecase.DeleteTask(ctx, req)
				Expect(err).To(BeNil())
			})
		})
		Context("When request for Update are failed", func() {
			It("error when get detail", func() {
				req := model.DeleteTaskRequest{
					TaskID: 1,
				}

				task = &model.Task{
					ID:          1,
					TodoID:      0,
					Title:       "Title 1",
					Description: "Description Title 1",
					File:        "filetitle1.txt",
				}

				reqDetail := model.DetailTaskRequest{
					TaskID: req.TaskID,
				}
				mockTaskRepo.EXPECT().DetailTask(ctx, reqDetail).Return(task, err)
				_, err := mockUsecase.DeleteTask(ctx, req)
				Expect(err).ToNot(BeNil())
			})
			It("error when delete sub task", func() {
				req := model.DeleteTaskRequest{
					TaskID: 1,
				}

				task = &model.Task{
					ID:          1,
					TodoID:      0,
					Title:       "Title 1",
					Description: "Description Title 1",
					File:        "filetitle1.txt",
				}

				reqDetail := model.DetailTaskRequest{
					TaskID: req.TaskID,
				}
				mockTaskRepo.EXPECT().DetailTask(ctx, reqDetail).Return(task, nil)
				mockTaskRepo.EXPECT().DeleteSubTask(ctx, req.TaskID).Return(err)
				_, err := mockUsecase.DeleteTask(ctx, req)
				Expect(err).ToNot(BeNil())
			})
			It("error when delete task", func() {
				req := model.DeleteTaskRequest{
					TaskID: 1,
				}

				task = &model.Task{
					ID:          1,
					TodoID:      0,
					Title:       "Title 1",
					Description: "Description Title 1",
					File:        "filetitle1.txt",
				}

				reqDetail := model.DetailTaskRequest{
					TaskID: req.TaskID,
				}
				mockTaskRepo.EXPECT().DetailTask(ctx, reqDetail).Return(task, nil)
				mockTaskRepo.EXPECT().DeleteSubTask(ctx, req.TaskID).Return(nil)
				mockTaskRepo.EXPECT().DeleteTask(ctx, req.TaskID).Return(err)
				_, err := mockUsecase.DeleteTask(ctx, req)
				Expect(err).ToNot(BeNil())
			})
		})
	})
})
