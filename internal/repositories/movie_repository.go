package repositories

import (
	"ticket-booking/internal/models"

	"gorm.io/gorm"
)

type MovieRepository struct {
	db *gorm.DB
}

func NewMovieRepository(db *gorm.DB) *MovieRepository {
	return &MovieRepository{
		db: db,
	}
}

func (r *MovieRepository) Create(movie *models.Movie) error {
	return r.db.Create(movie).Error
}

func (r *MovieRepository) FindAll() ([]models.Movie, error) {

	var movies []models.Movie

	err := r.db.Find(&movies).Error

	return movies, err
}

func (r *MovieRepository) FindByID(id uint) (*models.Movie, error) {

	var movie models.Movie

	err := r.db.First(&movie, id).Error

	if err != nil {
		return nil, err
	}

	return &movie, nil
}

func (r *MovieRepository) Update(movie *models.Movie) error {
	return r.db.Save(movie).Error
}

func (r *MovieRepository) Delete(id uint) error {
	// return r.db.Delete(&models.Movie{}, id).Error
	result := r.db.Delete(&models.Movie{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

