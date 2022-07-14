package mapper

import (
	"github.com/jmerschbacher/videos-api/domain"
	"github.com/jmerschbacher/videos-api/entity"
)

type Video struct {
}

func ToDomain(entity *entity.Video) *domain.Video {
	return &domain.Video{
		Id:        entity.Id,
		Titulo:    entity.Titulo,
		Descricao: entity.Descricao,
		URL:       entity.URL,
	}
}
