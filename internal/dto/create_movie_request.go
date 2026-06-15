package dto

type CreateMovieRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Genre       string `json:"genre"`
	Duration    int    `json:"duration"`
	PosterURL   string `json:"poster_url"`
}