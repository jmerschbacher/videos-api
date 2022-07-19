package usecase

import (
	"github.com/jmerschbacher/videos-api/domain"
	"github.com/jmerschbacher/videos-api/entity"
	"github.com/jmerschbacher/videos-api/mapper/video"
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
		return nil, entity.ErrVideoNotFound
	}

	var videos []*domain.Video
	for _, videoEncontrado := range listaVideos {
		videos = append(videos, video.ToDomain(videoEncontrado))
	}

	return videos, nil
}

func (v *Video) BuscarPorId(id int) (*domain.Video, error) {
	videoEncontrado, err := v.repo.BuscarPorId(id)

	if err != nil {
		return nil, err
	}

	if videoEncontrado.Id == 0 {
		return nil, entity.ErrVideoNotFound
	}

	return video.ToDomain(videoEncontrado), nil
}

func (v *Video) Criar(videoDomain *domain.Video) (*domain.Video, error) {
	videoEntity := video.ToEntity(videoDomain)
	videoGravado, err := v.repo.Criar(videoEntity)

	if err != nil {
		return nil, err
	}

	return video.ToDomain(videoGravado), nil
}

func (v *Video) Excluir(id int) error {
	err := v.repo.Excluir(id)
	if err != nil {
		return err
	}
	return nil
}

func (v *Video) EditarVideo(videoDomain *domain.Video) (*domain.Video, error) {
	videoEntity := video.ToEntity(videoDomain)
	entityAtualizada, err := v.repo.Atualizar(videoEntity)
	if err != nil {
		return nil, err
	}
	return video.ToDomain(entityAtualizada), nil
}
