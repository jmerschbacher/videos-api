package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jmerschbacher/videos-api/interfaces/controller"
	"log"
)

func HandleRequests(video *controller.Video) {
	r := gin.Default()
	r.GET("/videos", video.ListarTodos)
	r.GET("/videos/:id", video.BuscarPorId)
	r.POST("/videos", video.Criar)
	r.DELETE("/videos/:id", video.Excluir)
	r.PATCH("/videos/:id", video.Editar)
	err := r.Run()
	if err != nil {
		log.Panic("Erro ao disponibilizar rotas")
	}
}
