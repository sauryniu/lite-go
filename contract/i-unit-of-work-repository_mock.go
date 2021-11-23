// Code generated by MockGen. DO NOT EDIT.
// Source: contract\i-unit-of-work-repository.go

// Package contract is a generated GoMock package.
package contract

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIUnitOfWorkRepository is a mock of IUnitOfWorkRepository interface.
type MockIUnitOfWorkRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIUnitOfWorkRepositoryMockRecorder
}

// MockIUnitOfWorkRepositoryMockRecorder is the mock recorder for MockIUnitOfWorkRepository.
type MockIUnitOfWorkRepositoryMockRecorder struct {
	mock *MockIUnitOfWorkRepository
}

// NewMockIUnitOfWorkRepository creates a new mock instance.
func NewMockIUnitOfWorkRepository(ctrl *gomock.Controller) *MockIUnitOfWorkRepository {
	mock := &MockIUnitOfWorkRepository{ctrl: ctrl}
	mock.recorder = &MockIUnitOfWorkRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUnitOfWorkRepository) EXPECT() *MockIUnitOfWorkRepositoryMockRecorder {
	return m.recorder
}

// Commit mocks base method.
func (m *MockIUnitOfWorkRepository) Commit() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit")
	ret0, _ := ret[0].(error)
	return ret0
}

// Commit indicates an expected call of Commit.
func (mr *MockIUnitOfWorkRepositoryMockRecorder) Commit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockIUnitOfWorkRepository)(nil).Commit))
}

// RegisterAdd mocks base method.
func (m *MockIUnitOfWorkRepository) RegisterAdd(entry IDbModel) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterAdd", entry)
}

// RegisterAdd indicates an expected call of RegisterAdd.
func (mr *MockIUnitOfWorkRepositoryMockRecorder) RegisterAdd(entry interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterAdd", reflect.TypeOf((*MockIUnitOfWorkRepository)(nil).RegisterAdd), entry)
}

// RegisterRemove mocks base method.
func (m *MockIUnitOfWorkRepository) RegisterRemove(entry IDbModel) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterRemove", entry)
}

// RegisterRemove indicates an expected call of RegisterRemove.
func (mr *MockIUnitOfWorkRepositoryMockRecorder) RegisterRemove(entry interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterRemove", reflect.TypeOf((*MockIUnitOfWorkRepository)(nil).RegisterRemove), entry)
}

// RegisterSave mocks base method.
func (m *MockIUnitOfWorkRepository) RegisterSave(entry IDbModel) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterSave", entry)
}

// RegisterSave indicates an expected call of RegisterSave.
func (mr *MockIUnitOfWorkRepositoryMockRecorder) RegisterSave(entry interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterSave", reflect.TypeOf((*MockIUnitOfWorkRepository)(nil).RegisterSave), entry)
}
