package usecases

import (
	"time"

	domain "goapp/domain"
	"goapp/usecases/interfaces"
)

type GetMovieMakersUseCase struct{}

func NewGetMovieMakersUseCase() interfaces.IGetMovieMakersUseCase {
	return &GetMovieMakersUseCase{}
}

func (usecase *GetMovieMakersUseCase) Perform() []domain.MovieMakers {
	var movieMakers = []domain.MovieMakers{
		{
			ID: 1, FullName: "John Coltrane", ArtWorks: []domain.ArtWork{
				{ID: 1, Name: "Blue Train", Description: "Test.....", PublishedAt: time.Now()},
			},
		},
		{
			ID: 2, FullName: "Gerry Mulligan", ArtWorks: []domain.ArtWork{
				{ID: 2, Name: "Jeru", Description: "Test.....", PublishedAt: time.Now()},
			},
		},
		{
			ID: 3, FullName: "Sarah Vaughan", ArtWorks: []domain.ArtWork{
				{ID: 3, Name: "Sarah Vaughan and Clifford Brown", Description: "Test...", PublishedAt: time.Now()},
			},
		},
	}
	return movieMakers
}
