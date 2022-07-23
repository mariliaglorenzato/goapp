package controllers

import (
	"goapp/usecases/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MovieMakersController struct {
	getMovieMakersUseCase interfaces.IGetMovieMakersUseCase
}

func NewMovieMakersController(getMovieMakersUseCase interfaces.IGetMovieMakersUseCase) *MovieMakersController {
	return &MovieMakersController{
		getMovieMakersUseCase: getMovieMakersUseCase,
	}
}
func (c *MovieMakersController) GetAll(ctx *gin.Context) {
	ucOutput := c.getMovieMakersUseCase.Perform()
	ctx.IndentedJSON(http.StatusOK, ucOutput)
}
