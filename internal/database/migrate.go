package database

import (
	"ticket-booking/internal/models"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.User{},
		&models.Movie{},
		&models.Studio{},
		&models.StudioSeat{},
		&models.Schedule{},
		&models.Booking{},
		&models.BookingSeat{},
		&models.Payment{},
	)
}