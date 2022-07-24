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
	ucOutput, err := c.getMovieMakersUseCase.Perform()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, ucOutput)
	return
}
