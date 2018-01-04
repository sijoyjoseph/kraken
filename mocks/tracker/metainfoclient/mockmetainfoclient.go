// Code generated by MockGen. DO NOT EDIT.
// Source: code.uber.internal/infra/kraken/tracker/metainfoclient (interfaces: Client)

// Package mockmetainfoclient is a generated GoMock package.
package mockmetainfoclient

import (
	torlib "code.uber.internal/infra/kraken/torlib"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockClient is a mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// Download mocks base method
func (m *MockClient) Download(arg0, arg1 string) (*torlib.MetaInfo, error) {
	ret := m.ctrl.Call(m, "Download", arg0, arg1)
	ret0, _ := ret[0].(*torlib.MetaInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Download indicates an expected call of Download
func (mr *MockClientMockRecorder) Download(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Download", reflect.TypeOf((*MockClient)(nil).Download), arg0, arg1)
}
