package domain

type MovieMaker struct {
	ID         int64   `json:"id" gorm:"primary_key"`
	FullName   string  `json:"name"`
	ArtistName *string `json:"artist_name"`
	// BirthDate  time.Time `json: birth_date`
	// DeathDate  time.Time `json:death_date`
	ArtWorks []ArtWork `json:"art_works"`
}
