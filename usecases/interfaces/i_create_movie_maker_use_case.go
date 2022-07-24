package interfaces

import (
	"goapp/domain"
	"goapp/usecases/inputs"
)

type ICreateMovieMakerUseCase interface {
	Perform(movieMakerInput *inputs.MovieMakerInput) (*domain.MovieMaker, error)
}
