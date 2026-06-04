package models

import (
	"time"

	"gorm.io/gorm"
)

type Booking struct {
	ID          uint           `gorm:"primaryKey"`
	UserID      uint           `gorm:"not null"`

	BookingCode string         `gorm:"size:50;uniqueIndex"`
	TotalPrice int64

	Status      string         `gorm:"size:20;default:pending"`
	ExpiredAt   *time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// pending
// paid
// expired
// cancelled