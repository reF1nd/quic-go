// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/quic-go/quic-go (interfaces: SendConn)
//
// Generated by this command:
//
//	mockgen -typed -build_flags=-tags=gomock -package quic -self_package github.com/quic-go/quic-go -destination mock_send_conn_test.go github.com/quic-go/quic-go SendConn
//

// Package quic is a generated GoMock package.
package quic

import (
	net "net"
	reflect "reflect"

	protocol "github.com/quic-go/quic-go/internal/protocol"
	gomock "go.uber.org/mock/gomock"
)

// MockSendConn is a mock of SendConn interface.
type MockSendConn struct {
	ctrl     *gomock.Controller
	recorder *MockSendConnMockRecorder
}

// MockSendConnMockRecorder is the mock recorder for MockSendConn.
type MockSendConnMockRecorder struct {
	mock *MockSendConn
}

// NewMockSendConn creates a new mock instance.
func NewMockSendConn(ctrl *gomock.Controller) *MockSendConn {
	mock := &MockSendConn{ctrl: ctrl}
	mock.recorder = &MockSendConnMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSendConn) EXPECT() *MockSendConnMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockSendConn) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockSendConnMockRecorder) Close() *MockSendConnCloseCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockSendConn)(nil).Close))
	return &MockSendConnCloseCall{Call: call}
}

// MockSendConnCloseCall wrap *gomock.Call
type MockSendConnCloseCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSendConnCloseCall) Return(arg0 error) *MockSendConnCloseCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSendConnCloseCall) Do(f func() error) *MockSendConnCloseCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSendConnCloseCall) DoAndReturn(f func() error) *MockSendConnCloseCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// LocalAddr mocks base method.
func (m *MockSendConn) LocalAddr() net.Addr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LocalAddr")
	ret0, _ := ret[0].(net.Addr)
	return ret0
}

// LocalAddr indicates an expected call of LocalAddr.
func (mr *MockSendConnMockRecorder) LocalAddr() *MockSendConnLocalAddrCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LocalAddr", reflect.TypeOf((*MockSendConn)(nil).LocalAddr))
	return &MockSendConnLocalAddrCall{Call: call}
}

// MockSendConnLocalAddrCall wrap *gomock.Call
type MockSendConnLocalAddrCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSendConnLocalAddrCall) Return(arg0 net.Addr) *MockSendConnLocalAddrCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSendConnLocalAddrCall) Do(f func() net.Addr) *MockSendConnLocalAddrCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSendConnLocalAddrCall) DoAndReturn(f func() net.Addr) *MockSendConnLocalAddrCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// RemoteAddr mocks base method.
func (m *MockSendConn) RemoteAddr() net.Addr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoteAddr")
	ret0, _ := ret[0].(net.Addr)
	return ret0
}

// RemoteAddr indicates an expected call of RemoteAddr.
func (mr *MockSendConnMockRecorder) RemoteAddr() *MockSendConnRemoteAddrCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoteAddr", reflect.TypeOf((*MockSendConn)(nil).RemoteAddr))
	return &MockSendConnRemoteAddrCall{Call: call}
}

// MockSendConnRemoteAddrCall wrap *gomock.Call
type MockSendConnRemoteAddrCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSendConnRemoteAddrCall) Return(arg0 net.Addr) *MockSendConnRemoteAddrCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSendConnRemoteAddrCall) Do(f func() net.Addr) *MockSendConnRemoteAddrCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSendConnRemoteAddrCall) DoAndReturn(f func() net.Addr) *MockSendConnRemoteAddrCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SetRemoteAddr mocks base method.
func (m *MockSendConn) SetRemoteAddr(arg0 net.Addr) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetRemoteAddr", arg0)
}

// SetRemoteAddr indicates an expected call of SetRemoteAddr.
func (mr *MockSendConnMockRecorder) SetRemoteAddr(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetRemoteAddr", reflect.TypeOf((*MockSendConn)(nil).SetRemoteAddr), arg0)
}

// Write mocks base method.
func (m *MockSendConn) Write(arg0 []byte, arg1 uint16, arg2 protocol.ECN) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Write indicates an expected call of Write.
func (mr *MockSendConnMockRecorder) Write(arg0, arg1, arg2 any) *MockSendConnWriteCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockSendConn)(nil).Write), arg0, arg1, arg2)
	return &MockSendConnWriteCall{Call: call}
}

// MockSendConnWriteCall wrap *gomock.Call
type MockSendConnWriteCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSendConnWriteCall) Return(arg0 error) *MockSendConnWriteCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSendConnWriteCall) Do(f func([]byte, uint16, protocol.ECN) error) *MockSendConnWriteCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSendConnWriteCall) DoAndReturn(f func([]byte, uint16, protocol.ECN) error) *MockSendConnWriteCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// capabilities mocks base method.
func (m *MockSendConn) capabilities() connCapabilities {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "capabilities")
	ret0, _ := ret[0].(connCapabilities)
	return ret0
}

// capabilities indicates an expected call of capabilities.
func (mr *MockSendConnMockRecorder) capabilities() *MockSendConncapabilitiesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "capabilities", reflect.TypeOf((*MockSendConn)(nil).capabilities))
	return &MockSendConncapabilitiesCall{Call: call}
}

// MockSendConncapabilitiesCall wrap *gomock.Call
type MockSendConncapabilitiesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSendConncapabilitiesCall) Return(arg0 connCapabilities) *MockSendConncapabilitiesCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSendConncapabilitiesCall) Do(f func() connCapabilities) *MockSendConncapabilitiesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSendConncapabilitiesCall) DoAndReturn(f func() connCapabilities) *MockSendConncapabilitiesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
