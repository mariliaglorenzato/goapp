package persistence

import (
	"goapp/domain"
	"goapp/usecases/interfaces"

	"github.com/jinzhu/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) interfaces.IRepository {
	return &Repository{db: db}
}

func (r *Repository) GetAllMovieMakers() ([]*domain.MovieMaker, error) {
	movieMakers := []*domain.MovieMaker{}
	result := r.db.Find(&movieMakers)
	if result.Error != nil {
		return nil, result.Error
	}

	return movieMakers, nil
}

func (r *Repository) CreateMovieMaker(movieMaker *domain.MovieMaker) (*domain.MovieMaker, error) {
	result := r.db.Create(&movieMaker)
	if result.Error != nil {
		return nil, result.Error
	}

	return movieMaker, nil
}
