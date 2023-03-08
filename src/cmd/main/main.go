package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"wiselink-challenge/src/cmd/routes"
	"wiselink-challenge/src/internal/database"
)

func main() {
	port := ":8080"
	fmt.Printf("Starting server on %v\n", port)
	database.NewPostgreSQL()
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
	})

	r.Route("/events", func(r chi.Router) {
		r.Get("/", routes.GetEvents)
		r.Post("/", routes.CreateEvent)
		//r.Get("/{id}", routes.GetEvent)
		//r.Patch("/{id}", routes.UpdateEvent)
		r.Delete("/{id}", routes.DeleteEvent)
	})

	return r
}
