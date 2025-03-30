package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"winebaby/internal/routes"
)

func main() {
	r := chi.NewRouter()
	r.Mount("/", routes.RegisterRoutes())

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", r)
}