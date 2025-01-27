// Code generated by MockGen. DO NOT EDIT.
// Source: me.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	schema "github.com/uselagoon/lagoon-cli/internal/schema"
)

// MockMe is a mock of Me interface.
type MockMe struct {
	ctrl     *gomock.Controller
	recorder *MockMeMockRecorder
}

// MockMeMockRecorder is the mock recorder for MockMe.
type MockMeMockRecorder struct {
	mock *MockMe
}

// NewMockMe creates a new mock instance.
func NewMockMe(ctrl *gomock.Controller) *MockMe {
	mock := &MockMe{ctrl: ctrl}
	mock.recorder = &MockMeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMe) EXPECT() *MockMeMockRecorder {
	return m.recorder
}

// Me mocks base method.
func (m *MockMe) Me(ctx context.Context, user *schema.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Me", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Me indicates an expected call of Me.
func (mr *MockMeMockRecorder) Me(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Me", reflect.TypeOf((*MockMe)(nil).Me), ctx, user)
}
