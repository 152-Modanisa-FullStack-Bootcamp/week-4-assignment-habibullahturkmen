package repository

import (
	"github.com/stretchr/testify/assert"
	"os"
	"path"
	"runtime"
	"testing"
)

func TestVideoExist(t *testing.T) {
	videos, err := NewVideoRepository().GetAllVideos()
	assert.Nil(t, err)
	assert.Equal(t, len(videos.Videos), 12)

}

// fixes relative path problems while testing
func init() {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "..")
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}
}