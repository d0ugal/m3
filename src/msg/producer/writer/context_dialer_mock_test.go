// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/m3db/m3/src/msg/producer/writer/consumer_writer_test.go

// Copyright (c) 2022 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Package writer is a generated GoMock package.
package writer

import (
	"context"
	"net"
	"reflect"

	"github.com/golang/mock/gomock"
)

// MockcontextDialer is a mock of contextDialer interface.
type MockcontextDialer struct {
	ctrl     *gomock.Controller
	recorder *MockcontextDialerMockRecorder
}

// MockcontextDialerMockRecorder is the mock recorder for MockcontextDialer.
type MockcontextDialerMockRecorder struct {
	mock *MockcontextDialer
}

// NewMockcontextDialer creates a new mock instance.
func NewMockcontextDialer(ctrl *gomock.Controller) *MockcontextDialer {
	mock := &MockcontextDialer{ctrl: ctrl}
	mock.recorder = &MockcontextDialerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockcontextDialer) EXPECT() *MockcontextDialerMockRecorder {
	return m.recorder
}

// DialContext mocks base method.
func (m *MockcontextDialer) DialContext(ctx context.Context, network, addr string) (net.Conn, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DialContext", ctx, network, addr)
	ret0, _ := ret[0].(net.Conn)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DialContext indicates an expected call of DialContext.
func (mr *MockcontextDialerMockRecorder) DialContext(ctx, network, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DialContext", reflect.TypeOf((*MockcontextDialer)(nil).DialContext), ctx, network, addr)
}
