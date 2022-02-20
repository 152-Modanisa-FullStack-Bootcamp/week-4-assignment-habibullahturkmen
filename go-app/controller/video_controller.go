package controller

import (
	"encoding/json"
	"go-app/service"
	"net/http"
)

type IVideoController interface {
	GetVideos(w http.ResponseWriter)
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type VideoController struct {
	videoService service.IVideoService
}

func (v *VideoController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		v.GetVideos(w)
	} else {
		w.WriteHeader(http.StatusNotImplemented)
	}
}

func (v *VideoController) GetVideos(w http.ResponseWriter) {
	videos, err := v.videoService.Videos()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	videoBytes, err := json.Marshal(videos)

	w.Header().Add("content-type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(videoBytes)
}

func NewVideoController(service service.IVideoService) IVideoController {
	return &VideoController{videoService: service}
}
