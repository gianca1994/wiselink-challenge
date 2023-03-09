package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/golang-jwt/jwt"
	"net/http"
	"wiselink-challenge/src/cmd/service"
	jwt_auth "wiselink-challenge/src/internal/jwt"
)

func handlerTokenClaims(w http.ResponseWriter, r *http.Request) jwt.MapClaims {
	claims := jwt_auth.TokenGetClaims(jwtauth.TokenFromHeader(r))
	if claims == nil {
		_, _ = w.Write([]byte("Invalid token"))
		return nil
	}
	return claims
}

func GetEvents(w http.ResponseWriter, r *http.Request) {
	claims := handlerTokenClaims(w, r)

	data := service.GetEventsService(claims)
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(data)
}

func GetEvent(w http.ResponseWriter, r *http.Request) {
	claims := handlerTokenClaims(w, r)

	data, _ := service.GetEvent(claims, chi.URLParam(r, "id"))
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(data)
}

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	claims := handlerTokenClaims(w, r)

	data, _ := service.CreateEventService(claims, r)
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(data)
}

func UpdateEvent(w http.ResponseWriter, r *http.Request) {
	claims := handlerTokenClaims(w, r)

	data, _ := service.UpdateEventService(claims, chi.URLParam(r, "id"), r)
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(data)
}

func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	claims := handlerTokenClaims(w, r)

	data, _ := service.DeleteEventService(claims, chi.URLParam(r, "id"))
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(data)
}
