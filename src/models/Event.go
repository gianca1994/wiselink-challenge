package models

import (
	"gorm.io/gorm"
	"time"
)

type Event struct {
	gorm.Model
	Id        uint      `gorm:"primaryKey not null"`
	Title     string    `gorm:"not null unique"`
	ShortDesc string    `gorm:"not null"`
	LongDesc  string    `gorm:"not null"`
	Date      time.Time `gorm:"not null"`
	Time      time.Time `gorm:"not null"`
	Organizer string    `gorm:"not null"`
	Place     string    `gorm:"not null"`
	Status    string    `gorm:"not null"`
}

type EventResponse struct {
	Id        uint   `json:"id"`
	Title     string `json:"title"`
	ShortDesc string `json:"short_desc"`
	LongDesc  string `json:"long_desc"`
	Date      string `json:"date"`
	Time      string `json:"time"`
	Organizer string `json:"organizer"`
	Place     string `json:"place"`
	Status    string `json:"status"`
}

type EventCreate struct {
	Title     string `json:"title"`
	ShortDesc string `json:"short_desc"`
	LongDesc  string `json:"long_desc"`
	Date      string `json:"date"`
	Time      string `json:"time"`
	Organizer string `json:"organizer"`
	Place     string `json:"place"`
}

type EventUpdate struct {
	ShortDesc string `json:"short_desc"`
	LongDesc  string `json:"long_desc"`
	Date      string `json:"date"`
	Time      string `json:"time"`
	Organizer string `json:"organizer"`
	Place     string `json:"place"`
	Status    string `json:"status"`
}
