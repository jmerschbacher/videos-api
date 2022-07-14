package mapper

import (
	"github.com/jmerschbacher/videos-api/domain"
	"github.com/jmerschbacher/videos-api/entity"
)

func ToDomain(entity *entity.Video) *domain.Video {
	return &domain.Video{
		Id:        entity.Id,
		Titulo:    entity.Titulo,
		Descricao: entity.Descricao,
		URL:       entity.URL,
	}
}

func ToEntity(domain *domain.Video) *entity.Video {
	return &entity.Video{
		Id:        domain.Id,
		Titulo:    domain.Titulo,
		Descricao: domain.Descricao,
		URL:       domain.URL,
	}
}
