// Code generated by MockGen. DO NOT EDIT.
// Source: repository/gorm/interface.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	model "reza/todolist-api/model"

	gomock "github.com/golang/mock/gomock"
)

// MockPostgresSQLRepository is a mock of PostgresSQLRepository interface.
type MockPostgresSQLRepository struct {
	ctrl     *gomock.Controller
	recorder *MockPostgresSQLRepositoryMockRecorder
}

// MockPostgresSQLRepositoryMockRecorder is the mock recorder for MockPostgresSQLRepository.
type MockPostgresSQLRepositoryMockRecorder struct {
	mock *MockPostgresSQLRepository
}

// NewMockPostgresSQLRepository creates a new mock instance.
func NewMockPostgresSQLRepository(ctrl *gomock.Controller) *MockPostgresSQLRepository {
	mock := &MockPostgresSQLRepository{ctrl: ctrl}
	mock.recorder = &MockPostgresSQLRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPostgresSQLRepository) EXPECT() *MockPostgresSQLRepositoryMockRecorder {
	return m.recorder
}

// CreateTask mocks base method.
func (m *MockPostgresSQLRepository) CreateTask(ctx context.Context, req model.CreateTaskRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTask", ctx, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateTask indicates an expected call of CreateTask.
func (mr *MockPostgresSQLRepositoryMockRecorder) CreateTask(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTask", reflect.TypeOf((*MockPostgresSQLRepository)(nil).CreateTask), ctx, req)
}
