package usecase

import (
	"github.com/jmerschbacher/videos-api/domain"
	"github.com/jmerschbacher/videos-api/mapper/categoria"
	"github.com/jmerschbacher/videos-api/repository"
)

type Categoria struct {
	repo *repository.Categoria
}

func NewCategoriaUseCase(repository *repository.Categoria) *Categoria {
	return &Categoria{repo: repository}
}

//func (v *Categoria) ListarTodas() ([]*domain.Categoria, error) {
//	listaCategorias, err := v.repo.ListarTodos()
//
//	if err != nil {
//		return nil, err
//	}
//
//	if len(listaCategorias) == 0 {
//		return nil, entity.ErrCategoriaNotFound
//	}
//
//	var categorias []*domain.Categoria
//	for _, categoria := range listaCategorias {
//		categorias = append(categorias, mapper.ToDomain(categoria))
//	}
//
//	return videos, nil
//}

func (c *Categoria) Criar(categoriaDomain *domain.Categoria) error {
	categoriaEntity := categoria.ToEntity(categoriaDomain)
	err := c.repo.Criar(&categoriaEntity)
	if err != nil {
		return err
	}
	*categoriaDomain = categoria.ToDomain(categoriaEntity)
	return nil
}
