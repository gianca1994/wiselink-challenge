package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	port := ":8080"
	fmt.Printf("Starting server on %v\n", port)

	http.ListenAndServe(port, routerHandler())
}

type Message struct {
	Text string `json:"text"`
}

func routerHandler() http.Handler {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		message := Message{"Hello, World!"}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		_ = json.NewEncoder(w).Encode(message)
	})

	return r
}
