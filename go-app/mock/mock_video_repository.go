// Code generated by MockGen. DO NOT EDIT.
// Source: repository/video_repository.go

// Package mock is a generated GoMock package.
package mock

import (
	model "go-app/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIVideoRepository is a mock of IVideoRepository interface.
type MockIVideoRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIVideoRepositoryMockRecorder
}

// MockIVideoRepositoryMockRecorder is the mock recorder for MockIVideoRepository.
type MockIVideoRepositoryMockRecorder struct {
	mock *MockIVideoRepository
}

// NewMockIVideoRepository creates a new mock instance.
func NewMockIVideoRepository(ctrl *gomock.Controller) *MockIVideoRepository {
	mock := &MockIVideoRepository{ctrl: ctrl}
	mock.recorder = &MockIVideoRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIVideoRepository) EXPECT() *MockIVideoRepositoryMockRecorder {
	return m.recorder
}

// GetAllVideos mocks base method.
func (m *MockIVideoRepository) GetAllVideos() (model.Videos, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllVideos")
	ret0, _ := ret[0].(model.Videos)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllVideos indicates an expected call of GetAllVideos.
func (mr *MockIVideoRepositoryMockRecorder) GetAllVideos() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllVideos", reflect.TypeOf((*MockIVideoRepository)(nil).GetAllVideos))
}