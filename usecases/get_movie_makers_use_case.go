package usecases

import (
	"time"

	domain "goapp/domain"
	"goapp/usecases/interfaces"

	"github.com/jinzhu/gorm"
)

type GetMovieMakersUseCase struct {
	DB *gorm.DB
}

func NewGetMovieMakersUseCase(database *gorm.DB) interfaces.IGetMovieMakersUseCase {
	return &GetMovieMakersUseCase{DB: database}
}

func (usecase *GetMovieMakersUseCase) Perform() []domain.MovieMaker {
	var movieMakers = []domain.MovieMaker{
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
