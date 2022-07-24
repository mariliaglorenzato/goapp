package interfaces

import "goapp/domain"

type IRepository interface {
	GetAllMovieMakers() ([]*domain.MovieMaker, error)
}
