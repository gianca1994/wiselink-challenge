package routes

import (
	"encoding/json"
	"net/http"
	"os"
)

func Home(w http.ResponseWriter, r *http.Request) {
	port := os.Getenv("API_PORT")
	available_routes := availableRoutes(port)

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(available_routes)
}

func availableRoutes(port string) []byte {
	routes_auth, _ := json.Marshal(map[string]string{
		"POST, Login":    "http://localhost:" + port + "api/v1/auth/login",
		"POST, Register": "http://localhost:" + port + "api/v1/auth/register",
	})

	routes_user, _ := json.Marshal(map[string]string{
		"GET, Profile":            "http://localhost:" + port + "api/v1/users/profile",
		"POST, Register to event": "http://localhost:" + port + "api/v1/users/register-event/{event_id}",
		"GET, Registered events":  "http://localhost:" + port + "api/v1/users/registered-events",
	})

	routes_event, _ := json.Marshal(map[string]string{
		"GET, Events":   "http://localhost:" + port + "api/v1/events",
		"GET, Event":    "http://localhost:" + port + "api/v1/events/{id}",
		"POST, Event":   "http://localhost:" + port + "api/v1/events",
		"PATCH, Event":  "http://localhost:" + port + "api/v1/events/{id}",
		"DELETE, Event": "http://localhost:" + port + "api/v1/events/{id}",
	})

	return []byte(`
		{
			"auth": ` + string(routes_auth) + `,
			"user": ` + string(routes_user) + `,
			"event": ` + string(routes_event) + `
		}
	`)
}
