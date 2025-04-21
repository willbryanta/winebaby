package routes

import (
	"database/sql"
	"net/http"

	"winebaby/internal/handlers"

	"github.com/go-chi/chi/v5"
)
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

	repo := &Repository{DB: db}

	r.Get("/", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Welcome to Winebaby"))
	})

	r.Get("/wines", func(w http.ResponseWriter, r *http.Request){
		handlers.GetWines(w, r, repo, db)
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

	r.Get("/user/{username}", func(w http.ResponseWriter, r *http.Request){
		handlers.GetUserProfile(w,r, repo, db)})
	r.Put("/user/{username}", func(w http.ResponseWriter, r *http.Request){
		handlers.UpdateUserProfile(w,r, repo, db)})
	r.Delete("/user/{username}", func(w http.ResponseWriter, r *http.Request){
		handlers.DeleteUser(w,r, repo, db)})


	r.Get("/user/{username}/wines", func(w http.ResponseWriter, r *http.Request){
		handlers.GetUserFavoriteWines(w,r, repo, db)})
	r.Post("/user/{username}/wines/{wineId}", func(w http.ResponseWriter, r *http.Request){
		handlers.AddUserFavoriteWine(w,r, repo, db)})
	r.Delete("/user/{username}/wines/{wineId}", func(w http.ResponseWriter, r *http.Request){
		handlers.RemoveUserFavoriteWine(w,r, repo, db)})
	r.Get("/user/{username}/reviews", func(w http.ResponseWriter, r *http.Request){
		handlers.GetUserReviews(w,r, repo, db)})
	r.Post("/user/{username}/reviews", func(w http.ResponseWriter, r *http.Request){
		handlers.CreateUserReview(w, r, repo, db)})
	r.Put("/user/{username}/reviews/{reviewId}", func(w http.ResponseWriter, r *http.Request){
		handlers.UpdateUserReview(w, r, repo, db)})
	r.Delete("/user/{username}/reviews/{reviewId}", func(w http.ResponseWriter, r *http.Request){
		handlers.DeleteUserReview(w, r, repo, db)})
	r.Get("/user/{username}/reviews/{reviewId}", func(w http.ResponseWriter, r *http.Request){
		handlers.GetUserReviewById(w, r, repo, db)})
		

	return r
}