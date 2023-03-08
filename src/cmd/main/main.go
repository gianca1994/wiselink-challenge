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

	r.Group(func(r chi.Router) {
		r.Get("/profile", routes.GetProfile)
	})

	return r
}
