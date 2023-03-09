package routes

import (
	"encoding/json"
	"github.com/go-chi/jwtauth/v5"
	"net/http"
	"os"
	"wiselink-challenge/src/cmd/service"
	jwt_auth "wiselink-challenge/src/internal/jwt"
)

func Home(w http.ResponseWriter, r *http.Request) {
	port := os.Getenv("API_PORT")
	available_routes, _ := json.Marshal(map[string]string{
		"POST, Login":    "http://localhost:" + port + "/auth/login",
		"POST, Register": "http://localhost:" + port + "/auth/register",
		"GET, Profile":   "http://localhost:" + port + "/users/profile",
		"GET, Events":    "http://localhost:" + port + "/events",
		"GET, Event":     "http://localhost:" + port + "/events/{id}",
		"POST, Event":    "http://localhost:" + port + "/events",
		"PATCH, Event":   "http://localhost:" + port + "/events/{id}",
		"DELETE, Event":  "http://localhost:" + port + "/events/{id}",
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
	data := service.GetProfileService(claims)

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(data)
}
