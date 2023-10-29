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

var _ = Describe("Get Detail Task", func() {
	var (
		mockTaskRepo *mocks.MockTaskRepository
		mockUsecase  TaskUsecase
		ctx          context.Context
		task         *model.Task
		err          error
	)

	BeforeEach(func() {
		mockCtrl := gomock.NewController(GinkgoT())
		mockCtrl.Finish()
		ctx = context.Background()
		mockTaskRepo = mocks.NewMockTaskRepository(mockCtrl)
		mockUsecase = NewTaskUsecase(mockTaskRepo)
		err = errors.New("errors")
	})

	Describe("Get Detail Usecase", func() {
		Context("When request for get detail", func() {
			It("success", func() {
				req := model.DetailTaskRequest{
					TaskID: 1,
				}

				task = &model.Task{
					ID:          1,
					TodoID:      0,
					Title:       "Title 1",
					Description: "Description Title 1",
					File:        "filetitle1.txt",
				}
				mockTaskRepo.EXPECT().DetailTask(ctx, req).Return(task, nil)
				_, err := mockUsecase.Detail(ctx, req)
				Expect(err).To(BeNil())
			})
			It("return error when GetByID has failed", func() {
				req := model.DetailTaskRequest{
					TaskID: 1,
				}

				task = &model.Task{}
				mockTaskRepo.EXPECT().DetailTask(ctx, req).Return(task, err)
				_, err := mockUsecase.Detail(ctx, req)
				Expect(err).ToNot(BeNil())
			})
		})
	})
})
