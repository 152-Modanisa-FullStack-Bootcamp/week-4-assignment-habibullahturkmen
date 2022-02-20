package service

import (
	"go-app/model"
	"go-app/repository"
)

type IVideoService interface {
	Videos() (model.Videos, error)
}

type VideoService struct {
	videoRepository repository.IVideoRepository
}

func (v *VideoService) Videos() (model.Videos, error) {
	return v.videoRepository.GetAllVideos()
}

func NewVideoService(repository repository.IVideoRepository) IVideoService {
	return &VideoService{videoRepository: repository}
}