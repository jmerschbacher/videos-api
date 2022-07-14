package repository

import (
	"github.com/jmerschbacher/videos-api/entity"
	"gorm.io/gorm"
)

type Video struct {
	db *gorm.DB
}

func NewVideoRepository(db *gorm.DB) *Video {
	return &Video{db: db}
}

func (v *Video) ListarTodos() ([]*entity.Video, error) {
	var videos []*entity.Video
	v.db.Find(&videos)
	err := v.db.Error
	if err != nil {
		return nil, err
	}
	return videos, nil
}
