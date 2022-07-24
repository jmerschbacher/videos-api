package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jmerschbacher/videos-api/interfaces/controller"
	"github.com/jmerschbacher/videos-api/interfaces/handler/categoria"
	"github.com/jmerschbacher/videos-api/interfaces/handler/video"
	"log"
)

func HandleRequests(videoController *controller.Video, categoriaController *controller.Categoria) {
	r := gin.Default()
	video.Endpoints(videoController, r)
	categoria.Endpoints(categoriaController, r)
	err := r.Run()
	if err != nil {
		log.Panic("Erro ao disponibilizar rotas")
	}
}
