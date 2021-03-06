// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/baetyl/baetyl-cloud/plugin (interfaces: Object)

// Package plugin is a generated GoMock package.
package plugin

import (
	models "github.com/baetyl/baetyl-cloud/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockObject is a mock of Object interface
type MockObject struct {
	ctrl     *gomock.Controller
	recorder *MockObjectMockRecorder
}

// MockObjectMockRecorder is the mock recorder for MockObject
type MockObjectMockRecorder struct {
	mock *MockObject
}

// NewMockObject creates a new mock instance
func NewMockObject(ctrl *gomock.Controller) *MockObject {
	mock := &MockObject{ctrl: ctrl}
	mock.recorder = &MockObjectMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockObject) EXPECT() *MockObjectMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockObject) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockObjectMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockObject)(nil).Close))
}

// CreateBucket mocks base method
func (m *MockObject) CreateBucket(arg0, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBucket", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateBucket indicates an expected call of CreateBucket
func (mr *MockObjectMockRecorder) CreateBucket(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBucket", reflect.TypeOf((*MockObject)(nil).CreateBucket), arg0, arg1, arg2)
}

// DeleteObject mocks base method
func (m *MockObject) DeleteObject(arg0, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteObject", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteObject indicates an expected call of DeleteObject
func (mr *MockObjectMockRecorder) DeleteObject(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteObject", reflect.TypeOf((*MockObject)(nil).DeleteObject), arg0, arg1, arg2)
}

// GenObjectURL mocks base method
func (m *MockObject) GenObjectURL(arg0, arg1, arg2 string) (*models.ObjectURL, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenObjectURL", arg0, arg1, arg2)
	ret0, _ := ret[0].(*models.ObjectURL)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenObjectURL indicates an expected call of GenObjectURL
func (mr *MockObjectMockRecorder) GenObjectURL(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenObjectURL", reflect.TypeOf((*MockObject)(nil).GenObjectURL), arg0, arg1, arg2)
}

// GetObject mocks base method
func (m *MockObject) GetObject(arg0, arg1, arg2 string) (*models.Object, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetObject", arg0, arg1, arg2)
	ret0, _ := ret[0].(*models.Object)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetObject indicates an expected call of GetObject
func (mr *MockObjectMockRecorder) GetObject(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObject", reflect.TypeOf((*MockObject)(nil).GetObject), arg0, arg1, arg2)
}

// HeadBucket mocks base method
func (m *MockObject) HeadBucket(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HeadBucket", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// HeadBucket indicates an expected call of HeadBucket
func (mr *MockObjectMockRecorder) HeadBucket(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HeadBucket", reflect.TypeOf((*MockObject)(nil).HeadBucket), arg0, arg1)
}

// HeadObject mocks base method
func (m *MockObject) HeadObject(arg0, arg1, arg2 string) (*models.ObjectMeta, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HeadObject", arg0, arg1, arg2)
	ret0, _ := ret[0].(*models.ObjectMeta)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HeadObject indicates an expected call of HeadObject
func (mr *MockObjectMockRecorder) HeadObject(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HeadObject", reflect.TypeOf((*MockObject)(nil).HeadObject), arg0, arg1, arg2)
}

// ListBucketObjects mocks base method
func (m *MockObject) ListBucketObjects(arg0, arg1 string, arg2 *models.ObjectParams) (*models.ListObjectsResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBucketObjects", arg0, arg1, arg2)
	ret0, _ := ret[0].(*models.ListObjectsResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBucketObjects indicates an expected call of ListBucketObjects
func (mr *MockObjectMockRecorder) ListBucketObjects(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBucketObjects", reflect.TypeOf((*MockObject)(nil).ListBucketObjects), arg0, arg1, arg2)
}

// ListBuckets mocks base method
func (m *MockObject) ListBuckets(arg0 string) ([]models.Bucket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBuckets", arg0)
	ret0, _ := ret[0].([]models.Bucket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBuckets indicates an expected call of ListBuckets
func (mr *MockObjectMockRecorder) ListBuckets(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBuckets", reflect.TypeOf((*MockObject)(nil).ListBuckets), arg0)
}

// PutObject mocks base method
func (m *MockObject) PutObject(arg0, arg1, arg2 string, arg3 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutObject", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// PutObject indicates an expected call of PutObject
func (mr *MockObjectMockRecorder) PutObject(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutObject", reflect.TypeOf((*MockObject)(nil).PutObject), arg0, arg1, arg2, arg3)
}

// PutObjectFromURL mocks base method
func (m *MockObject) PutObjectFromURL(arg0, arg1, arg2, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutObjectFromURL", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// PutObjectFromURL indicates an expected call of PutObjectFromURL
func (mr *MockObjectMockRecorder) PutObjectFromURL(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutObjectFromURL", reflect.TypeOf((*MockObject)(nil).PutObjectFromURL), arg0, arg1, arg2, arg3)
}
