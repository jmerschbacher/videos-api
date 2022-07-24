package repository

import (
	"github.com/jmerschbacher/videos-api/entity"
	"gorm.io/gorm"
)

type Categoria struct {
	db *gorm.DB
}

func NewCategoriaRepository(db *gorm.DB) *Categoria {
	return &Categoria{db: db}
}

func (c *Categoria) ListarTodos(categorias *[]entity.Categoria) error {
	c.db.Order("id").Find(&categorias)
	err := c.db.Error

	if err != nil {
		return err
	}

	return nil
}

func (c *Categoria) Criar(categoria *entity.Categoria) error {
	resultado := c.db.First(&categoria, &categoria.Id)
	if resultado.RowsAffected != 0 {
		return entity.ErrCategoriaJaExiste
	}

	err := c.db.Create(&categoria)
	if err != nil {
		return err.Error
	}

	return nil
}
