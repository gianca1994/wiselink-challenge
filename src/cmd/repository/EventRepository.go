package repository

import (
	"time"
	"wiselink-challenge/src/internal/database"
	"wiselink-challenge/src/models"
)

func CheckUserIsAdmin(claims map[string]interface{}) (bool, error) {
	var user models.User
	db, _ := database.PostgreSQL()
	defer db.Statement.Context.Done()
	db.Where("username = ?", claims["username"]).First(&user)

	sqlDB, _ := db.DB()
	_ = sqlDB.Close()

	return user.Admin, nil
}

func GetEvents() ([]models.Event, error) {
	var events []models.Event
	db, err := database.PostgreSQL()
	if err != nil {
		return events, err
	}
	defer db.Statement.Context.Done()
	db.Find(&events)

	sqlDB, _ := db.DB()
	_ = sqlDB.Close()

	if err != nil {
		return events, err
	}
	return events, nil
}

func GetEvent(id string) (models.Event, error) {
	var event models.Event
	db, err := database.PostgreSQL()
	if err != nil {
		return event, err
	}
	defer db.Statement.Context.Done()
	db.Where("id = ?", id).First(&event)

	sqlDB, _ := db.DB()
	_ = sqlDB.Close()

	if err != nil {
		return event, err
	}
	return event, nil
}

func GetUsersForEvent(event *models.Event) ([]models.UserEventResponse, error) {
	db, err := database.PostgreSQL()
	if err != nil {
		return nil, err
	}
	defer db.Statement.Context.Done()

	err = db.Model(&event).Association("Users").Find(&event.Users)
	if err != nil {
		return nil, err
	}

	var usersResponse []models.UserEventResponse
	for _, user := range event.Users {
		usersResponse = append(usersResponse, models.UserEventResponse{
			Username: user.Username,
			Email:    user.Email,
		})
	}

	return usersResponse, nil
}

func CreateEvent(event models.EventCreate) ([]byte, error) {
	db, _ := database.PostgreSQL()

	dateEvent, _ := time.Parse("2006-01-02", event.Date)
	timeEvent, _ := time.Parse("15:04", event.Time)
	db.Create(&models.Event{
		Title:     event.Title,
		ShortDesc: event.ShortDesc,
		LongDesc:  event.LongDesc,
		Date:      dateEvent,
		Time:      timeEvent,
		Organizer: event.Organizer,
		Place:     event.Place,
		Status:    "draft",
	})
	sqlDB, _ := db.DB()
	_ = sqlDB.Close()
	return []byte("Event created"), nil
}

func UpdateEvent(param string, event models.EventUpdate) ([]byte, error) {
	var eventDB models.Event
	db, _ := database.PostgreSQL()
	db.Where("id = ?", param).First(&eventDB)

	if eventDB.Id == 0 {
		return []byte("Event not found"), nil
	}

	dateEvent, _ := time.Parse("2006-01-02", event.Date)
	timeEvent, _ := time.Parse("15:04", event.Time)
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
	sqlDB, _ := db.DB()
	_ = sqlDB.Close()
	return []byte("Event updated"), nil
}

func DeleteEvent(idDeleted string) ([]byte, error) {
	var event models.Event
	db, _ := database.PostgreSQL()
	db.Where("id = ?", idDeleted).First(&event)
	if event.Id == 0 {
		return []byte("Invalid event"), nil
	}
	db.Delete(&event)
	sqlDB, _ := db.DB()
	_ = sqlDB.Close()
	return []byte("Event deleted"), nil
}
