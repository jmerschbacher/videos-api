package main

import (
	"github.com/jmerschbacher/videos-api/infrastructure/database"
	"github.com/jmerschbacher/videos-api/interfaces/controller"
	"github.com/jmerschbacher/videos-api/interfaces/handler"
	"github.com/jmerschbacher/videos-api/repository"
	"github.com/jmerschbacher/videos-api/usecase"
)

func main() {
	db := database.PostgreSQLStarter()
	videoRepository := repository.NewVideoRepository(db)
	videoUsecase := usecase.NewVideoUseCase(videoRepository)
	videoController := controller.NewVideoController(videoUsecase)

	categoriaRepository := repository.NewCategoriaRepository(db)
	categoriaUsecase := usecase.NewCategoriaUseCase(categoriaRepository)
	categoriaController := controller.NewCategoriaController(categoriaUsecase)

	handler.HandleRequests(videoController, categoriaController)
}
