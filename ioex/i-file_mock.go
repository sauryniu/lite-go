// Code generated by MockGen. DO NOT EDIT.
// Source: ioex\i-file.go

// Package ioex is a generated GoMock package.
package ioex

import (
	gomock "github.com/golang/mock/gomock"
	os "os"
	reflect "reflect"
)

// MockIFile is a mock of IFile interface
type MockIFile struct {
	ctrl     *gomock.Controller
	recorder *MockIFileMockRecorder
}

// MockIFileMockRecorder is the mock recorder for MockIFile
type MockIFileMockRecorder struct {
	mock *MockIFile
}

// NewMockIFile creates a new mock instance
func NewMockIFile(ctrl *gomock.Controller) *MockIFile {
	mock := &MockIFile{ctrl: ctrl}
	mock.recorder = &MockIFileMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIFile) EXPECT() *MockIFileMockRecorder {
	return m.recorder
}

// GetName mocks base method
func (m *MockIFile) GetName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetName indicates an expected call of GetName
func (mr *MockIFileMockRecorder) GetName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetName", reflect.TypeOf((*MockIFile)(nil).GetName))
}

// GetParent mocks base method
func (m *MockIFile) GetParent() IDirectory {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetParent")
	ret0, _ := ret[0].(IDirectory)
	return ret0
}

// GetParent indicates an expected call of GetParent
func (mr *MockIFileMockRecorder) GetParent() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetParent", reflect.TypeOf((*MockIFile)(nil).GetParent))
}

// GetPath mocks base method
func (m *MockIFile) GetPath() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPath")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetPath indicates an expected call of GetPath
func (mr *MockIFileMockRecorder) GetPath() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPath", reflect.TypeOf((*MockIFile)(nil).GetPath))
}

// IsExist mocks base method
func (m *MockIFile) IsExist() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsExist")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsExist indicates an expected call of IsExist
func (mr *MockIFileMockRecorder) IsExist() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsExist", reflect.TypeOf((*MockIFile)(nil).IsExist))
}

// Move mocks base method
func (m *MockIFile) Move(dstPath string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Move", dstPath)
	ret0, _ := ret[0].(error)
	return ret0
}

// Move indicates an expected call of Move
func (mr *MockIFileMockRecorder) Move(dstPath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Move", reflect.TypeOf((*MockIFile)(nil).Move), dstPath)
}

// Remove mocks base method
func (m *MockIFile) Remove() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove")
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove
func (mr *MockIFileMockRecorder) Remove() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockIFile)(nil).Remove))
}

// GetExt mocks base method
func (m *MockIFile) GetExt() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExt")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetExt indicates an expected call of GetExt
func (mr *MockIFileMockRecorder) GetExt() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExt", reflect.TypeOf((*MockIFile)(nil).GetExt))
}

// GetFile mocks base method
func (m *MockIFile) GetFile() (*os.File, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFile")
	ret0, _ := ret[0].(*os.File)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFile indicates an expected call of GetFile
func (mr *MockIFileMockRecorder) GetFile() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFile", reflect.TypeOf((*MockIFile)(nil).GetFile))
}

// Read mocks base method
func (m *MockIFile) Read(data interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", data)
	ret0, _ := ret[0].(error)
	return ret0
}

// Read indicates an expected call of Read
func (mr *MockIFileMockRecorder) Read(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockIFile)(nil).Read), data)
}

// ReadJSON mocks base method
func (m *MockIFile) ReadJSON(data interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadJSON", data)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReadJSON indicates an expected call of ReadJSON
func (mr *MockIFileMockRecorder) ReadJSON(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadJSON", reflect.TypeOf((*MockIFile)(nil).ReadJSON), data)
}

// ReadYaml mocks base method
func (m *MockIFile) ReadYaml(data interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadYaml", data)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReadYaml indicates an expected call of ReadYaml
func (mr *MockIFileMockRecorder) ReadYaml(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadYaml", reflect.TypeOf((*MockIFile)(nil).ReadYaml), data)
}

// Write mocks base method
func (m *MockIFile) Write(data interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", data)
	ret0, _ := ret[0].(error)
	return ret0
}

// Write indicates an expected call of Write
func (mr *MockIFileMockRecorder) Write(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockIFile)(nil).Write), data)
}
