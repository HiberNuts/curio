// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/filecoin-project/curio/lib/paths (interfaces: PartialFileHandler)

// Package mocks is a generated GoMock package.
package mocks

import (
	io "io"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	abi "github.com/filecoin-project/go-state-types/abi"

	partialfile "github.com/filecoin-project/curio/lib/partialfile"
	storiface "github.com/filecoin-project/curio/lib/storiface"
)

// MockPartialFileHandler is a mock of PartialFileHandler interface.
type MockPartialFileHandler struct {
	ctrl     *gomock.Controller
	recorder *MockPartialFileHandlerMockRecorder
}

// MockPartialFileHandlerMockRecorder is the mock recorder for MockPartialFileHandler.
type MockPartialFileHandlerMockRecorder struct {
	mock *MockPartialFileHandler
}

// NewMockPartialFileHandler creates a new mock instance.
func NewMockPartialFileHandler(ctrl *gomock.Controller) *MockPartialFileHandler {
	mock := &MockPartialFileHandler{ctrl: ctrl}
	mock.recorder = &MockPartialFileHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPartialFileHandler) EXPECT() *MockPartialFileHandlerMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockPartialFileHandler) Close(arg0 *partialfile.PartialFile) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockPartialFileHandlerMockRecorder) Close(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockPartialFileHandler)(nil).Close), arg0)
}

// HasAllocated mocks base method.
func (m *MockPartialFileHandler) HasAllocated(arg0 *partialfile.PartialFile, arg1 storiface.UnpaddedByteIndex, arg2 abi.UnpaddedPieceSize) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasAllocated", arg0, arg1, arg2)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HasAllocated indicates an expected call of HasAllocated.
func (mr *MockPartialFileHandlerMockRecorder) HasAllocated(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasAllocated", reflect.TypeOf((*MockPartialFileHandler)(nil).HasAllocated), arg0, arg1, arg2)
}

// OpenPartialFile mocks base method.
func (m *MockPartialFileHandler) OpenPartialFile(arg0 abi.PaddedPieceSize, arg1 string) (*partialfile.PartialFile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenPartialFile", arg0, arg1)
	ret0, _ := ret[0].(*partialfile.PartialFile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpenPartialFile indicates an expected call of OpenPartialFile.
func (mr *MockPartialFileHandlerMockRecorder) OpenPartialFile(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenPartialFile", reflect.TypeOf((*MockPartialFileHandler)(nil).OpenPartialFile), arg0, arg1)
}

// Reader mocks base method.
func (m *MockPartialFileHandler) Reader(arg0 *partialfile.PartialFile, arg1 storiface.PaddedByteIndex, arg2 abi.PaddedPieceSize) (io.Reader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reader", arg0, arg1, arg2)
	ret0, _ := ret[0].(io.Reader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Reader indicates an expected call of Reader.
func (mr *MockPartialFileHandlerMockRecorder) Reader(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reader", reflect.TypeOf((*MockPartialFileHandler)(nil).Reader), arg0, arg1, arg2)
}
