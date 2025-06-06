// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mfreeman451/goflow2/v2/transport (interfaces: TransportDriver,TransportInterface)
//
// Generated by this command:
//
//	mockgen -destination=mock_transport.go -package=transport github.com/mfreeman451/goflow2/v2/transport TransportDriver,TransportInterface
//

// Package transport is a generated GoMock package.
package transport

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockTransportDriver is a mock of TransportDriver interface.
type MockTransportDriver struct {
	ctrl     *gomock.Controller
	recorder *MockTransportDriverMockRecorder
	isgomock struct{}
}

// MockTransportDriverMockRecorder is the mock recorder for MockTransportDriver.
type MockTransportDriverMockRecorder struct {
	mock *MockTransportDriver
}

// NewMockTransportDriver creates a new mock instance.
func NewMockTransportDriver(ctrl *gomock.Controller) *MockTransportDriver {
	mock := &MockTransportDriver{ctrl: ctrl}
	mock.recorder = &MockTransportDriverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransportDriver) EXPECT() *MockTransportDriverMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockTransportDriver) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockTransportDriverMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockTransportDriver)(nil).Close))
}

// Init mocks base method.
func (m *MockTransportDriver) Init() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init")
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init.
func (mr *MockTransportDriverMockRecorder) Init() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockTransportDriver)(nil).Init))
}

// Prepare mocks base method.
func (m *MockTransportDriver) Prepare() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Prepare")
	ret0, _ := ret[0].(error)
	return ret0
}

// Prepare indicates an expected call of Prepare.
func (mr *MockTransportDriverMockRecorder) Prepare() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Prepare", reflect.TypeOf((*MockTransportDriver)(nil).Prepare))
}

// Send mocks base method.
func (m *MockTransportDriver) Send(key, data []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", key, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockTransportDriverMockRecorder) Send(key, data any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockTransportDriver)(nil).Send), key, data)
}

// MockTransportInterface is a mock of TransportInterface interface.
type MockTransportInterface struct {
	ctrl     *gomock.Controller
	recorder *MockTransportInterfaceMockRecorder
	isgomock struct{}
}

// MockTransportInterfaceMockRecorder is the mock recorder for MockTransportInterface.
type MockTransportInterfaceMockRecorder struct {
	mock *MockTransportInterface
}

// NewMockTransportInterface creates a new mock instance.
func NewMockTransportInterface(ctrl *gomock.Controller) *MockTransportInterface {
	mock := &MockTransportInterface{ctrl: ctrl}
	mock.recorder = &MockTransportInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransportInterface) EXPECT() *MockTransportInterfaceMockRecorder {
	return m.recorder
}

// Send mocks base method.
func (m *MockTransportInterface) Send(key, data []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", key, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockTransportInterfaceMockRecorder) Send(key, data any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockTransportInterface)(nil).Send), key, data)
}
