package repository

import (
	"time"
	"wiselink-challenge/src/internal/database"
	"wiselink-challenge/src/models"
)

func CreateEvent(event models.EventCreate, dateEvent, timeEvent time.Time) error {
	db := database.PostgreSQL()
	db.Create(&models.Event{
		Title:     event.Title,
		ShortDesc: event.ShortDesc,
		LongDesc:  event.LongDesc,
		Date:      dateEvent,
		Time:      timeEvent,
		Organizer: event.Organizer,
		Place:     event.Place,
		Status:    "draft",
	})
	return nil
}
