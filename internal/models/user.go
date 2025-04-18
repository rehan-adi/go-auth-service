package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"uniqueIndex;not null"`
	Username string `gorm:"not null"`
	Password string `gorm:"not null"`
}
