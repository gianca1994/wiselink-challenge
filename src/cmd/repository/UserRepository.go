package repository

import (
	"wiselink-challenge/src/internal/database"
	"wiselink-challenge/src/models"
)

func CreateUser(user models.User) (models.User, error) {
	db, err := database.PostgreSQL()
	if err != nil {
		return user, err
	}
	defer db.Statement.Context.Done()

	err = db.Create(&user).Error
	if sqlDB, err := db.DB(); err == nil {
		_ = sqlDB.Close()
	}
	if err != nil {
		return user, err
	}
	return user, nil
}

func GetUserByUsername(username string) (models.User, error) {
	var user models.User
	db, err := database.PostgreSQL()
	if err != nil {
		return user, err
	}
	defer db.Statement.Context.Done()

	err = db.Where("username = ?", username).First(&user).Error
	if sqlDB, err := db.DB(); err == nil {
		_ = sqlDB.Close()
	}
	if err != nil {
		return user, err
	}
	return user, nil
}

func GetUserByEmail(email string) (models.User, error) {
	var user models.User
	db, err := database.PostgreSQL()
	if err != nil {
		return user, err
	}
	defer db.Statement.Context.Done()

	err = db.Where("email = ?", email).First(&user).Error
	if sqlDB, err := db.DB(); err == nil {
		_ = sqlDB.Close()
	}
	if err != nil {
		return user, err
	}
	return user, nil
}

func GetRegisteredEvents(userId uint) ([]models.Event, error) {
	var events []models.Event
	db, err := database.PostgreSQL()
	if err != nil {
		return events, err
	}
	defer db.Statement.Context.Done()

	err = db.Model(&models.User{Id: userId}).Association("Events").Find(&events)
	if sqlDB, err := db.DB(); err == nil {
		_ = sqlDB.Close()
	}
	if err != nil {
		return events, err
	}
	return events, nil
}
