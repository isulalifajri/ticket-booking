package models

import (
	"time"

	"gorm.io/gorm"
)

type StudioSeat struct {
	ID         uint           `gorm:"primaryKey"`
	StudioID   uint           `gorm:"not null"`
	SeatNumber string         `gorm:"size:10;not null"`

	Studio     Studio

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}