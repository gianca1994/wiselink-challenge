package routes

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/jwtauth/v5"
	"gorm.io/gorm"
	"net/http"
	"os"
	"wiselink-challenge/src/internal/database"
	jwt_auth "wiselink-challenge/src/internal/jwt_bearer"
	"wiselink-challenge/src/models"
)

func DbConnection() *gorm.DB {
	db := database.NewPostgreSQL()

	if db == nil {
		fmt.Println("Error connecting to database")
		os.Exit(0)
	}
	return db
}

func Home(w http.ResponseWriter, r *http.Request) {
	port := ":8080"
	available_routes, _ := json.Marshal(map[string]string{
		"POST, Login":    "http://localhost" + port + "/auth/login",
		"POST, Register": "http://localhost" + port + "/auth/register",
		"GET, Profile":   "http://localhost" + port + "/profile",
	})

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(available_routes)
}

func GetProfile(w http.ResponseWriter, r *http.Request) {
	claims := jwt_auth.TokenGetClaims(jwtauth.TokenFromHeader(r))

	if claims == nil {
		_, _ = w.Write([]byte("Invalid token"))
		return
	}

	db := DbConnection()
	var user models.User

	db.Where("username = ?", claims["username"]).First(&user)
	data, _ := json.Marshal(user)

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(data)
}
