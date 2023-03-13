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
	claims := jwt_auth.TokenGetClaims(jwtauth.TokenFromHeader(r), w)
	if claims == nil {
		return nil
	}
	return claims
}

func GetEvents(w http.ResponseWriter, r *http.Request) {
	claims := handlerTokenClaims(w, r)

	filterBy, filterValue := extractFilter(r)
	data := service.GetEventsService(claims, filterBy, filterValue)
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

///////////////////// Helpers //////////////////////

func extractFilter(r *http.Request) (string, string) {
	var filterSelected string
	var filter string

	filterDate := r.URL.Query().Get("date")
	if filterDate != "" {
		filter = filterDate
		filterSelected = "date"
	}
	filterStatus := r.URL.Query().Get("status")
	if filterStatus != "" {
		filter = filterStatus
		filterSelected = "status"
	}
	filterTitle := r.URL.Query().Get("title")
	if filterTitle != "" {
		filter = filterTitle
		filterSelected = "title"
	}

	return filterSelected, filter
}
