package model

import (
	"time"
)

// User nil
type User struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
