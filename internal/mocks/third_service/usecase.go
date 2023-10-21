// Code generated by MockGen. DO NOT EDIT.
// Source: internal/third_service/usecase.go

// Package mock_third_service is a generated GoMock package.
package mock_third_service

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/marcussss1/fio_service/internal/models"
)

// MockUsecase is a mock of Usecase interface.
type MockUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockUsecaseMockRecorder
}

// MockUsecaseMockRecorder is the mock recorder for MockUsecase.
type MockUsecaseMockRecorder struct {
	mock *MockUsecase
}

// NewMockUsecase creates a new mock instance.
func NewMockUsecase(ctrl *gomock.Controller) *MockUsecase {
	mock := &MockUsecase{ctrl: ctrl}
	mock.recorder = &MockUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUsecase) EXPECT() *MockUsecaseMockRecorder {
	return m.recorder
}

// GetFullPeopleInfo mocks base method.
func (m *MockUsecase) GetFullPeopleInfo(ctx context.Context, req models.AbbreviatedPeopleRequest) (models.People, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFullPeopleInfo", ctx, req)
	ret0, _ := ret[0].(models.People)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFullPeopleInfo indicates an expected call of GetFullPeopleInfo.
func (mr *MockUsecaseMockRecorder) GetFullPeopleInfo(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFullPeopleInfo", reflect.TypeOf((*MockUsecase)(nil).GetFullPeopleInfo), ctx, req)
}

// GetPeopleAge mocks base method.
func (m *MockUsecase) GetPeopleAge(ctx context.Context, name string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPeopleAge", ctx, name)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPeopleAge indicates an expected call of GetPeopleAge.
func (mr *MockUsecaseMockRecorder) GetPeopleAge(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPeopleAge", reflect.TypeOf((*MockUsecase)(nil).GetPeopleAge), ctx, name)
}

// GetPeopleGender mocks base method.
func (m *MockUsecase) GetPeopleGender(ctx context.Context, name string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPeopleGender", ctx, name)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPeopleGender indicates an expected call of GetPeopleGender.
func (mr *MockUsecaseMockRecorder) GetPeopleGender(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPeopleGender", reflect.TypeOf((*MockUsecase)(nil).GetPeopleGender), ctx, name)
}

// GetPeopleNationality mocks base method.
func (m *MockUsecase) GetPeopleNationality(ctx context.Context, name string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPeopleNationality", ctx, name)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPeopleNationality indicates an expected call of GetPeopleNationality.
func (mr *MockUsecaseMockRecorder) GetPeopleNationality(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPeopleNationality", reflect.TypeOf((*MockUsecase)(nil).GetPeopleNationality), ctx, name)
}
