package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"net/http"
	"os"
	"wiselink-challenge/src/cmd/routes"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(0)
	}
	//database.Init()
	//database.Migrate()

	port := ":" + os.Getenv("API_PORT")
	fmt.Printf("Starting server on %v\n", port)
	_ = http.ListenAndServe(port, routerHandler())
}

func routerHandler() http.Handler {
	r := chi.NewRouter()

	corsOrigin := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PATCH", "DELETE"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type"},
		MaxAge:         300,
	})
	r.Use(corsOrigin.Handler)

	r.Get("/", routes.Home)

	r.Route("/api/v1/auth", func(r chi.Router) {
		r.Post("/login", routes.Login)
		r.Post("/register", routes.Register)
	})

	r.Route("/api/v1/users", func(r chi.Router) {
		r.Get("/profile", routes.GetProfile)
		r.Post("/register-event/{event_id}", routes.RegisterToEvent)
		r.Get("/registered-events", routes.GetRegisteredEvents)
	})

	r.Route("/api/v1/events", func(r chi.Router) {
		r.Get("/", routes.GetEvents)
		r.Get("/{id}", routes.GetEvent)

		r.Post("/", routes.CreateEvent)

		r.Patch("/{id}", routes.UpdateEvent)
		r.Delete("/{id}", routes.DeleteEvent)
	})

	return r
}
