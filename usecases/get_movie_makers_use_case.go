package usecases

import (
	domain "goapp/domain"
	"goapp/usecases/interfaces"
)

type GetMovieMakersUseCase struct {
	Repository interfaces.IRepository
}

func NewGetMovieMakersUseCase(repository interfaces.IRepository) interfaces.IGetMovieMakersUseCase {
	return &GetMovieMakersUseCase{Repository: repository}
}

func (usecase *GetMovieMakersUseCase) Perform() ([]*domain.MovieMaker, error) {
	output, err := usecase.Repository.GetAllMovieMakers()

	if err != nil {
		return nil, err
	}

	return output, nil
	// 	{
	// 		ID: 1, FullName: "John Coltrane",
	// 		// ArtWorks: []domain.ArtWork{
	// 		// 	{ID: 1, Name: "Blue Train", Description: "Test.....", PublishedAt: time.Now()},
	// 		// },
	// 	},
	// 	{
	// 		ID: 2, FullName: "Gerry Mulligan",
	// 		// ArtWorks: []domain.ArtWork{
	// 		// 	{ID: 2, Name: "Jeru", Description: "Test.....", PublishedAt: time.Now()},
	// 		// },
	// 	},
	// 	{
	// 		ID: 3, FullName: "Sarah Vaughan",
	// 		// ArtWorks: []domain.ArtWork{
	// 		// 	{ID: 3, Name: "Sarah Vaughan and Clifford Brown", Description: "Test...", PublishedAt: time.Now()},
	// 		// },
	// 	},
	// }
}
