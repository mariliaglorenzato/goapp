package usecases

import (
	domain "goapp/domain"
	inputs "goapp/usecases/inputs"
	"goapp/usecases/interfaces"
)

type CreateMovieMakerUseCase struct {
	Repository interfaces.IRepository
}

func NewCreateMovieMakerUseCase(repository interfaces.IRepository) interfaces.ICreateMovieMakerUseCase {
	return &CreateMovieMakerUseCase{Repository: repository}
}

func (usecase *CreateMovieMakerUseCase) Perform(movieMakerInput *inputs.MovieMakerInput) (*domain.MovieMaker, error) {
	movieMaker := domain.MovieMaker{
		FullName:   movieMakerInput.FullName,
		ArtistName: movieMakerInput.ArtistName,
	}

	_, err := usecase.Repository.CreateMovieMaker(&movieMaker)
	if err != nil {
		return nil, err
	}

	return &movieMaker, nil
}
