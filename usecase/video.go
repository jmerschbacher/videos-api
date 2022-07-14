package usecase

import (
	"github.com/jmerschbacher/videos-api/domain"
	"github.com/jmerschbacher/videos-api/entity"
	"github.com/jmerschbacher/videos-api/mapper"
	"github.com/jmerschbacher/videos-api/repository"
)

type Video struct {
	repo *repository.Video
}

func NewVideoUseCase(repository *repository.Video) *Video {
	return &Video{repo: repository}
}

func (v *Video) ListarTodos() ([]*domain.Video, error) {
	listaVideos, err := v.repo.ListarTodos()

	if err != nil {
		return nil, err
	}

	if len(listaVideos) == 0 {
		return nil, entity.ErrNotFound
	}

	var videos []*domain.Video
	for _, video := range listaVideos {
		videos = append(videos, mapper.ToDomain(video))
	}

	return videos, nil
}

func (v *Video) BuscarPorId(id int) (*domain.Video, error) {
	video, err := v.repo.BuscarPorId(id)

	if err != nil {
		return nil, err
	}

	if video.Id == 0 {
		return nil, entity.ErrNotFound
	}

	return mapper.ToDomain(video), nil
}

func (v *Video) Criar(videoDomain *domain.Video) (*domain.Video, error) {
	videoEntity := mapper.ToEntity(videoDomain)
	videoGravado, err := v.repo.Criar(videoEntity)

	if err != nil {
		return nil, err
	}

	return mapper.ToDomain(videoGravado), nil
}
