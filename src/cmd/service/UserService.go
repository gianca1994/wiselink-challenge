package service

import (
	"encoding/json"
	"wiselink-challenge/src/internal/database"
	"wiselink-challenge/src/models"
)

func GetProfileService(claims map[string]interface{}) []byte {
	db := database.PostgreSQL()
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

func RegisterToEvent(claims map[string]interface{}, event_id string) {
	db := database.PostgreSQL()
	var user models.User
	var event models.Event
	db.Where("username = ?", claims["username"]).First(&user)
	db.Where("id = ?", event_id).First(&event)
	err := db.Model(&user).Association("Events").Append(&event)
	if err != nil {
		return
	}
}
