package categoria

import (
	"github.com/gin-gonic/gin"
	"github.com/jmerschbacher/videos-api/interfaces/controller"
)

func Endpoints(categoria *controller.Categoria, r *gin.Engine) {
	r.POST("/categorias", categoria.Criar)
	r.GET("/categorias", categoria.ListarTodas)
}
