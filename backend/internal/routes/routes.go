package routes

import (
	"database/sql"
	"net/http"

	"winebaby/internal/handlers"
	"winebaby/internal/repository"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(db *sql.DB) *chi.Mux {
	r := chi.NewRouter()

	repo := &repository.MainRepository{DB: db}

	r.Get("/verify-token", func(w http.ResponseWriter, r *http.Request){
		handlers.VerifyToken(w, r)
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Welcome to Winebaby"))
	})

	r.Get("/wines", func(w http.ResponseWriter, r *http.Request){
		handlers.GetWines(w, r, repo, db)
	})
	r.Post("/wines", func(w http.ResponseWriter, r *http.Request){
		handlers.AddWine(w, r, repo, db)
	})

	r.Get("/reviews", func(w http.ResponseWriter, r *http.Request){
		handlers.GetReviews(w, r, repo, db)})
	r.Get("/reviews/{reviewId}", func(w http.ResponseWriter, r *http.Request){
		handlers.GetReviewById(w, r, repo, db)})
	r.Post("/reviews", func(w http.ResponseWriter, r *http.Request){
		handlers.CreateReview(w, r, repo, db)})
	r.Put("/reviews/{reviewId}", func(w http.ResponseWriter, r *http.Request){
		handlers.UpdateReview(w, r, repo, db)})
	r.Delete("/reviews/{reviewId}", func(w http.ResponseWriter, r *http.Request){
		handlers.DeleteReview(w,r, repo, db)})

	r.Post("/signup", func(w http.ResponseWriter, r *http.Request){
		handlers.SignUp(w,r, repo, db)})
	r.Post("/signin", func(w http.ResponseWriter, r *http.Request){
		handlers.SignIn(w,r, repo, db)})
	r.Post("/signout", func(w http.ResponseWriter, r *http.Request){
		handlers.SignOut(w,r)})

	r.Get("/user/{username}", func(w http.ResponseWriter, r *http.Request){
		handlers.GetUserProfile(w,r, repo, db)})
	r.Put("/user/{username}", func(w http.ResponseWriter, r *http.Request){
		handlers.UpdateUserProfile(w,r, repo, db)})
	
	return r
}