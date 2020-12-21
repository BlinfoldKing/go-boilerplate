// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package users is a generated GoMock package.
package users

import (
	gomock "github.com/golang/mock/gomock"
	entity "go-boilerplate/entity"
	reflect "reflect"
)

// MockRepository is a mock of Repository interface
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// Save mocks base method
func (m *MockRepository) Save(arg0 entity.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save
func (mr *MockRepositoryMockRecorder) Save(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockRepository)(nil).Save), arg0)
}

// FindByEmail mocks base method
func (m *MockRepository) FindByEmail(email string) (entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByEmail", email)
	ret0, _ := ret[0].(entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByEmail indicates an expected call of FindByEmail
func (mr *MockRepositoryMockRecorder) FindByEmail(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByEmail", reflect.TypeOf((*MockRepository)(nil).FindByEmail), email)
}

// Update mocks base method
func (m *MockRepository) Update(id string, changeset entity.UserChangeSet) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", id, changeset)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockRepositoryMockRecorder) Update(id, changeset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockRepository)(nil).Update), id, changeset)
}

// GetList mocks base method
func (m *MockRepository) GetList(limit, offset int) ([]entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetList", limit, offset)
	ret0, _ := ret[0].([]entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetList indicates an expected call of GetList
func (mr *MockRepositoryMockRecorder) GetList(limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetList", reflect.TypeOf((*MockRepository)(nil).GetList), limit, offset)
}

// FindByID mocks base method
func (m *MockRepository) FindByID(id string) (entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", id)
	ret0, _ := ret[0].(entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID
func (mr *MockRepositoryMockRecorder) FindByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockRepository)(nil).FindByID), id)
}

// DeleteByID mocks base method
func (m *MockRepository) DeleteByID(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByID", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByID indicates an expected call of DeleteByID
func (mr *MockRepositoryMockRecorder) DeleteByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockRepository)(nil).DeleteByID), id)
}
