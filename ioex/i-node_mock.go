// Code generated by MockGen. DO NOT EDIT.
// Source: ioex\i-node.go

// Package ioex is a generated GoMock package.
package ioex

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockINode is a mock of INode interface
type MockINode struct {
	ctrl     *gomock.Controller
	recorder *MockINodeMockRecorder
}

// MockINodeMockRecorder is the mock recorder for MockINode
type MockINodeMockRecorder struct {
	mock *MockINode
}

// NewMockINode creates a new mock instance
func NewMockINode(ctrl *gomock.Controller) *MockINode {
	mock := &MockINode{ctrl: ctrl}
	mock.recorder = &MockINodeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockINode) EXPECT() *MockINodeMockRecorder {
	return m.recorder
}

// GetName mocks base method
func (m *MockINode) GetName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetName indicates an expected call of GetName
func (mr *MockINodeMockRecorder) GetName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetName", reflect.TypeOf((*MockINode)(nil).GetName))
}

// GetParent mocks base method
func (m *MockINode) GetParent() IDirectory {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetParent")
	ret0, _ := ret[0].(IDirectory)
	return ret0
}

// GetParent indicates an expected call of GetParent
func (mr *MockINodeMockRecorder) GetParent() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetParent", reflect.TypeOf((*MockINode)(nil).GetParent))
}

// GetPath mocks base method
func (m *MockINode) GetPath() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPath")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetPath indicates an expected call of GetPath
func (mr *MockINodeMockRecorder) GetPath() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPath", reflect.TypeOf((*MockINode)(nil).GetPath))
}

// IsExist mocks base method
func (m *MockINode) IsExist() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsExist")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsExist indicates an expected call of IsExist
func (mr *MockINodeMockRecorder) IsExist() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsExist", reflect.TypeOf((*MockINode)(nil).IsExist))
}

// Move mocks base method
func (m *MockINode) Move(dstPath string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Move", dstPath)
	ret0, _ := ret[0].(error)
	return ret0
}

// Move indicates an expected call of Move
func (mr *MockINodeMockRecorder) Move(dstPath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Move", reflect.TypeOf((*MockINode)(nil).Move), dstPath)
}

// Remove mocks base method
func (m *MockINode) Remove() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove")
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove
func (mr *MockINodeMockRecorder) Remove() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockINode)(nil).Remove))
}
