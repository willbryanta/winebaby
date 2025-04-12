package routes

import (
	"database/sql"
	"net/http"

	"winebaby/internal/handlers"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(db *sql.DB) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Welcome to Winebaby"))
	})

	r.Get("/wines", func(w http.ResponseWriter, r *http.Request){
		handlers.GetWines(w, r, db)
	})

	r.Get("/reviews", func(w http.ResponseWriter, r *http.Request){
		handlers.GetReviews(w, r, db)})
	r.Get("/reviews/{id}", func(w http.ResponseWriter, r *http.Request){
		handlers.GetReviewById(w, r, db)})
	r.Post("/reviews", func(w http.ResponseWriter, r *http.Request){
		handlers.CreateReview(w, r, db)})
	r.Put("/reviews/{id}", func(w http.ResponseWriter, r *http.Request){
		handlers.UpdateReview(w, r, db)})
	r.Delete("/reviews/{id}", func(w http.ResponseWriter, r *http.Request){
		handlers.DeleteReview(w,r, db)})

	r.Post("/signup", func(w http.ResponseWriter, r *http.Request){
		handlers.SignUp(w,r)})
	r.Post("/signin", func(w http.ResponseWriter, r *http.Request){
		handlers.SignIn(w,r)})

	r.Get("/user/{username}", func(w http.ResponseWriter, r *http.Request){
		handlers.GetUserProfile(w,r,db)})
	r.Put("/user/{username}", func(w http.ResponseWriter, r *http.Request){
		handlers.UpdateUserProfile(w,r,db)})
	r.Delete("/user/{username}", func(w http.ResponseWriter, r *http.Request){
		handlers.DeleteUser(w,r,db)})


	r.Get("/user/{username}/wines", func(w http.ResponseWriter, r *http.Request){
		handlers.GetUserFavoriteWines(w,r,db)})
	r.Post("/user/{username}/wines", func(w http.ResponseWriter, r *http.Request){
		handlers.AddUserFavoriteWine(w,r,db)})
	r.Delete("/user/{username}/wines/{wineId}", func(w http.ResponseWriter, r *http.Request){
		handlers.RemoveUserFavoriteWine(w,r,db)})
	r.Get("/user/{username}/reviews", func(w http.ResponseWriter, r *http.Request){
		handlers.GetUserReviews(w,r,db)})
	r.Post("/user/{username}/reviews", func(w http.ResponseWriter, r *http.Request){
		handlers.CreateUserReview(w, r, db)})
	r.Put("/user/{username}/reviews/{reviewId}", func(w http.ResponseWriter, r *http.Request){
		handlers.UpdateUserReview(w, r, db)})
	r.Delete("/user/{username}/reviews/{reviewId}", func(w http.ResponseWriter, r *http.Request){
		handlers.DeleteUserReview(w, r, db)})
	r.Get("/user/{username}/reviews/{reviewId}", func(w http.ResponseWriter, r *http.Request){
		handlers.GetUserReviewById(w, r, db)})

	return r
}