package service

import (
	"encoding/json"
	"time"
	"wiselink-challenge/src/internal/database"
	"wiselink-challenge/src/models"
)

func GetProfileService(claims map[string]interface{}) []byte {
	db, _ := database.PostgreSQL()
	var user models.User
	var userResponse models.UserProfileResponse
	var eventsResponse []models.EventResponseProfileUser

	db.Where("username = ?", claims["username"]).First(&user)

	userResponse.Username = user.Username
	userResponse.Email = user.Email
	userResponse.Admin = user.Admin

	_ = db.Model(&user).Association("Events").Find(&user.Events)

	for _, event := range user.Events {
		eventsResponse = append(eventsResponse, models.EventResponseProfileUser{
			Id:        event.Id,
			Title:     event.Title,
			ShortDesc: event.ShortDesc,
			LongDesc:  event.LongDesc,
			Date:      event.Date.Format("2006-01-02"),
			Time:      event.Time.Format("15:04"),
			Organizer: event.Organizer,
			Place:     event.Place,
			Status:    event.Status,
		})
	}
	userResponse.Events = eventsResponse
	data, _ := json.Marshal(userResponse)
	return data
}

func GetRegisteredEvents(claims map[string]interface{}, filter string) []byte {
	db, _ := database.PostgreSQL()
	var user models.User
	var eventsResponse []models.EventResponseProfileUser

	db.Where("username = ?", claims["username"]).First(&user)

	_ = db.Model(&user).Association("Events").Find(&user.Events)

	for _, event := range user.Events {
		if filter == "completed" {
			eventTime := time.Date(event.Date.Year(), event.Date.Month(), event.Date.Day(), event.Time.Hour(), event.Time.Minute(), 0, 0, time.UTC)
			now := time.Now().UTC()
			if eventTime.Before(now) {
				eventsResponse = append(eventsResponse, models.EventResponseProfileUser{
					Id:        event.Id,
					Title:     event.Title,
					ShortDesc: event.ShortDesc,
					LongDesc:  event.LongDesc,
					Date:      event.Date.Format("2006-01-02"),
					Time:      event.Time.Format("15:04"),
					Organizer: event.Organizer,
					Place:     event.Place,
					Status:    event.Status,
				})
			}
		} else if filter == "active" {
			eventTime := time.Date(event.Date.Year(), event.Date.Month(), event.Date.Day(), event.Time.Hour(), event.Time.Minute(), 0, 0, time.UTC)
			now := time.Now().UTC()
			if eventTime.After(now) {
				eventsResponse = append(eventsResponse, models.EventResponseProfileUser{
					Id:        event.Id,
					Title:     event.Title,
					ShortDesc: event.ShortDesc,
					LongDesc:  event.LongDesc,
					Date:      event.Date.Format("2006-01-02"),
					Time:      event.Time.Format("15:04"),
					Organizer: event.Organizer,
					Place:     event.Place,
					Status:    event.Status,
				})
			}
		} else {
			eventsResponse = append(eventsResponse, models.EventResponseProfileUser{
				Id:        event.Id,
				Title:     event.Title,
				ShortDesc: event.ShortDesc,
				LongDesc:  event.LongDesc,
				Date:      event.Date.Format("2006-01-02"),
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

func RegisterToEvent(claims map[string]interface{}, event_id string) []byte {
	db, _ := database.PostgreSQL()
	var user models.User
	var event models.Event
	db.Where("username = ?", claims["username"]).First(&user)
	db.Where("id = ?", event_id).First(&event)

	eventTime := time.Date(event.Date.Year(), event.Date.Month(), event.Date.Day(), event.Time.Hour(), event.Time.Minute(), 0, 0, time.UTC)
	now := time.Now().UTC()
	if eventTime.Before(now) {
		return []byte("Event has already started")
	}

	err := db.Model(&user).Association("Events").Append(&event)
	if err != nil {
		return nil
	}
	return []byte("Registered to event")
}
