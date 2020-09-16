package model

import (
	"gorm.io/gorm"
)

// Article nil
type Article struct {
	ID          uint64 `gorm:"primaryKey"`
	Title       string
	Description string
	Content     string
	UserID      uint
	gorm.Model
}
