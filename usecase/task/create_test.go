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

var _ = Describe("Create Task", func() {
	var (
		mockPostgre *mocks.MockTaskRepository
		mockUsecase TaskUsecase
		ctx         context.Context
		err         error
	)

	BeforeEach(func() {
		mockCtrl := gomock.NewController(GinkgoT())
		mockCtrl.Finish()
		ctx = context.Background()
		mockPostgre = mocks.NewMockTaskRepository(mockCtrl)
		mockUsecase = NewTaskUsecase(mockPostgre)
		err = errors.New("errors")
	})

	Describe("Create Usecase", func() {
		Context("When request for Create", func() {
			It("success", func() {
				req := model.CreateTaskRequest{
					Title:       "test title 1",
					Description: "test title 1",
					File:        "testfile1.txt",
				}

				mockPostgre.EXPECT().CreateTask(ctx, req).Return(nil)
				_, err := mockUsecase.Create(ctx, req)
				Expect(err).To(BeNil())
			})
		})
		Context("When request for Create has failed when storing to db", func() {
			It("success", func() {
				req := model.CreateTaskRequest{
					Title:       "test title 1",
					Description: "test title 1",
					File:        "testfile1.txt",
				}

				mockPostgre.EXPECT().CreateTask(ctx, req).Return(err)
				_, err := mockUsecase.Create(ctx, req)
				Expect(err).ToNot(BeNil())
			})
		})
	})
})
