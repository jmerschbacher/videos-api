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
		if errors.Is(err, entity.ErrNotFound) {
			c.JSON(http.StatusNotFound, domain.Data{
				Data: domain.Error{
					Rota:     c.FullPath(),
					Codigo:   http.StatusNotFound,
					Mensagem: err.Error(),
				},
			})
			return
		}
		c.JSON(http.StatusInternalServerError, domain.Data{
			Data: domain.Error{
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
		if errors.Is(err, entity.ErrNotFound) {
			c.JSON(http.StatusNotFound, domain.Data{
				Data: domain.Error{
					Rota:     c.FullPath(),
					Codigo:   http.StatusNotFound,
					Mensagem: err.Error(),
				},
			})
			return
		}

		c.JSON(http.StatusInternalServerError, domain.Data{
			Data: domain.Error{
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
				Rota:     c.FullPath(),
				Codigo:   http.StatusBadRequest,
				Mensagem: err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusCreated, domain.Data{Data: &video})
}
