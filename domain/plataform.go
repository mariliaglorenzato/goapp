package domain

type Plataform struct {
	ID         int64
	Name       string `json:name`
	Url        string `json:url`
	OpenSource bool   `json:open_sourced`
}
