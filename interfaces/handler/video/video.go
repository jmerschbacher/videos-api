package video

import (
	"github.com/gin-gonic/gin"
	"github.com/jmerschbacher/videos-api/interfaces/controller"
)

func Endpoints(video *controller.Video, r *gin.Engine) {
	r.GET("/videos", video.ListarTodos)
	r.GET("/videos/:id", video.BuscarPorId)
	r.POST("/videos", video.Criar)
	r.DELETE("/videos/:id", video.Excluir)
	r.PATCH("/videos/:id", video.Editar)
}
