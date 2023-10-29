package endpoint

import (
	"reza/todolist-api/mocks"
	taskUc "reza/todolist-api/usecase/task"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Create Products", func() {
	var (
		mockPostgre *mocks.MockTaskRepository
		mockUsecase taskUc.TaskUsecase
	)

	BeforeEach(func() {
		mockCtrl := gomock.NewController(GinkgoT())
		mockCtrl.Finish()
		mockPostgre = mocks.NewMockTaskRepository(mockCtrl)
		mockUsecase = taskUc.NewTaskUsecase(mockPostgre)
	})

	Describe("Create Usecase", func() {
		Context("When request for Create", func() {
			It("success", func() {
				test := MakeCreateTaskEndpoint(mockUsecase)
				Expect(test).ToNot(BeNil())
			})
		})
	})
})
