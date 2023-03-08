package service

import (
	"encoding/json"
	"wiselink-challenge/src/internal/database"
	"wiselink-challenge/src/models"
)

func GetProfileService(claims map[string]interface{}) []byte {
	db := database.DbConnection()
	var user models.User
	var userResponse models.UserProfileResponse

	db.Where("username = ?", claims["username"]).First(&user)

	userResponse.Username = user.Username
	userResponse.Email = user.Email
	userResponse.Admin = user.Admin

	data, _ := json.Marshal(userResponse)

	return data
}
