package service

import (
	"encoding/json"
	"strings"
	"time"
	"wiselink-challenge/src/cmd/repository"
	"wiselink-challenge/src/internal/database"

	_ "github.com/jackc/pgx/v5/pgtype"
	"net/http"

	"wiselink-challenge/src/models"
)

func GetEventsService(claims map[string]interface{}, filterSelected string, filter string) []byte {
	var eventsResponse []models.EventResponseProfileUser

	events, _ := repository.GetEvents()
	adminRequired, _ := repository.CheckUserIsAdmin(claims)

	applyFilter := func(filter string) func(event models.Event) bool {
		switch filterSelected {
		case "date":
			return func(event models.Event) bool {
				return event.Date.Format("2006-01-02") == filter
			}
		case "status":
			return func(event models.Event) bool {
				return event.Status == filter
			}
		case "title":
			return func(event models.Event) bool {
				return strings.Contains(strings.ToLower(event.Title), strings.ToLower(filter))
			}
		default:
			return func(event models.Event) bool {
				return true
			}
		}
	}

	filterFunc := applyFilter(filter)

	for _, event := range events {
		if adminRequired || event.Status != "draft" {
			if filterFunc(event) {
				eventsResponse = append(eventsResponse, models.EventResponseProfileUser{
					Id:        event.Id,
					Title:     event.Title,
					ShortDesc: event.ShortDesc,
					Date:      event.Date.Format("2006-01-02"),
					Time:      event.Time.Format("15:04"),
					Place:     event.Place,
					Status:    event.Status,
				})
			}
		}
	}

	data, _ := json.Marshal(eventsResponse)
	if len(eventsResponse) == 0 {
		return []byte("[]")
	}
	return data
}

func GetEvent(claims map[string]interface{}, id string) ([]byte, error) {
	event, _ := repository.GetEvent(id)
	if event.Id == 0 {
		return []byte("Event not found"), nil
	}

	adminRequired, _ := repository.CheckUserIsAdmin(claims)
	if adminRequired == false && event.Status == "draft" {
		return []byte("Only admins can see posts in draft status."), nil
	}

	db, _ := database.PostgreSQL()
	_ = db.Model(&event).Association("Users").Find(&event.Users)

	sqlDB, _ := db.DB()
	_ = sqlDB.Close()

	var usersResponse []models.UserEventResponse
	for _, user := range event.Users {
		usersResponse = append(usersResponse, models.UserEventResponse{
			Username: user.Username,
			Email:    user.Email,
		})
	}

	data, _ := json.Marshal(models.EventResponse{
		Id:        event.Id,
		Title:     event.Title,
		ShortDesc: event.ShortDesc,
		LongDesc:  event.LongDesc,
		Date:      event.Date.Format("2006-01-02"),
		Time:      event.Time.Format("15:04"),
		Organizer: event.Organizer,
		Place:     event.Place,
		Status:    event.Status,
		Users:     usersResponse,
	})
	return data, nil
}

func CreateEventService(claims map[string]interface{}, r *http.Request) ([]byte, error) {
	adminRequired, _ := repository.CheckUserIsAdmin(claims)
	if adminRequired == false {
		return []byte("Unauthorized"), nil
	}

	var event models.EventCreate
	_ = json.NewDecoder(r.Body).Decode(&event)

	if event.Title == "" || event.ShortDesc == "" || event.LongDesc == "" ||
		event.Date == "" || event.Time == "" || event.Organizer == "" || event.Place == "" {
		return []byte("Invalid data"), nil
	}

	date, err := formatDate(event.Date)
	if err != nil {
		return []byte("Invalid date format"), err
	}
	timeEvent, err := formatTime(event.Time)
	if err != nil {
		return []byte("Invalid time format"), err
	}

	return repository.CreateEvent(event, date, timeEvent)
}

func UpdateEventService(claims map[string]interface{}, param string, r *http.Request) ([]byte, error) {
	adminRequired, _ := repository.CheckUserIsAdmin(claims)
	if adminRequired == false {
		return []byte("Unauthorized"), nil
	}

	var event models.EventUpdate
	_ = json.NewDecoder(r.Body).Decode(&event)

	if event.ShortDesc == "" || event.LongDesc == "" ||
		event.Date == "" || event.Time == "" || event.Organizer == "" ||
		event.Place == "" || event.Status == "" {
		return []byte("Invalid data"), nil
	}

	date, err := formatDate(event.Date)
	if err != nil {
		return []byte("Invalid date format"), err
	}

	timeEvent, err := formatTime(event.Time)
	if err != nil {
		return []byte("Invalid time format"), err
	}

	return repository.UpdateEvent(param, event, date, timeEvent)
}

func DeleteEventService(claims map[string]interface{}, idDeleted string) ([]byte, error) {
	adminRequired, _ := repository.CheckUserIsAdmin(claims)
	if adminRequired == false {
		return []byte("Unauthorized"), nil
	}
	return repository.DeleteEvent(idDeleted)
}

func formatDate(date string) (time.Time, error) {
	dateFormat := "2006-01-02"
	dateEvent, err := time.Parse(dateFormat, date)
	return dateEvent, err
}

func formatTime(timeDate string) (time.Time, error) {
	timeFormat := "15:04"
	timeEvent, err := time.Parse(timeFormat, timeDate)
	return timeEvent, err
}
