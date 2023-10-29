// Code generated by MockGen. DO NOT EDIT.
// Source: usecase/task/interface.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	model "reza/todolist-api/model"

	gomock "github.com/golang/mock/gomock"
)

// MockTaskUsecase is a mock of TaskUsecase interface.
type MockTaskUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockTaskUsecaseMockRecorder
}

// MockTaskUsecaseMockRecorder is the mock recorder for MockTaskUsecase.
type MockTaskUsecaseMockRecorder struct {
	mock *MockTaskUsecase
}

// NewMockTaskUsecase creates a new mock instance.
func NewMockTaskUsecase(ctrl *gomock.Controller) *MockTaskUsecase {
	mock := &MockTaskUsecase{ctrl: ctrl}
	mock.recorder = &MockTaskUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTaskUsecase) EXPECT() *MockTaskUsecaseMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockTaskUsecase) Create(ctx context.Context, req model.CreateTaskRequest) (model.CreateResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, req)
	ret0, _ := ret[0].(model.CreateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockTaskUsecaseMockRecorder) Create(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTaskUsecase)(nil).Create), ctx, req)
}
