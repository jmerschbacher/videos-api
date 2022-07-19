package categoria

import (
	"github.com/jmerschbacher/videos-api/domain"
	"github.com/jmerschbacher/videos-api/entity"
)

func ToDomain(entity entity.Categoria) domain.Categoria {
	return domain.Categoria{
		Id:     entity.Id,
		Titulo: entity.Titulo,
		Cor:    entity.Cor,
	}
}

func ToEntity(domain *domain.Categoria) entity.Categoria {
	return entity.Categoria{
		Id:     domain.Id,
		Titulo: domain.Titulo,
		Cor:    domain.Cor,
	}
}
