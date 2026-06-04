package models

import "time"

type BookingSeat struct {
	ID        uint `gorm:"primaryKey"`

	BookingID uint
	ScheduleID uint
	StudioSeatID uint

	CreatedAt time.Time
}