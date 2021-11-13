// Code generated by MockGen. DO NOT EDIT.
// Source: contract\i-context-wrapper.go

// Package contract is a generated GoMock package.
package contract

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIContextWrapper is a mock of IContextWrapper interface.
type MockIContextWrapper struct {
	ctrl     *gomock.Controller
	recorder *MockIContextWrapperMockRecorder
}

// MockIContextWrapperMockRecorder is the mock recorder for MockIContextWrapper.
type MockIContextWrapperMockRecorder struct {
	mock *MockIContextWrapper
}

// NewMockIContextWrapper creates a new mock instance.
func NewMockIContextWrapper(ctrl *gomock.Controller) *MockIContextWrapper {
	mock := &MockIContextWrapper{ctrl: ctrl}
	mock.recorder = &MockIContextWrapperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIContextWrapper) EXPECT() *MockIContextWrapperMockRecorder {
	return m.recorder
}

// WithContext mocks base method.
func (m *MockIContextWrapper) WithContext(arg0 context.Context) interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithContext", arg0)
	ret0, _ := ret[0].(interface{})
	return ret0
}

// WithContext indicates an expected call of WithContext.
func (mr *MockIContextWrapperMockRecorder) WithContext(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithContext", reflect.TypeOf((*MockIContextWrapper)(nil).WithContext), arg0)
}
