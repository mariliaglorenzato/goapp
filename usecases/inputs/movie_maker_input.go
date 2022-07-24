package inputs

type MovieMakerInput struct {
	FullName   string  `json:"name,binding:required"`
	ArtistName *string `json:"artist_name,omitempty"`
}
