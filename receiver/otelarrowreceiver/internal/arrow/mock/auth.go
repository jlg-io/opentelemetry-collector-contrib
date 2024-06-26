// Code generated by MockGen. DO NOT EDIT.
// Source: go.opentelemetry.io/collector/extension/auth (interfaces: Server)
//
// Generated by this command:
//
//	mockgen -package mock go.opentelemetry.io/collector/extension/auth Server
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	component "go.opentelemetry.io/collector/component"
	gomock "go.uber.org/mock/gomock"
)

// MockServer is a mock of Server interface.
type MockServer struct {
	ctrl     *gomock.Controller
	recorder *MockServerMockRecorder
}

// MockServerMockRecorder is the mock recorder for MockServer.
type MockServerMockRecorder struct {
	mock *MockServer
}

// NewMockServer creates a new mock instance.
func NewMockServer(ctrl *gomock.Controller) *MockServer {
	mock := &MockServer{ctrl: ctrl}
	mock.recorder = &MockServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServer) EXPECT() *MockServerMockRecorder {
	return m.recorder
}

// Authenticate mocks base method.
func (m *MockServer) Authenticate(arg0 context.Context, arg1 map[string][]string) (context.Context, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Authenticate", arg0, arg1)
	ret0, _ := ret[0].(context.Context)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Authenticate indicates an expected call of Authenticate.
func (mr *MockServerMockRecorder) Authenticate(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authenticate", reflect.TypeOf((*MockServer)(nil).Authenticate), arg0, arg1)
}

// Shutdown mocks base method.
func (m *MockServer) Shutdown(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Shutdown", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Shutdown indicates an expected call of Shutdown.
func (mr *MockServerMockRecorder) Shutdown(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Shutdown", reflect.TypeOf((*MockServer)(nil).Shutdown), arg0)
}

// Start mocks base method.
func (m *MockServer) Start(arg0 context.Context, arg1 component.Host) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start.
func (mr *MockServerMockRecorder) Start(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockServer)(nil).Start), arg0, arg1)
}
