package server

import (
	"fmt"
	"go-app/controller"
	"go-app/repository"
	"go-app/service"
	"net/http"
)

type Server struct {
}

func (s *Server) StartServer(port int) error {
	videoRepository := repository.NewVideoRepository()
	videoService := service.NewVideoService(videoRepository)
	videoController := controller.NewVideoController(videoService)

	http.HandleFunc("/api/v1/videos", videoController.ServeHTTP)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	return err
}

func NewServer() *Server {
	return &Server{}
}
