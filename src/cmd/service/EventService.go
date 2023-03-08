package service

import (
	"encoding/json"
	"time"

	_ "github.com/jackc/pgx/v5/pgtype"
	"net/http"

	"wiselink-challenge/src/internal/database"
	"wiselink-challenge/src/models"
)

func GetEventsService() []byte {
	db := database.DbConnection()
	var events []models.Event
	var eventsResponse []models.EventResponse
	db.Find(&events)
	for _, event := range events {
		eventsResponse = append(eventsResponse, models.EventResponse{
			Id:        event.Id,
			Title:     event.Title,
			ShortDesc: event.ShortDesc,
			LongDesc:  event.LongDesc,
			Date:      event.Date.Format("2006:01:02"),
			Time:      event.Time.Format("15:04"),
			Organizer: event.Organizer,
			Place:     event.Place,
			Status:    event.Status,
		})
	}
	data, _ := json.Marshal(eventsResponse)
	return data
}

func GetEvent(id string) ([]byte, error) {
	db := database.DbConnection()
	var event models.Event
	db.Where("id = ?", id).First(&event)
	if event.Id == 0 {
		return []byte("Event not found"), nil
	}
	data, _ := json.Marshal(models.EventResponse{
		Id:        event.Id,
		Title:     event.Title,
		ShortDesc: event.ShortDesc,
		LongDesc:  event.LongDesc,
		Date:      event.Date.Format("2006:01:02"),
		Time:      event.Time.Format("15:04"),
		Organizer: event.Organizer,
		Place:     event.Place,
		Status:    event.Status,
	})
	return data, nil
}

func CreateEventService(claims map[string]interface{}, r *http.Request) ([]byte, error) {
	db := database.DbConnection()
	var user models.User
	db.Where("username = ?", claims["username"]).First(&user)

	if user.Admin == false {
		return []byte("Unauthorized"), nil
	}

	var event models.EventCreate
	_ = json.NewDecoder(r.Body).Decode(&event)

	if event.Title == "" || event.ShortDesc == "" || event.LongDesc == "" ||
		event.Date == "" || event.Time == "" || event.Organizer == "" || event.Place == "" {
		return []byte("Invalid data"), nil
	}

	location, _ := time.LoadLocation("America/Argentina/Mendoza")
	dateEvent, err := time.ParseInLocation("2006-01-02", event.Date, location)
	if err != nil {
		return []byte("Invalid date"), nil
	}
	timeEvent, err := time.ParseInLocation("15:04", event.Time, location)
	if err != nil {
		return []byte("Invalid time"), nil
	}

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

func UpdateEventService(claims map[string]interface{}, param string, r *http.Request) ([]byte, error) {
	db := database.DbConnection()
	var user models.User
	db.Where("username = ?", claims["username"]).First(&user)

	if user.Admin == false {
		return []byte("Unauthorized"), nil
	}

	var event models.EventUpdate
	_ = json.NewDecoder(r.Body).Decode(&event)

	if event.ShortDesc == "" || event.LongDesc == "" ||
		event.Date == "" || event.Time == "" || event.Organizer == "" || event.Place == "" || event.Status == "" {
		return []byte("Invalid data"), nil
	}

	location, _ := time.LoadLocation("America/Argentina/Mendoza")
	dateEvent, err := time.ParseInLocation("2006-01-02", event.Date, location)
	if err != nil {
		return []byte("Invalid date"), nil
	}
	timeEvent, err := time.ParseInLocation("15:04", event.Time, location)
	if err != nil {
		return []byte("Invalid time"), nil
	}

	var eventDB models.Event
	db.Where("id = ?", param).First(&eventDB)
	if eventDB.Id == 0 {
		return []byte("Invalid event"), nil
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

func DeleteEventService(claims map[string]interface{}, idDeleted string) ([]byte, error) {
	db := database.DbConnection()
	var user models.User
	db.Where("username = ?", claims["username"]).First(&user)

	if user.Admin == false {
		return []byte("Unauthorized"), nil
	}

	var event models.Event
	db.Where("id = ?", idDeleted).First(&event)

	if event.Id == 0 {
		return []byte("Invalid event"), nil
	}

	db.Delete(&event)
	return []byte("Event deleted"), nil
}
