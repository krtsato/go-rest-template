// Code generated by MockGen. DO NOT EDIT.
// Source: abstract.go

// Package rest is a generated GoMock package.
package rest

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockAbstractServer is a mock of AbstractServer interface.
type MockAbstractServer struct {
	ctrl     *gomock.Controller
	recorder *MockAbstractServerMockRecorder
}

// MockAbstractServerMockRecorder is the mock recorder for MockAbstractServer.
type MockAbstractServerMockRecorder struct {
	mock *MockAbstractServer
}

// NewMockAbstractServer creates a new mock instance.
func NewMockAbstractServer(ctrl *gomock.Controller) *MockAbstractServer {
	mock := &MockAbstractServer{ctrl: ctrl}
	mock.recorder = &MockAbstractServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAbstractServer) EXPECT() *MockAbstractServerMockRecorder {
	return m.recorder
}

// ListenAndServe mocks base method.
func (m *MockAbstractServer) ListenAndServe(ctx context.Context, port string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListenAndServe", ctx, port)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListenAndServe indicates an expected call of ListenAndServe.
func (mr *MockAbstractServerMockRecorder) ListenAndServe(ctx, port interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListenAndServe", reflect.TypeOf((*MockAbstractServer)(nil).ListenAndServe), ctx, port)
}

// Shutdown mocks base method.
func (m *MockAbstractServer) Shutdown(ctx context.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Shutdown", ctx)
}

// Shutdown indicates an expected call of Shutdown.
func (mr *MockAbstractServerMockRecorder) Shutdown(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Shutdown", reflect.TypeOf((*MockAbstractServer)(nil).Shutdown), ctx)
}
