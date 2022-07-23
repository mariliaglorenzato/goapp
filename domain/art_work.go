package domain

import "time"

type ArtWork struct {
	ID                  int64
	Name                string    `json:name`
	Description         string    `json:description`
	PublishedAt         time.Time `json:published_at`
	Plataforms          []*Plataform
	PlataformsMovieUrls []*string `json:plataform_movie_urls`
	MovieMakerId        int64
}
