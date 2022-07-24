package controllers

import (
	"net/http"

	"goapp/controllers/payloads"
	"goapp/usecases/inputs"
	"goapp/usecases/interfaces"

	"github.com/gin-gonic/gin"
)

type MovieMakersController struct {
	getMovieMakersUseCase   interfaces.IGetMovieMakersUseCase
	createMovieMakerUseCase interfaces.ICreateMovieMakerUseCase
}

func NewMovieMakersController(
	getMovieMakersUseCase interfaces.IGetMovieMakersUseCase,
	createMovieMakerUseCase interfaces.ICreateMovieMakerUseCase,
) *MovieMakersController {
	return &MovieMakersController{
		getMovieMakersUseCase:   getMovieMakersUseCase,
		createMovieMakerUseCase: createMovieMakerUseCase,
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

func (c *MovieMakersController) Create(ctx *gin.Context) {
	var payload *payloads.MovieMakerPayload
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input := inputs.MovieMakerInput{
		FullName:   payload.FullName,
		ArtistName: payload.ArtistName,
	}
	ucOutput, err := c.createMovieMakerUseCase.Perform(&input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"data": ucOutput})
}
