package repository

import (
	"time"
	"wiselink-challenge/src/internal/database"
	"wiselink-challenge/src/models"
)

func CheckUserIsAdmin(claims map[string]interface{}) bool {
	db := database.PostgreSQL()
	var user models.User
	db.Where("username = ?", claims["username"]).First(&user)
	return user.Admin
}

func GetEvents() []models.Event {
	db := database.PostgreSQL()
	var events []models.Event
	db.Find(&events)
	return events
}

func GetEvent(id string) models.Event {
	db := database.PostgreSQL()
	var event models.Event
	db.Where("id = ?", id).First(&event)
	return event
}

func CreateEvent(event models.EventCreate, dateEvent, timeEvent time.Time) ([]byte, error) {
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
	return []byte("Event created"), nil
}

func UpdateEvent(param string, event models.EventUpdate, dateEvent, timeEvent time.Time) ([]byte, error) {
	var eventDB models.Event
	db := database.PostgreSQL()
	db.Where("id = ?", param).First(&eventDB)
	if eventDB.Id == 0 {
		return []byte("Event not found"), nil
	}
	db.Model(&eventDB).Updates(models.Event{
		Title:     eventDB.Title,
		ShortDesc: event.ShortDesc,
		LongDesc:  event.LongDesc,
		Date:      dateEvent,
		Time:      timeEvent,
		Organizer: event.Organizer,
		Place:     event.Place,
		Status:    event.Status,
	})
	return []byte("Event updated"), nil
}

func DeleteEvent(idDeleted string) ([]byte, error) {
	var event models.Event
	db := database.PostgreSQL()
	db.Where("id = ?", idDeleted).First(&event)
	if event.Id == 0 {
		return []byte("Invalid event"), nil
	}
	db.Delete(&event)
	return []byte("Event deleted"), nil
}