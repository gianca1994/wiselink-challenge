package routes

import (
	"encoding/json"
	"github.com/go-chi/jwtauth/v5"
	"net/http"
	"wiselink-challenge/src/cmd/service"
	jwt_auth "wiselink-challenge/src/internal/jwt"
)

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
	data := service.GetProfileService(claims)

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(data)
}
