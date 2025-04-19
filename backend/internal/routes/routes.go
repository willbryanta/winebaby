package routes

import (
	"database/sql"
	"net/http"

	"winebaby/internal/handlers"

	"github.com/go-chi/chi/v5"
)

//TODO - may want to move struct and NewRepo to a separate file
type Repository struct {
    DB *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		DB: db,
	}
}

func RegisterRoutes(db *sql.DB) *chi.Mux {
	r := chi.NewRouter()

	repo := NewRepository(db) //Todo - may need to update this to repository.Repository

	r.Get("/", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Welcome to Winebaby"))
	})

	r.Get("/wines", func(w http.ResponseWriter, r *http.Request){
		handlers.GetWines(w, r)
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
		handlers.SignUp(w,r, db)})
	r.Post("/signin", func(w http.ResponseWriter, r *http.Request){
		handlers.SignIn(w,r, db)})

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