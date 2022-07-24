package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jmerschbacher/videos-api/domain"
	"github.com/jmerschbacher/videos-api/entity"
	"github.com/jmerschbacher/videos-api/usecase"
	"net/http"
)

type Categoria struct {
	useCase *usecase.Categoria
}

func NewCategoriaController(categoriaUseCase *usecase.Categoria) *Categoria {
	return &Categoria{useCase: categoriaUseCase}
}

func (cat *Categoria) Criar(c *gin.Context) {
	var categoria *domain.Categoria
	err := c.ShouldBindJSON(&categoria)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Data{
			Data: domain.Error{
				Method:   http.MethodPost,
				Rota:     c.FullPath(),
				Codigo:   http.StatusBadRequest,
				Mensagem: err.Error(),
			},
		})
		return
	}

	if !categoria.IsValid() {
		c.JSON(http.StatusBadRequest, domain.Data{
			Data: domain.Error{
				Method:   http.MethodPost,
				Rota:     c.FullPath(),
				Codigo:   http.StatusBadRequest,
				Mensagem: entity.ErrParametroInvalido.Error(),
			},
		})
		return
	}

	err = cat.useCase.Criar(categoria)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Data{
			Data: domain.Error{
				Method:   http.MethodPost,
				Rota:     c.FullPath(),
				Codigo:   http.StatusBadRequest,
				Mensagem: err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusCreated, domain.Data{Data: &categoria})
}

func (cat *Categoria) ListarTodas(c *gin.Context) {
	categorais, err := cat.useCase.ListarTodas()
	if err != nil {
		if errors.Is(err, entity.ErrCategoriaNotFound) {
			c.JSON(http.StatusNotFound, domain.Data{
				Data: domain.Error{
					Method:   http.MethodGet,
					Rota:     c.FullPath(),
					Codigo:   http.StatusNotFound,
					Mensagem: err.Error(),
				},
			})
			return
		}
		c.JSON(http.StatusInternalServerError, domain.Data{
			Data: domain.Error{
				Method:   http.MethodGet,
				Rota:     c.FullPath(),
				Codigo:   http.StatusInternalServerError,
				Mensagem: err.Error(),
			},
		})
		return
	}
	c.JSON(http.StatusOK, domain.Data{Data: categorais})
}
