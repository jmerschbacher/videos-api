package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jmerschbacher/videos-api/domain"
	"github.com/jmerschbacher/videos-api/entity"
	"github.com/jmerschbacher/videos-api/usecase"
	"net/http"
	"strconv"
)

type Video struct {
	useCase *usecase.Video
}

func NewVideoController(videoUseCase *usecase.Video) *Video {
	return &Video{useCase: videoUseCase}
}

func (v *Video) ListarTodos(c *gin.Context) {
	videos, err := v.useCase.ListarTodos()
	if err != nil {
		if errors.Is(err, entity.ErrVideoNotFound) {
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
	c.JSON(http.StatusOK, domain.Data{Data: videos})
}

func (v *Video) BuscarPorId(c *gin.Context) {
	id := c.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Data{
			Data: domain.Error{
				Method:   http.MethodGet,
				Rota:     c.FullPath(),
				Codigo:   http.StatusBadRequest,
				Mensagem: entity.ErrParametroInvalido.Error(),
			},
		})
		return
	}

	var video *domain.Video
	video, err = v.useCase.BuscarPorId(intId)

	if err != nil {
		if errors.Is(err, entity.ErrVideoNotFound) {
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

	c.JSON(http.StatusOK, domain.Data{Data: video})
}

func (v *Video) Criar(c *gin.Context) {
	var video *domain.Video
	err := c.ShouldBindJSON(&video)
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

	if !video.IsValido() {
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

	video, err = v.useCase.Criar(video)
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

	c.JSON(http.StatusCreated, domain.Data{Data: &video})
}

func (v *Video) Excluir(c *gin.Context) {
	id := c.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Data{
			Data: domain.Error{
				Method:   http.MethodDelete,
				Rota:     c.FullPath(),
				Codigo:   http.StatusBadRequest,
				Mensagem: entity.ErrParametroInvalido.Error(),
			},
		})
		return
	}

	err = v.useCase.Excluir(intId)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Data{
			Data: domain.Error{
				Method:   http.MethodDelete,
				Rota:     c.FullPath(),
				Codigo:   http.StatusBadRequest,
				Mensagem: err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func (v *Video) Editar(c *gin.Context) {
	id := c.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Data{
			Data: domain.Error{
				Method:   http.MethodPatch,
				Rota:     c.FullPath(),
				Codigo:   http.StatusBadRequest,
				Mensagem: entity.ErrParametroInvalido.Error(),
			},
		})
		return
	}

	var video *domain.Video
	err = c.ShouldBindJSON(&video)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Data{
			Data: domain.Error{
				Method:   http.MethodPatch,
				Rota:     c.FullPath(),
				Codigo:   http.StatusBadRequest,
				Mensagem: err.Error(),
			},
		})
		return
	}

	if video.Id == 0 {
		video.Id = uint(intId)
	}

	if video.Id != uint(intId) {
		c.JSON(http.StatusBadRequest, domain.Data{
			Data: domain.Error{
				Method:   http.MethodPatch,
				Rota:     c.FullPath(),
				Codigo:   http.StatusBadRequest,
				Mensagem: entity.ErrIdPathEBodyDiferentes.Error(),
			},
		})
		return
	}

	if !video.IsValidoEdicao() {
		c.JSON(http.StatusBadRequest, domain.Data{
			Data: domain.Error{
				Method:   http.MethodPatch,
				Rota:     c.FullPath(),
				Codigo:   http.StatusBadRequest,
				Mensagem: entity.ErrParametroInvalido.Error(),
			},
		})
		return
	}

	var videoAtualizado *domain.Video
	videoAtualizado, err = v.useCase.EditarVideo(video)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Data{
			Data: domain.Error{
				Method:   http.MethodPatch,
				Rota:     c.FullPath(),
				Codigo:   http.StatusBadRequest,
				Mensagem: err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, domain.Data{Data: videoAtualizado})
}
