package service_test

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go-app/mock"
	"go-app/model"
	"go-app/service"
	"testing"
)

func TestServiceShouldGetRepository(t *testing.T) {
	mockController := gomock.NewController(t)
	mockVideoRepository := mock.NewMockIVideoRepository(mockController)

	mockVideoRepository.EXPECT().GetAllVideos().Return(model.Videos{}, nil).Times(1)

	videoService := service.NewVideoService(mockVideoRepository)
	videos, err := videoService.Videos()
	assert.Nil(t, err)
	assert.Equal(t, model.Videos{}, videos)
}