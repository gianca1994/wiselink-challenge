package service

import (
	"encoding/json"
	"errors"
	"fmt"
	_ "github.com/jackc/pgx/v5/pgtype"
	"net/http"
	"strings"
	"time"
	"wiselink-challenge/src/cmd/repository"

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
	event, err := repository.GetEvent(id)
	if err != nil {
		return nil, err
	}
	adminRequired, _ := repository.CheckUserIsAdmin(claims)
	if !adminRequired && event.Status == "draft" {
		return []byte("Only admins can see posts in draft status."), nil
	}
	usersResponse, err := repository.GetUsersForEvent(&event)
	if err != nil {
		return nil, err
	}

	return json.Marshal(models.EventResponse{
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
	if err := validateDataAndTime(&event); err != nil {
		return []byte(err.Error()), nil
	}
	return repository.CreateEvent(event)
}

func UpdateEventService(claims map[string]interface{}, param string, r *http.Request) ([]byte, error) {
	adminRequired, _ := repository.CheckUserIsAdmin(claims)
	if adminRequired == false {
		return []byte("Unauthorized"), nil
	}

	var event models.EventUpdate
	_ = json.NewDecoder(r.Body).Decode(&event)
	if err := validateDataAndTime(&event); err != nil {
		return []byte(err.Error()), nil
	}
	return repository.UpdateEvent(param, event)
}

func DeleteEventService(claims map[string]interface{}, idDeleted string) ([]byte, error) {
	adminRequired, _ := repository.CheckUserIsAdmin(claims)
	if adminRequired == false {
		return []byte("Unauthorized"), nil
	}
	return repository.DeleteEvent(idDeleted)
}

/////////////////////// HELPERS ///////////////////////

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

func validateDataAndTime(event interface{}) error {
	var dateField, timeField string
	var err error

	switch e := event.(type) {
	case *models.EventCreate:
		dateField = e.Date
		timeField = e.Time
	case *models.EventUpdate:
		dateField = e.Date
		timeField = e.Time
	default:
		return fmt.Errorf("unexpected event type")
	}

	date, err := formatDate(dateField)
	if err != nil {
		return errors.New("invalid date format")
	}
	timeEvent, err := formatTime(timeField)
	if err != nil {
		return errors.New("invalid time format")
	}

	switch e := event.(type) {
	case *models.EventCreate:
		e.Date = date.Format("2006-01-02")
		e.Time = timeEvent.Format("15:04")
	case *models.EventUpdate:
		e.Date = date.Format("2006-01-02")
		e.Time = timeEvent.Format("15:04")
	}

	return nil
}
