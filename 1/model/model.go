package model

// album represents data about a record album.
type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

type Website struct {
	Url     string `json:"url"`
	Name    string `json:"name"`
	Details string `json:"details"`
}
