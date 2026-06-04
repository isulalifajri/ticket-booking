package models

import (
	"time"

	"gorm.io/gorm"
)

type Schedule struct {
	ID        uint           `gorm:"primaryKey"`
	MovieID   uint           `gorm:"not null"`
	StudioID  uint           `gorm:"not null"`

	Movie     Movie          `gorm:"foreignKey:MovieID"`
	Studio    Studio         `gorm:"foreignKey:StudioID"`

	ShowTime  time.Time      `gorm:"not null"`
	Price     int64          `gorm:"not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}