// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/peterstirrup/messages/internal/messages/usecases (interfaces: WhatsApp)

// Package mock_usecases is a generated GoMock package.
package mock_usecases

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	entities "github.com/peterstirrup/messages/internal/messages/entities"
	reflect "reflect"
)

// MockWhatsApp is a mock of WhatsApp interface
type MockWhatsApp struct {
	ctrl     *gomock.Controller
	recorder *MockWhatsAppMockRecorder
}

// MockWhatsAppMockRecorder is the mock recorder for MockWhatsApp
type MockWhatsAppMockRecorder struct {
	mock *MockWhatsApp
}

// NewMockWhatsApp creates a new mock instance
func NewMockWhatsApp(ctrl *gomock.Controller) *MockWhatsApp {
	mock := &MockWhatsApp{ctrl: ctrl}
	mock.recorder = &MockWhatsAppMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockWhatsApp) EXPECT() *MockWhatsAppMockRecorder {
	return m.recorder
}

// GetContacts mocks base method
func (m *MockWhatsApp) GetContacts(arg0 context.Context, arg1 int64) ([]entities.Contact, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContacts", arg0, arg1)
	ret0, _ := ret[0].([]entities.Contact)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContacts indicates an expected call of GetContacts
func (mr *MockWhatsAppMockRecorder) GetContacts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContacts", reflect.TypeOf((*MockWhatsApp)(nil).GetContacts), arg0, arg1)
}
