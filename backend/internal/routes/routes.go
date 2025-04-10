package routes

import (
	"net/http"
	"winebaby/internal/handlers"

	"github.com/go-chi/chi/v5"
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

	r.Post("/signup", handlers.SignUp)
	r.Post("/signin", handlers.SignIn)
	r.Get("/user/{username}", handlers.GetUserProfile)
	r.Put("/user/{username}", handlers.UpdateUserProfile)
	r.Delete("/user/{username}", handlers.DeleteUser)
	r.Get("/user/{username}/wines", handlers.GetUserFavoriteWines)
	r.Post("/user/{username}/wines", handlers.AddUserFavoriteWine)
	r.Delete("/user/{username}/wines/{wineId}", handlers.RemoveUserFavoriteWine)
	r.Get("/user/{username}/reviews", handlers.GetUserReviews)
	r.Post("/user/{username}/reviews", handlers.CreateUserReview)
	r.Put("/user/{username}/reviews/{reviewId}", handlers.UpdateUserReview)
	r.Delete("/user/{username}/reviews/{reviewId}", handlers.DeleteUserReview)
	r.Get("/user/{username}/reviews/{reviewId}", handlers.GetUserReviewById)

	return r
}