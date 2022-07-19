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

func (v *Categoria) ListarTodos() ([]*entity.Categoria, error) {
	var categoria []*entity.Categoria
	v.db.Order("id").Find(&categoria)
	err := v.db.Error

	if err != nil {
		return nil, err
	}

	return categoria, nil
}
