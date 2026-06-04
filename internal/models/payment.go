package models

import (
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	ID                     uint           `gorm:"primaryKey"`
	BookingID              uint           `gorm:"not null"`

	MidtransOrderID        string
	MidtransTransactionID  string

	PaymentMethod          string
	Status                 string

	PaidAt                 *time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}