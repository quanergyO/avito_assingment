// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	types "github.com/quanergyO/avito_assingment/types"
)

// MockAuthorization is a mock of Authorization interface.
type MockAuthorization struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorizationMockRecorder
}

// MockAuthorizationMockRecorder is the mock recorder for MockAuthorization.
type MockAuthorizationMockRecorder struct {
	mock *MockAuthorization
}

// NewMockAuthorization creates a new mock instance.
func NewMockAuthorization(ctrl *gomock.Controller) *MockAuthorization {
	mock := &MockAuthorization{ctrl: ctrl}
	mock.recorder = &MockAuthorizationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthorization) EXPECT() *MockAuthorizationMockRecorder {
	return m.recorder
}

// CheckAuthData mocks base method.
func (m *MockAuthorization) CheckAuthData(username, password string) (types.UserDAO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckAuthData", username, password)
	ret0, _ := ret[0].(types.UserDAO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckAuthData indicates an expected call of CheckAuthData.
func (mr *MockAuthorizationMockRecorder) CheckAuthData(username, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckAuthData", reflect.TypeOf((*MockAuthorization)(nil).CheckAuthData), username, password)
}

// CreateUser mocks base method.
func (m *MockAuthorization) CreateUser(user types.SignInInput) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", user)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockAuthorizationMockRecorder) CreateUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockAuthorization)(nil).CreateUser), user)
}

// GenerateToken mocks base method.
func (m *MockAuthorization) GenerateToken(user types.UserDAO) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateToken", user)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateToken indicates an expected call of GenerateToken.
func (mr *MockAuthorizationMockRecorder) GenerateToken(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateToken", reflect.TypeOf((*MockAuthorization)(nil).GenerateToken), user)
}

// ParserToken mocks base method.
func (m *MockAuthorization) ParserToken(accessToken string) (*types.TokenClaims, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParserToken", accessToken)
	ret0, _ := ret[0].(*types.TokenClaims)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParserToken indicates an expected call of ParserToken.
func (mr *MockAuthorizationMockRecorder) ParserToken(accessToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParserToken", reflect.TypeOf((*MockAuthorization)(nil).ParserToken), accessToken)
}

// MockUser is a mock of User interface.
type MockUser struct {
	ctrl     *gomock.Controller
	recorder *MockUserMockRecorder
}

// MockUserMockRecorder is the mock recorder for MockUser.
type MockUserMockRecorder struct {
	mock *MockUser
}

// NewMockUser creates a new mock instance.
func NewMockUser(ctrl *gomock.Controller) *MockUser {
	mock := &MockUser{ctrl: ctrl}
	mock.recorder = &MockUserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUser) EXPECT() *MockUserMockRecorder {
	return m.recorder
}

// BuyItem mocks base method.
func (m *MockUser) BuyItem(userID int, itemName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuyItem", userID, itemName)
	ret0, _ := ret[0].(error)
	return ret0
}

// BuyItem indicates an expected call of BuyItem.
func (mr *MockUserMockRecorder) BuyItem(userID, itemName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuyItem", reflect.TypeOf((*MockUser)(nil).BuyItem), userID, itemName)
}

// GetUserInfo mocks base method.
func (m *MockUser) GetUserInfo(userID int) (types.UserInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserInfo", userID)
	ret0, _ := ret[0].(types.UserInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserInfo indicates an expected call of GetUserInfo.
func (mr *MockUserMockRecorder) GetUserInfo(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserInfo", reflect.TypeOf((*MockUser)(nil).GetUserInfo), userID)
}

// SendCoins mocks base method.
func (m *MockUser) SendCoins(senderID, receiverID, amount int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendCoins", senderID, receiverID, amount)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendCoins indicates an expected call of SendCoins.
func (mr *MockUserMockRecorder) SendCoins(senderID, receiverID, amount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendCoins", reflect.TypeOf((*MockUser)(nil).SendCoins), senderID, receiverID, amount)
}
