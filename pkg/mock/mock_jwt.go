// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/vatsal278/UserManagementService/internal/repo/jwt (interfaces: JWTService)

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	jwt "github.com/dgrijalva/jwt-go"
	gomock "github.com/golang/mock/gomock"
)

// MockJWTService is a mock of JWTService interface.
type MockJWTService struct {
	ctrl     *gomock.Controller
	recorder *MockJWTServiceMockRecorder
}

// MockJWTServiceMockRecorder is the mock recorder for MockJWTService.
type MockJWTServiceMockRecorder struct {
	mock *MockJWTService
}

// NewMockJWTService creates a new mock instance.
func NewMockJWTService(ctrl *gomock.Controller) *MockJWTService {
	mock := &MockJWTService{ctrl: ctrl}
	mock.recorder = &MockJWTServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockJWTService) EXPECT() *MockJWTServiceMockRecorder {
	return m.recorder
}

// DecodeToken mocks base method.
func (m *MockJWTService) DecodeToken(arg0 string) (jwt.MapClaims, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DecodeToken", arg0)
	ret0, _ := ret[0].(jwt.MapClaims)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DecodeToken indicates an expected call of DecodeToken.
func (mr *MockJWTServiceMockRecorder) DecodeToken(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DecodeToken", reflect.TypeOf((*MockJWTService)(nil).DecodeToken), arg0)
}

// GenerateToken mocks base method.
func (m *MockJWTService) GenerateToken(arg0 jwt.SigningMethod, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateToken", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateToken indicates an expected call of GenerateToken.
func (mr *MockJWTServiceMockRecorder) GenerateToken(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateToken", reflect.TypeOf((*MockJWTService)(nil).GenerateToken), arg0, arg1)
}

// ValidateToken mocks base method.
func (m *MockJWTService) ValidateToken(arg0 string) (*jwt.Token, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateToken", arg0)
	ret0, _ := ret[0].(*jwt.Token)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidateToken indicates an expected call of ValidateToken.
func (mr *MockJWTServiceMockRecorder) ValidateToken(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateToken", reflect.TypeOf((*MockJWTService)(nil).ValidateToken), arg0)
}