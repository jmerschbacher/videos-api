package usecase

import (
	"github.com/jmerschbacher/videos-api/domain"
	"github.com/jmerschbacher/videos-api/entity"
	"github.com/jmerschbacher/videos-api/mapper/categoria"
	"github.com/jmerschbacher/videos-api/repository"
)

type Categoria struct {
	repo *repository.Categoria
}

func NewCategoriaUseCase(repository *repository.Categoria) *Categoria {
	return &Categoria{repo: repository}
}

func (v *Categoria) ListarTodas() ([]domain.Categoria, error) {
	var categoriasRetornadas []entity.Categoria
	err := v.repo.ListarTodos(&categoriasRetornadas)

	if err != nil {
		return nil, err
	}

	if len(categoriasRetornadas) == 0 {
		return nil, entity.ErrCategoriaNotFound
	}

	var categoriasDomain []domain.Categoria
	for _, categoriaEntity := range categoriasRetornadas {
		categoriasDomain = append(categoriasDomain, categoria.ToDomain(categoriaEntity))
	}

	return categoriasDomain, nil
}

func (c *Categoria) Criar(categoriaDomain *domain.Categoria) error {
	categoriaEntity := categoria.ToEntity(categoriaDomain)
	err := c.repo.Criar(&categoriaEntity)
	if err != nil {
		return err
	}
	*categoriaDomain = categoria.ToDomain(categoriaEntity)
	return nil
}
