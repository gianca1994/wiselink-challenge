package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"net/http"
	"os"
	"wiselink-challenge/src/cmd/routes"
	"wiselink-challenge/src/internal/database"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(0)
	}

	port := ":" + os.Getenv("API_PORT")
	database.Migrate()

	fmt.Printf("Starting server on %v\n", port)
	_ = http.ListenAndServe(port, routerHandler())
}

func routerHandler() http.Handler {
	r := chi.NewRouter()

	r.Get("/", routes.Home)

	r.Route("/auth", func(r chi.Router) {
		r.Post("/login", routes.Login)
		r.Post("/register", routes.Register)
	})

	r.Route("/users", func(r chi.Router) {
		r.Get("/profile", routes.GetProfile)
		r.Post("/register-event/{event_id}", routes.RegisterToEvent)
	})

	r.Route("/events", func(r chi.Router) {
		r.Get("/", routes.GetEvents)
		r.Get("/{id}", routes.GetEvent)

		r.Post("/", routes.CreateEvent)

		r.Patch("/{id}", routes.UpdateEvent)
		r.Delete("/{id}", routes.DeleteEvent)
	})

	return r
}
