package routes

import (
	"github.com/go-chi/chi/v5"
	"winebaby/internal/handlers"
	"net/http"
)

func RegisterRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r * http.Request){
		w.Write([]byte("Welcome to Winebaby"))
	})

	r.Get("/wines", handlers.GetWines)

	r.Get("/reviews", handlers.GetReviews)
	r.Get("/reviews/{id}", handlers.GetReviewById)
	r.Post("/reviews", handlers.CreateReview)
	r.Put("/reviews/{id}", handlers.UpdateReview)
	r.Delete("/reviews/{id}", handlers.DeleteReview)

	return r
}