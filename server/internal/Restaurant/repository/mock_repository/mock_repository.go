// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	reflect "reflect"
	dto "server/internal/domain/dto"
	entity "server/internal/domain/entity"

	gomock "github.com/golang/mock/gomock"
)

// MockRestaurantRepositoryI is a mock of RestaurantRepositoryI interface.
type MockRestaurantRepositoryI struct {
	ctrl     *gomock.Controller
	recorder *MockRestaurantRepositoryIMockRecorder
}

// MockRestaurantRepositoryIMockRecorder is the mock recorder for MockRestaurantRepositoryI.
type MockRestaurantRepositoryIMockRecorder struct {
	mock *MockRestaurantRepositoryI
}

// NewMockRestaurantRepositoryI creates a new mock instance.
func NewMockRestaurantRepositoryI(ctrl *gomock.Controller) *MockRestaurantRepositoryI {
	mock := &MockRestaurantRepositoryI{ctrl: ctrl}
	mock.recorder = &MockRestaurantRepositoryIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRestaurantRepositoryI) EXPECT() *MockRestaurantRepositoryIMockRecorder {
	return m.recorder
}

// GetCategories mocks base method.
func (m *MockRestaurantRepositoryI) GetCategories() ([]*entity.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCategories")
	ret0, _ := ret[0].([]*entity.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCategories indicates an expected call of GetCategories.
func (mr *MockRestaurantRepositoryIMockRecorder) GetCategories() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCategories", reflect.TypeOf((*MockRestaurantRepositoryI)(nil).GetCategories))
}

// GetCategoriesByRestaurantID mocks base method.
func (m *MockRestaurantRepositoryI) GetCategoriesByRestaurantID(id uint) ([]*entity.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCategoriesByRestaurantID", id)
	ret0, _ := ret[0].([]*entity.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCategoriesByRestaurantID indicates an expected call of GetCategoriesByRestaurantID.
func (mr *MockRestaurantRepositoryIMockRecorder) GetCategoriesByRestaurantID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCategoriesByRestaurantID", reflect.TypeOf((*MockRestaurantRepositoryI)(nil).GetCategoriesByRestaurantID), id)
}

// GetMenuTypesByRestaurantID mocks base method.
func (m *MockRestaurantRepositoryI) GetMenuTypesByRestaurantID(id uint) ([]*entity.MenuType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMenuTypesByRestaurantID", id)
	ret0, _ := ret[0].([]*entity.MenuType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMenuTypesByRestaurantID indicates an expected call of GetMenuTypesByRestaurantID.
func (mr *MockRestaurantRepositoryIMockRecorder) GetMenuTypesByRestaurantID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMenuTypesByRestaurantID", reflect.TypeOf((*MockRestaurantRepositoryI)(nil).GetMenuTypesByRestaurantID), id)
}

// GetRestaurantByID mocks base method.
func (m *MockRestaurantRepositoryI) GetRestaurantByID(id uint) (*entity.Restaurant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRestaurantByID", id)
	ret0, _ := ret[0].(*entity.Restaurant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRestaurantByID indicates an expected call of GetRestaurantByID.
func (mr *MockRestaurantRepositoryIMockRecorder) GetRestaurantByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRestaurantByID", reflect.TypeOf((*MockRestaurantRepositoryI)(nil).GetRestaurantByID), id)
}

// GetRestaurantByName mocks base method.
func (m *MockRestaurantRepositoryI) GetRestaurantByName(name string) (*entity.Restaurant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRestaurantByName", name)
	ret0, _ := ret[0].(*entity.Restaurant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRestaurantByName indicates an expected call of GetRestaurantByName.
func (mr *MockRestaurantRepositoryIMockRecorder) GetRestaurantByName(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRestaurantByName", reflect.TypeOf((*MockRestaurantRepositoryI)(nil).GetRestaurantByName), name)
}

// GetRestaurants mocks base method.
func (m *MockRestaurantRepositoryI) GetRestaurants() ([]*entity.Restaurant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRestaurants")
	ret0, _ := ret[0].([]*entity.Restaurant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRestaurants indicates an expected call of GetRestaurants.
func (mr *MockRestaurantRepositoryIMockRecorder) GetRestaurants() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRestaurants", reflect.TypeOf((*MockRestaurantRepositoryI)(nil).GetRestaurants))
}

// GetRestaurantsByCategory mocks base method.
func (m *MockRestaurantRepositoryI) GetRestaurantsByCategory(name string) ([]*entity.Restaurant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRestaurantsByCategory", name)
	ret0, _ := ret[0].([]*entity.Restaurant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRestaurantsByCategory indicates an expected call of GetRestaurantsByCategory.
func (mr *MockRestaurantRepositoryIMockRecorder) GetRestaurantsByCategory(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRestaurantsByCategory", reflect.TypeOf((*MockRestaurantRepositoryI)(nil).GetRestaurantsByCategory), name)
}

// SearchCategories mocks base method.
func (m *MockRestaurantRepositoryI) SearchCategories(word string) ([]*entity.Restaurant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchCategories", word)
	ret0, _ := ret[0].([]*entity.Restaurant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchCategories indicates an expected call of SearchCategories.
func (mr *MockRestaurantRepositoryIMockRecorder) SearchCategories(word interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchCategories", reflect.TypeOf((*MockRestaurantRepositoryI)(nil).SearchCategories), word)
}

// SearchRestaurants mocks base method.
func (m *MockRestaurantRepositoryI) SearchRestaurants(word string) ([]*entity.Restaurant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchRestaurants", word)
	ret0, _ := ret[0].([]*entity.Restaurant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchRestaurants indicates an expected call of SearchRestaurants.
func (mr *MockRestaurantRepositoryIMockRecorder) SearchRestaurants(word interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchRestaurants", reflect.TypeOf((*MockRestaurantRepositoryI)(nil).SearchRestaurants), word)
}

// UpdateComments mocks base method.
func (m *MockRestaurantRepositoryI) UpdateComments(comment *dto.ReqCreateComment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateComments", comment)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateComments indicates an expected call of UpdateComments.
func (mr *MockRestaurantRepositoryIMockRecorder) UpdateComments(comment interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateComments", reflect.TypeOf((*MockRestaurantRepositoryI)(nil).UpdateComments), comment)
}
