package domain

type MovieMakers struct {
	ID         int64   `json: id`
	FullName   string  `json: name`
	ArtistName *string `json: artist_name`
	// BirthDate  time.Time `json: birth_date`
	// DeathDate  time.Time `json:death_date`
	ArtWorks []ArtWork `json:art_works`
}
