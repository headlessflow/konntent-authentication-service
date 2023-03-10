// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/claimer/claimer.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	claimer "konntent-authentication-service/pkg/claimer"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockClaimer is a mock of Claimer interface.
type MockClaimer struct {
	ctrl     *gomock.Controller
	recorder *MockClaimerMockRecorder
}

// MockClaimerMockRecorder is the mock recorder for MockClaimer.
type MockClaimerMockRecorder struct {
	mock *MockClaimer
}

// NewMockClaimer creates a new mock instance.
func NewMockClaimer(ctrl *gomock.Controller) *MockClaimer {
	mock := &MockClaimer{ctrl: ctrl}
	mock.recorder = &MockClaimerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClaimer) EXPECT() *MockClaimerMockRecorder {
	return m.recorder
}

// GetModel mocks base method.
func (m *MockClaimer) GetModel(ctx context.Context) *claimer.Model {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetModel", ctx)
	ret0, _ := ret[0].(*claimer.Model)
	return ret0
}

// GetModel indicates an expected call of GetModel.
func (mr *MockClaimerMockRecorder) GetModel(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetModel", reflect.TypeOf((*MockClaimer)(nil).GetModel), ctx)
}

// IsValid mocks base method.
func (m *MockClaimer) IsValid(ctx context.Context, rawToken []byte) ([]byte, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsValid", ctx, rawToken)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// IsValid indicates an expected call of IsValid.
func (mr *MockClaimerMockRecorder) IsValid(ctx, rawToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsValid", reflect.TypeOf((*MockClaimer)(nil).IsValid), ctx, rawToken)
}
