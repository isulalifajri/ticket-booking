package services

import (
	"ticket-booking/internal/dto"
	"ticket-booking/internal/models"
	"ticket-booking/internal/repositories"
)

type MovieService struct {
	movieRepo *repositories.MovieRepository
}

func NewMovieService(
	movieRepo *repositories.MovieRepository,
) *MovieService {
	return &MovieService{
		movieRepo: movieRepo,
	}
}

func (s *MovieService) Create(
	req dto.CreateMovieRequest,
) error {

	movie := models.Movie{
		Title:       req.Title,
		Description: req.Description,
		Genre:       req.Genre,
		Duration:    req.Duration,
		PosterURL:   req.PosterURL,
	}

	return s.movieRepo.Create(&movie)
}

func (s *MovieService) GetAll() ([]models.Movie, error) {
	return s.movieRepo.FindAll()
}

func (s *MovieService) GetByID(id uint) (*models.Movie, error) {
	return s.movieRepo.FindByID(id)
}

func (s *MovieService) Update(
	id uint,
	req dto.UpdateMovieRequest,
) error {

	movie, err := s.movieRepo.FindByID(id)

	if err != nil {
		return err
	}

	movie.Title = req.Title
	movie.Description = req.Description
	movie.Genre = req.Genre
	movie.Duration = req.Duration
	movie.PosterURL = req.PosterURL

	return s.movieRepo.Update(movie)
}

func (s *MovieService) Delete(id uint) error {
	return s.movieRepo.Delete(id)
}