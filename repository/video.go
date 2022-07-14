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

func (v *Video) BuscarPorId(id int) (*entity.Video, error) {
	var video *entity.Video
	v.db.First(&video, id)
	err := v.db.Error

	if err != nil {
		return nil, err
	}

	return video, nil
}

func (v *Video) Criar(video *entity.Video) (*entity.Video, error) {
	v.db.Create(&video)
	err := v.db.Error

	if err != nil {
		return nil, err
	}

	return video, nil
}
