package main

import (
	controllers "goapp/controllers"
	"goapp/usecases"

	"github.com/gin-gonic/gin"
)

func main() {
	//load use cases
	getMovieMakersUseCase := usecases.NewGetMovieMakersUseCase()

	//load controllers
	movieMakersController := controllers.NewMovieMakersController(getMovieMakersUseCase)

	//load routes
	router := gin.Default()
	router.GET("/moviemakers", movieMakersController.GetAll)

	router.Run("localhost:8080")
}
