package interfaces

import "goapp/domain"

type IGetMovieMakersUseCase interface {
	Perform() []domain.MovieMaker
}
