package main

import (
	"fmt"
	"os"

	controllers "goapp/controllers"
	persistence "goapp/persistence"
	usecases "goapp/usecases"

	"github.com/gin-gonic/gin"
)

func main() {
	// load database
	GormImpl := persistence.NewGormImpl()

	// load repository
	repository := persistence.NewRepository(GormImpl.DB)

	// load use cases
	getMovieMakersUseCase := usecases.NewGetMovieMakersUseCase(repository)

	// load controllers
	movieMakersController := controllers.NewMovieMakersController(getMovieMakersUseCase)

	// load routes
	router := gin.Default()
	router.GET("/moviemakers", movieMakersController.GetAll)

	router.Run(fmt.Sprintf(":%v", os.Getenv("PORT")))
}
