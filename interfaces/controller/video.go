package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jmerschbacher/videos-api/domain"
	"github.com/jmerschbacher/videos-api/entity"
	"github.com/jmerschbacher/videos-api/usecase"
	"net/http"
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
			c.JSON(http.StatusNotFound, err)
		}
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, domain.Data{Data: videos})
}
