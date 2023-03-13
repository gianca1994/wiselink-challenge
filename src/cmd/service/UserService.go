package service

import (
	"encoding/json"
	"time"
	"wiselink-challenge/src/cmd/repository"
	"wiselink-challenge/src/models"
)

func GetProfileService(claims map[string]interface{}) ([]byte, error) {
	var user models.User
	user, _ = repository.GetUserByUsername(claims["username"].(string))

	return json.Marshal(models.UserProfileResponse{
		Username: user.Username,
		Email:    user.Email,
		Admin:    user.Admin,
	})
}

func GetRegisteredEvents(claims map[string]interface{}, filter string) ([]byte, error) {
	var user models.User

	userID, err := claims["user_id"].(float64)
	if err != true {
		return []byte("Error"), nil
	}
	user.Events, _ = repository.GetRegisteredEvents(uint(userID))
	return json.Marshal(filterEvents(user.Events, filter))
}

func RegisterToEvent(claims map[string]interface{}, event_id string) []byte {
	var user models.User
	var event models.Event

	user, err := repository.GetUserByUsername(claims["username"].(string))
	if err != nil {
		return []byte("User not found")
	}
	event, err = repository.GetEvent(event_id)
	if err != nil {
		return []byte("Event not found")
	}
	if event.Status != "active" {
		return []byte("Event is not active")
	}
	if event.Date.Before(time.Now().UTC()) {
		return []byte("Event has already started")
	}

	err = repository.RegisterUserInEvent(user, event)
	if err != nil {
		return nil
	}
	return []byte("Registered to event")
}

//////////////////////// HELPER FUNCTIONS ////////////////////////

func filterEvents(events []models.Event, filter string) []models.EventResponseProfileUser {
	filteredEvents := make([]models.EventResponseProfileUser, 0)

	for _, event := range events {
		switch filter {
		case "completed":
			if event.Date.Before(time.Now().UTC()) && event.Status == "active" {
				filteredEvents = append(filteredEvents, createEventResponse(event))
			}
		case "active":
			if event.Date.After(time.Now().UTC()) && event.Status == "active" {
				filteredEvents = append(filteredEvents, createEventResponse(event))
			}
		default:
			filteredEvents = append(filteredEvents, createEventResponse(event))
		}
	}
	return filteredEvents
}

func createEventResponse(event models.Event) models.EventResponseProfileUser {
	return models.EventResponseProfileUser{
		Id:        event.Id,
		Title:     event.Title,
		ShortDesc: event.ShortDesc,
		Date:      event.Date.Format("2006-01-02"),
		Time:      event.Time.Format("15:04"),
		Place:     event.Place,
		Status:    event.Status,
	}
}
