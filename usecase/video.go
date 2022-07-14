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

func (s *Video) ListarTodos() ([]*domain.Video, error) {
	listaVideos, err := s.repo.ListarTodos()

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
