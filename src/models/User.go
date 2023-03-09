package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id       uint    `gorm:"primaryKey not null"`
	Username string  `gorm:"uniqueIndex not null"`
	Email    string  `gorm:"uniqueIndex not null"`
	Password string  `gorm:"not null"`
	Admin    bool    `gorm:"not null"`
	Events   []Event `gorm:"many2many:user_events"`
}

type UserProfileResponse struct {
	Username string          `json:"username"`
	Email    string          `json:"email"`
	Admin    bool            `json:"admin"`
	Events   []EventResponse `json:"events"`
}

type UserRegisterDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UserLoginDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
