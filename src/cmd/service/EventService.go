package service

import (
	"encoding/json"
	"time"
	"wiselink-challenge/src/cmd/repository"

	_ "github.com/jackc/pgx/v5/pgtype"
	"net/http"

	"wiselink-challenge/src/internal/database"
	"wiselink-challenge/src/models"
)

func checkUserIsAdmin(claims map[string]interface{}) bool {
	db := database.PostgreSQL()
	var user models.User
	db.Where("username = ?", claims["username"]).First(&user)
	return user.Admin
}

func GetEventsService(claims map[string]interface{}) []byte {
	db := database.PostgreSQL()
	var events []models.Event
	db.Find(&events)

	adminRequired := checkUserIsAdmin(claims)
	var eventsResponse []models.EventResponse

	for _, event := range events {
		if adminRequired || event.Status != "draft" {
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
	}

	data, _ := json.Marshal(eventsResponse)
	return data
}

func GetEvent(claims map[string]interface{}, id string) ([]byte, error) {
	db := database.PostgreSQL()
	var event models.Event
	db.Where("id = ?", id).First(&event)
	if event.Id == 0 {
		return []byte("Event not found"), nil
	}

	adminRequired := checkUserIsAdmin(claims)
	if adminRequired == false && event.Status == "draft" {
		return []byte("Only admins can see posts in draft status."), nil
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
	adminRequired := checkUserIsAdmin(claims)
	if adminRequired == false {
		return []byte("Unauthorized"), nil
	}

	var event models.EventCreate
	_ = json.NewDecoder(r.Body).Decode(&event)

	if event.Title == "" || event.ShortDesc == "" || event.LongDesc == "" ||
		event.Date == "" || event.Time == "" || event.Organizer == "" || event.Place == "" {
		return []byte("Invalid data"), nil
	}

	location, _ := time.LoadLocation("America/Argentina/Mendoza")
	dateEvent, _ := time.ParseInLocation("2006-01-02", event.Date, location)
	timeEvent, _ := time.ParseInLocation("15:04", event.Time, location)

	err := repository.CreateEvent(event, dateEvent, timeEvent)
	if err != nil {
		return nil, err
	}
	return []byte("Event created"), nil
}

func UpdateEventService(claims map[string]interface{}, param string, r *http.Request) ([]byte, error) {
	adminRequired := checkUserIsAdmin(claims)
	if adminRequired == false {
		return []byte("Unauthorized"), nil
	}

	var event models.EventUpdate
	_ = json.NewDecoder(r.Body).Decode(&event)

	if event.ShortDesc == "" || event.LongDesc == "" ||
		event.Date == "" || event.Time == "" || event.Organizer == "" || event.Place == "" || event.Status == "" {
		return []byte("Invalid data"), nil
	}

	location, _ := time.LoadLocation("America/Argentina/Mendoza")
	dateEvent, _ := time.ParseInLocation("2006-01-02", event.Date, location)
	timeEvent, _ := time.ParseInLocation("15:04", event.Time, location)

	var eventDB models.Event
	db := database.PostgreSQL()
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
	adminRequired := checkUserIsAdmin(claims)
	if adminRequired == false {
		return []byte("Unauthorized"), nil
	}

	var event models.Event
	db := database.PostgreSQL()
	db.Where("id = ?", idDeleted).First(&event)
	if event.Id == 0 {
		return []byte("Invalid event"), nil
	}

	db.Delete(&event)
	return []byte("Event deleted"), nil
}
