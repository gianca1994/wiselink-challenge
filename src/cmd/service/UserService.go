package service

import (
	"encoding/json"
	"time"
	"wiselink-challenge/src/cmd/repository"
	"wiselink-challenge/src/models"
)

func GetProfileService(claims map[string]interface{}) []byte {
	var user models.User
	var userResponse models.UserProfileResponse

	user, _ = repository.GetUserByUsername(claims["username"].(string))
	userResponse.Username = user.Username
	userResponse.Email = user.Email
	userResponse.Admin = user.Admin

	data, _ := json.Marshal(userResponse)
	return data
}

func GetRegisteredEvents(claims map[string]interface{}, filter string) []byte {
	var user models.User
	var eventsResponse []models.EventResponseProfileUser

	userIDInterface, _ := claims["user_id"]
	userID, _ := userIDInterface.(float64)
	user.Events, _ = repository.GetRegisteredEvents(uint(userID))

	for _, event := range user.Events {
		if filter == "completed" {
			if eventIsCompleted(event) {
				eventsResponse = append(eventsResponse, createEventResponse(event))
			}
		} else if filter == "active" {
			if eventIsActive(event) {
				eventsResponse = append(eventsResponse, createEventResponse(event))
			}
		} else {
			eventsResponse = append(eventsResponse, createEventResponse(event))
		}
	}
	data, _ := json.Marshal(eventsResponse)
	return data
}

func eventIsActive(event models.Event) bool {
	if event.Status != "active" {
		return false
	}

	eventTime := time.Date(event.Date.Year(), event.Date.Month(), event.Date.Day(), event.Time.Hour(), event.Time.Minute(), 0, 0, time.UTC)
	now := time.Now().UTC()
	return eventTime.After(now)
}

func eventIsCompleted(event models.Event) bool {
	if event.Status != "active" {
		return false
	}

	eventTime := time.Date(event.Date.Year(), event.Date.Month(), event.Date.Day(), event.Time.Hour(), event.Time.Minute(), 0, 0, time.UTC)
	now := time.Now().UTC()
	return eventTime.Before(now)
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

	eventTime := time.Date(event.Date.Year(), event.Date.Month(), event.Date.Day(), event.Time.Hour(), event.Time.Minute(), 0, 0, time.UTC)
	now := time.Now().UTC()
	if eventTime.Before(now) {
		return []byte("Event has already started")
	}

	err = repository.RegisterUserInEvent(user, event)
	if err != nil {
		return nil
	}

	return []byte("Registered to event")
}
