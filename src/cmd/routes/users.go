package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"net/http"
	"wiselink-challenge/src/cmd/service"
	jwt_auth "wiselink-challenge/src/internal/jwt"
)

func GetProfile(w http.ResponseWriter, r *http.Request) {
	claims := jwt_auth.TokenGetClaims(jwtauth.TokenFromHeader(r), w)
	if claims == nil {
		return
	}

	data, _ := service.GetProfileService(claims)
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(data)
}

func GetRegisteredEvents(w http.ResponseWriter, r *http.Request) {
	claims := jwt_auth.TokenGetClaims(jwtauth.TokenFromHeader(r), w)
	if claims == nil {
		return
	}

	filter := r.URL.Query().Get("filter")
	data, _ := service.GetRegisteredEvents(claims, filter)
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(data)
}

func RegisterToEvent(w http.ResponseWriter, r *http.Request) {
	claims := jwt_auth.TokenGetClaims(jwtauth.TokenFromHeader(r), w)
	if claims == nil {
		return
	}

	response := service.RegisterToEvent(claims, chi.URLParam(r, "event_id"))
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(response)
}
