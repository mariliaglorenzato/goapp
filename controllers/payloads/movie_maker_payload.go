package payloads

type MovieMakerPayload struct {
	FullName   string  `json:"name,binding:required"`
	ArtistName *string `json:"artist_name,omitempty"`
}
