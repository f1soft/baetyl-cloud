// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/baetyl/baetyl-cloud/plugin/default/pki (interfaces: PVC)

// Package plugin is a generated GoMock package.
package plugin

import (
	models "github.com/baetyl/baetyl-cloud/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockPVC is a mock of PVC interface
type MockPVC struct {
	ctrl     *gomock.Controller
	recorder *MockPVCMockRecorder
}

// MockPVCMockRecorder is the mock recorder for MockPVC
type MockPVCMockRecorder struct {
	mock *MockPVC
}

// NewMockPVC creates a new mock instance
func NewMockPVC(ctrl *gomock.Controller) *MockPVC {
	mock := &MockPVC{ctrl: ctrl}
	mock.recorder = &MockPVCMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPVC) EXPECT() *MockPVCMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockPVC) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockPVCMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockPVC)(nil).Close))
}

// CountCertByParentId mocks base method
func (m *MockPVC) CountCertByParentId(arg0 string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountCertByParentId", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountCertByParentId indicates an expected call of CountCertByParentId
func (mr *MockPVCMockRecorder) CountCertByParentId(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountCertByParentId", reflect.TypeOf((*MockPVC)(nil).CountCertByParentId), arg0)
}

// CreateCert mocks base method
func (m *MockPVC) CreateCert(arg0 models.Cert) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCert", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateCert indicates an expected call of CreateCert
func (mr *MockPVCMockRecorder) CreateCert(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCert", reflect.TypeOf((*MockPVC)(nil).CreateCert), arg0)
}

// DeleteCert mocks base method
func (m *MockPVC) DeleteCert(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCert", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCert indicates an expected call of DeleteCert
func (mr *MockPVCMockRecorder) DeleteCert(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCert", reflect.TypeOf((*MockPVC)(nil).DeleteCert), arg0)
}

// GetCert mocks base method
func (m *MockPVC) GetCert(arg0 string) (*models.Cert, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCert", arg0)
	ret0, _ := ret[0].(*models.Cert)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCert indicates an expected call of GetCert
func (mr *MockPVCMockRecorder) GetCert(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCert", reflect.TypeOf((*MockPVC)(nil).GetCert), arg0)
}

// UpdateCert mocks base method
func (m *MockPVC) UpdateCert(arg0 models.Cert) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCert", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCert indicates an expected call of UpdateCert
func (mr *MockPVCMockRecorder) UpdateCert(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCert", reflect.TypeOf((*MockPVC)(nil).UpdateCert), arg0)
}
