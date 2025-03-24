package models

import (
	//	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       int    `gorm:"primaryKey"`
	Email    string `gorm:"uniqueIndex;not null"`
	Username string `gorm:"not null"`
	Password string `gorm:"not null"`
}
