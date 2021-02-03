// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ahl5esoft/lite-go/api (interfaces: IAPI)

// Package api is a generated GoMock package.
package api

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockIAPI is a mock of IAPI interface
type MockIAPI struct {
	ctrl     *gomock.Controller
	recorder *MockIAPIMockRecorder
}

// MockIAPIMockRecorder is the mock recorder for MockIAPI
type MockIAPIMockRecorder struct {
	mock *MockIAPI
}

// NewMockIAPI creates a new mock instance
func NewMockIAPI(ctrl *gomock.Controller) *MockIAPI {
	mock := &MockIAPI{ctrl: ctrl}
	mock.recorder = &MockIAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIAPI) EXPECT() *MockIAPIMockRecorder {
	return m.recorder
}

// Call mocks base method
func (m *MockIAPI) Call() (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Call")
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Call indicates an expected call of Call
func (mr *MockIAPIMockRecorder) Call() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Call", reflect.TypeOf((*MockIAPI)(nil).Call))
}
