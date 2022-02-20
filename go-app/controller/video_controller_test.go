package controller_test

import (
	"encoding/json"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go-app/controller"
	"go-app/mock"
	"go-app/model"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetVideos(t *testing.T) {
	t.Run("GET all videos", func(t *testing.T) {
		mockController := gomock.NewController(t)
		mockService := mock.NewMockIVideoService(mockController)

		mockService.EXPECT().Videos().Return(model.Videos{model.Video{ID: 2}}, nil).Times(1)

		videoController := controller.NewVideoController(mockService)

		request := httptest.NewRequest(http.MethodGet, "/api/v1/videos", http.NoBody)
		responseWriter := httptest.NewRecorder()

		videoController.ServeHTTP(responseWriter, request)

		expectedResponseBody := model.Videos{}
		err := json.Unmarshal(responseWriter.Body.Bytes(), &expectedResponseBody)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, responseWriter.Result().StatusCode)
		assert.Equal(t, "application/json; charset=UTF-8", responseWriter.Header().Get("content-type"))
		assert.Equal(t, expectedResponseBody[0].ID, 2)
	})

	t.Run("Service fail check", func(t *testing.T) {
		mockController := gomock.NewController(t)
		mockService := mock.NewMockIVideoService(mockController)

		mockService.EXPECT().Videos().Return(model.Videos{}, errors.New("error occurred")).Times(1)

		videoController := controller.NewVideoController(mockService)

		request := httptest.NewRequest(http.MethodGet, "/api/v1/videos", http.NoBody)
		responseWriter := httptest.NewRecorder()

		videoController.ServeHTTP(responseWriter, request)

		assert.Equal(t, http.StatusInternalServerError, responseWriter.Result().StatusCode)
	})

	t.Run("method not implemented", func(t *testing.T) {
		videoController := controller.NewVideoController(nil)

		request := httptest.NewRequest(http.MethodDelete, "/api/v1/videos", http.NoBody)
		responseWriter := httptest.NewRecorder()

		videoController.ServeHTTP(responseWriter, request)

		assert.Equal(t, http.StatusNotImplemented, responseWriter.Result().StatusCode)
	})
}
