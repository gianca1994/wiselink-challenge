package routes

import (
	"encoding/json"
	"net/http"
	"os"
)

func Home(w http.ResponseWriter, r *http.Request) {
	port := os.Getenv("API_PORT")

	routes_auth, _ := json.Marshal(map[string]string{
		"POST, Login":    "http://localhost:" + port + "/auth/login",
		"POST, Register": "http://localhost:" + port + "/auth/register",
	})

	routes_user, _ := json.Marshal(map[string]string{
		"GET, Profile": "http://localhost:" + port + "/users/profile",
	})

	routes_event, _ := json.Marshal(map[string]string{
		"GET, Events":   "http://localhost:" + port + "/events",
		"GET, Event":    "http://localhost:" + port + "/events/{id}",
		"POST, Event":   "http://localhost:" + port + "/events",
		"PATCH, Event":  "http://localhost:" + port + "/events/{id}",
		"DELETE, Event": "http://localhost:" + port + "/events/{id}",
	})

	available_routes := []byte(`
		{
			"auth": ` + string(routes_auth) + `,
			"user": ` + string(routes_user) + `,
			"event": ` + string(routes_event) + `
		}
	`)

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(available_routes)
}
