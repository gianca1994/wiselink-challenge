package database

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"regexp"
	"strconv"
	"wiselink-challenge/src/models"
)

func DbConnection() *gorm.DB {
	db := NewPostgreSQL()

	if db == nil {
		fmt.Println("Error connecting to database")
		os.Exit(0)
	}
	return db
}

func NewPostgreSQL() *gorm.DB {
	projectName := regexp.MustCompile(`^(.*` + "wiselink-challenge" + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	if err := godotenv.Load(string(rootPath) + `/.env`); err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(0)
	}

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

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		return nil
	}

	return db
}
