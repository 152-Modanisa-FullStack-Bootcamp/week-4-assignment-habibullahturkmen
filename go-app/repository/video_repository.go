package repository

import (
	"encoding/json"
	"go-app/model"
	"io/ioutil"
	"os"
)

type IVideoRepository interface {
	GetAllVideos() (model.Videos, error)
}

type VideoRepository struct {
}

func (v *VideoRepository) GetAllVideos() (model.Videos, error) {
	jsonFile, err := os.Open("./.config/videos.json")
	defer jsonFile.Close()

	bytes, _ := ioutil.ReadAll(jsonFile)

	var videos model.Videos

	json.Unmarshal(bytes, &videos)

	return videos, err
}

func NewVideoRepository() IVideoRepository {
	return &VideoRepository{}
}