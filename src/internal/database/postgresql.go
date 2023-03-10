package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"strconv"
	"wiselink-challenge/src/models"
)

func PostgreSQL() (*gorm.DB, error) {
	dbHost := os.Getenv("DB_HOST")
	dbPort, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbSSLMode := os.Getenv("DB_SSL_MODE")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=%s",
		dbHost, dbPort, dbUser, dbPass, dbName, dbSSLMode)

	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db, nil
}

func Migrate() {
	db, err := PostgreSQL()
	if err != nil {
		panic(err)
	}
	defer db.Statement.Context.Done()

	err_user := db.AutoMigrate(&models.User{})
	if err_user != nil {
		return
	}
	err_event := db.AutoMigrate(&models.Event{})
	if err_event != nil {
		return
	}

	if sqlDB, err := db.DB(); err != nil {
		_ = sqlDB.Close()
	}
}
