// Code generated by MockGen. DO NOT EDIT.
// Source: taskstore.go

// Package taskstore_test is a generated GoMock package.
package taskstore_test

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	storages "github.com/kulti/task-list/server/internal/storages"
	reflect "reflect"
)

// MockDBStore is a mock of dbStore interface
type MockDBStore struct {
	ctrl     *gomock.Controller
	recorder *MockDBStoreMockRecorder
}

// MockDBStoreMockRecorder is the mock recorder for MockDBStore
type MockDBStoreMockRecorder struct {
	mock *MockDBStore
}

// NewMockDBStore creates a new mock instance
func NewMockDBStore(ctrl *gomock.Controller) *MockDBStore {
	mock := &MockDBStore{ctrl: ctrl}
	mock.recorder = &MockDBStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDBStore) EXPECT() *MockDBStoreMockRecorder {
	return m.recorder
}

// UpdateTask mocks base method
func (m *MockDBStore) UpdateTask(ctx context.Context, taskID int64, fn storages.UpdateTaskFn) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTask", ctx, taskID, fn)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTask indicates an expected call of UpdateTask
func (mr *MockDBStoreMockRecorder) UpdateTask(ctx, taskID, fn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTask", reflect.TypeOf((*MockDBStore)(nil).UpdateTask), ctx, taskID, fn)
}

// PostponeTask mocks base method
func (m *MockDBStore) PostponeTask(ctx context.Context, taskID int64, fn storages.PostponeTaskFn) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostponeTask", ctx, taskID, fn)
	ret0, _ := ret[0].(error)
	return ret0
}

// PostponeTask indicates an expected call of PostponeTask
func (mr *MockDBStoreMockRecorder) PostponeTask(ctx, taskID, fn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostponeTask", reflect.TypeOf((*MockDBStore)(nil).PostponeTask), ctx, taskID, fn)
}

// DeleteTask mocks base method
func (m *MockDBStore) DeleteTask(ctx context.Context, taskID int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTask", ctx, taskID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTask indicates an expected call of DeleteTask
func (mr *MockDBStoreMockRecorder) DeleteTask(ctx, taskID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTask", reflect.TypeOf((*MockDBStore)(nil).DeleteTask), ctx, taskID)
}
