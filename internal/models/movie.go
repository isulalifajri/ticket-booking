package models

import (
	"time"

	"gorm.io/gorm"
)

type Movie struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Title       string         `gorm:"size:255;not null" json:"title"`
	Description string         `gorm:"type:text" json:"description"`
	Duration    int            `json:"duration"`
	PosterURL   string         `gorm:"type:text" json:"poster_url"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}