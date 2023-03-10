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
	Events   []Event `gorm:"many2many:event_users"`
}

type UserProfileResponse struct {
	Username string                     `json:"username"`
	Email    string                     `json:"email"`
	Admin    bool                       `json:"admin"`
	Events   []EventResponseProfileUser `json:"events"`
}

type UserEventResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
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
