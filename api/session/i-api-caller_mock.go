// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ahl5esoft/lite-go/api/session (interfaces: IAPICaller)

// Package session is a generated GoMock package.
package session

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	time "time"
)

// MockIAPICaller is a mock of IAPICaller interface
type MockIAPICaller struct {
	ctrl     *gomock.Controller
	recorder *MockIAPICallerMockRecorder
}

// MockIAPICallerMockRecorder is the mock recorder for MockIAPICaller
type MockIAPICallerMockRecorder struct {
	mock *MockIAPICaller
}

// NewMockIAPICaller creates a new mock instance
func NewMockIAPICaller(ctrl *gomock.Controller) *MockIAPICaller {
	mock := &MockIAPICaller{ctrl: ctrl}
	mock.recorder = &MockIAPICallerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIAPICaller) EXPECT() *MockIAPICallerMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockIAPICaller) Get(arg0 string, arg1 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Get indicates an expected call of Get
func (mr *MockIAPICallerMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockIAPICaller)(nil).Get), arg0, arg1)
}

// Set mocks base method
func (m *MockIAPICaller) Set(arg0 interface{}, arg1, arg2 time.Duration) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Set indicates an expected call of Set
func (mr *MockIAPICallerMockRecorder) Set(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockIAPICaller)(nil).Set), arg0, arg1, arg2)
}