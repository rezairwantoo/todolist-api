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

var _ = Describe("Get List Task", func() {
	var (
		mockTaskRepo *mocks.MockTaskRepository
		mockUsecase  SubTaskUsecase
		ctx          context.Context
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

	Describe("Get List Usecase", func() {
		Context("When request for get list", func() {
			It("success", func() {
				req := model.ListSubTaskRequest{
					Limit:  10,
					Page:   1,
					Search: "test",
					TaskID: int64(1),
				}
				var paginate model.Pagination
				paginate.Limit = int(req.Limit)
				paginate.Page = int(req.Page)

				mockTaskRepo.EXPECT().ListSubTask(*&paginate, req.Search, req.TaskID)
				_, err := mockUsecase.List(ctx, req)
				Expect(err).To(BeNil())
			})
			It("return error when GetByID has failed", func() {
				req := model.ListSubTaskRequest{
					Limit:  10,
					Page:   1,
					Search: "test",
					TaskID: int64(1),
				}
				var paginate model.Pagination
				paginate.Limit = int(req.Limit)
				paginate.Page = int(req.Page)

				mockTaskRepo.EXPECT().ListSubTask(*&paginate, req.Search, req.TaskID).Return(&paginate, err)
				_, err := mockUsecase.List(ctx, req)
				Expect(err).ToNot(BeNil())
			})
			It("success with req page and limit is zero", func() {
				req := model.ListSubTaskRequest{
					Limit:  0,
					Page:   0,
					Search: "test",
					TaskID: int64(1),
				}
				var paginate model.Pagination
				paginate.Limit = 0
				paginate.Page = 0

				mockTaskRepo.EXPECT().ListSubTask(*&paginate, req.Search, req.TaskID)
				_, err := mockUsecase.List(ctx, req)
				Expect(err).To(BeNil())
			})
		})
	})
})
