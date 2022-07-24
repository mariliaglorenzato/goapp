package persistence

import (
	"fmt"
	"goapp/domain"

	"github.com/jinzhu/gorm"

	_ "github.com/mattn/go-sqlite3"
)

type GormImpl struct {
	DB *gorm.DB
}

func NewGormImpl() *GormImpl {
	database, err := gorm.Open("sqlite3", "test.db")

	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %v", err))
	}

	//todo: this should not be here, move as soon as possible
	database.AutoMigrate(&domain.MovieMaker{})

	return &GormImpl{DB: database}
}
