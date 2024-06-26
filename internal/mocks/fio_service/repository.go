// Code generated by MockGen. DO NOT EDIT.
// Source: internal/fio_service/repository.go

// Package mock_fio_service is a generated GoMock package.
package mock_fio_service

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/marcussss1/fio_service/internal/models"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// CreatePeople mocks base method.
func (m *MockRepository) CreatePeople(ctx context.Context, people models.People) (models.People, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePeople", ctx, people)
	ret0, _ := ret[0].(models.People)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePeople indicates an expected call of CreatePeople.
func (mr *MockRepositoryMockRecorder) CreatePeople(ctx, people interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePeople", reflect.TypeOf((*MockRepository)(nil).CreatePeople), ctx, people)
}

// DeletePeopleByID mocks base method.
func (m *MockRepository) DeletePeopleByID(ctx context.Context, peopleID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePeopleByID", ctx, peopleID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePeopleByID indicates an expected call of DeletePeopleByID.
func (mr *MockRepositoryMockRecorder) DeletePeopleByID(ctx, peopleID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePeopleByID", reflect.TypeOf((*MockRepository)(nil).DeletePeopleByID), ctx, peopleID)
}

// GetPeople mocks base method.
func (m *MockRepository) GetPeople(ctx context.Context, req models.GetPeopleRequest) ([]models.People, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPeople", ctx, req)
	ret0, _ := ret[0].([]models.People)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPeople indicates an expected call of GetPeople.
func (mr *MockRepositoryMockRecorder) GetPeople(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPeople", reflect.TypeOf((*MockRepository)(nil).GetPeople), ctx, req)
}

// GetPeopleByID mocks base method.
func (m *MockRepository) GetPeopleByID(ctx context.Context, peopleID string) (models.People, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPeopleByID", ctx, peopleID)
	ret0, _ := ret[0].(models.People)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPeopleByID indicates an expected call of GetPeopleByID.
func (mr *MockRepositoryMockRecorder) GetPeopleByID(ctx, peopleID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPeopleByID", reflect.TypeOf((*MockRepository)(nil).GetPeopleByID), ctx, peopleID)
}

// UpdatePeopleByID mocks base method.
func (m *MockRepository) UpdatePeopleByID(ctx context.Context, req models.People, peopleID string) (models.People, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePeopleByID", ctx, req, peopleID)
	ret0, _ := ret[0].(models.People)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdatePeopleByID indicates an expected call of UpdatePeopleByID.
func (mr *MockRepositoryMockRecorder) UpdatePeopleByID(ctx, req, peopleID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePeopleByID", reflect.TypeOf((*MockRepository)(nil).UpdatePeopleByID), ctx, req, peopleID)
}
