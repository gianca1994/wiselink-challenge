package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id       uint   `gorm:"primaryKey not null"`
	Username string `gorm:"uniqueIndex not null"`
	Email    string `gorm:"uniqueIndex not null"`
	Password string `gorm:"not null"`
}
