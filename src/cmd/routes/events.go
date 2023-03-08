package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"net/http"
	"wiselink-challenge/src/cmd/service"
	jwt_auth "wiselink-challenge/src/internal/jwt_bearer"
)

func GetEvents(w http.ResponseWriter, r *http.Request) {
	claims := jwt_auth.TokenGetClaims(jwtauth.TokenFromHeader(r))
	if claims == nil {
		_, _ = w.Write([]byte("Invalid token"))
		return
	}

	data := service.GetEventsService()
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(data)
}

func GetEvent(w http.ResponseWriter, r *http.Request) {
	claims := jwt_auth.TokenGetClaims(jwtauth.TokenFromHeader(r))
	if claims == nil {
		_, _ = w.Write([]byte("Invalid token"))
		return
	}

	data, _ := service.GetEvent(chi.URLParam(r, "id"))

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(data)
}

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	claims := jwt_auth.TokenGetClaims(jwtauth.TokenFromHeader(r))
	if claims == nil {
		_, _ = w.Write([]byte("Invalid token"))
		return
	}
	data, _ := service.CreateEventService(claims, r)

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(data)
}

func UpdateEvent(w http.ResponseWriter, r *http.Request) {
	claims := jwt_auth.TokenGetClaims(jwtauth.TokenFromHeader(r))
	if claims == nil {
		_, _ = w.Write([]byte("Invalid token"))
		return
	}

	data, _ := service.UpdateEventService(claims, chi.URLParam(r, "id"), r)

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(data)
}

func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	claims := jwt_auth.TokenGetClaims(jwtauth.TokenFromHeader(r))
	if claims == nil {
		_, _ = w.Write([]byte("Invalid token"))
		return
	}

	data, _ := service.DeleteEventService(claims, chi.URLParam(r, "id"))

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(data)
}
