package service

import (
	"encoding/json"
	"time"
	"wiselink-challenge/src/internal/database"
	"wiselink-challenge/src/models"
)

func GetProfileService(claims map[string]interface{}) []byte {
	db := database.PostgreSQL()
	var user models.User
	var userResponse models.UserProfileResponse

	db.Where("username = ?", claims["username"]).First(&user)

	userResponse.Username = user.Username
	userResponse.Email = user.Email
	userResponse.Admin = user.Admin
	_ = db.Model(&user).Association("Events").Find(&userResponse.Events)

	for i := range userResponse.Events {
		dateEvent, _ := time.Parse(time.RFC3339Nano, userResponse.Events[i].Date)
		timeEvent, _ := time.Parse(time.RFC3339, userResponse.Events[i].Time)
		userResponse.Events[i].Date = dateEvent.Format("02-01-2006")
		userResponse.Events[i].Time = timeEvent.Format("15:04")
	}
	data, _ := json.Marshal(userResponse)
	return data
}
