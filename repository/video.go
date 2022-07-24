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
	v.db.Order("id").Find(&videos)
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

func (v *Video) Excluir(id int) error {
	var video entity.Video
	v.db.First(&video, id)

	if video.Id == 0 {
		return entity.ErrVideoInexistente
	}

	v.db.Delete(&video, id)
	if v.db.Error != nil {
		return v.db.Error
	}

	return nil
}

func (v *Video) Atualizar(video *entity.Video) (*entity.Video, error) {
	var videoAtual entity.Video
	v.db.First(&videoAtual, video.Id)

	if videoAtual.Id == 0 {
		return nil, entity.ErrVideoInexistente
	}

	v.db.Model(&videoAtual).UpdateColumns(video)
	if v.db.Error != nil {
		return nil, v.db.Error
	}

	return &videoAtual, nil
}
