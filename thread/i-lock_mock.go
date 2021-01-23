// Code generated by MockGen. DO NOT EDIT.
// Source: thread\i-lock.go

// Package thread is a generated GoMock package.
package thread

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	time "time"
)

// MockILock is a mock of ILock interface
type MockILock struct {
	ctrl     *gomock.Controller
	recorder *MockILockMockRecorder
}

// MockILockMockRecorder is the mock recorder for MockILock
type MockILockMockRecorder struct {
	mock *MockILock
}

// NewMockILock creates a new mock instance
func NewMockILock(ctrl *gomock.Controller) *MockILock {
	mock := &MockILock{ctrl: ctrl}
	mock.recorder = &MockILockMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockILock) EXPECT() *MockILockMockRecorder {
	return m.recorder
}

// Lock mocks base method
func (m *MockILock) Lock(arg0 string, arg1 ...interface{}) (func(), error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Lock", varargs...)
	ret0, _ := ret[0].(func())
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Lock indicates an expected call of Lock
func (mr *MockILockMockRecorder) Lock(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Lock", reflect.TypeOf((*MockILock)(nil).Lock), varargs...)
}

// SetExpire mocks base method
func (m *MockILock) SetExpire(seconds time.Duration) ILock {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetExpire", seconds)
	ret0, _ := ret[0].(ILock)
	return ret0
}

// SetExpire indicates an expected call of SetExpire
func (mr *MockILockMockRecorder) SetExpire(seconds interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetExpire", reflect.TypeOf((*MockILock)(nil).SetExpire), seconds)
}
