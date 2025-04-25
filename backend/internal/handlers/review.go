package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"winebaby/internal/models"
	"winebaby/internal/repository"

	"github.com/go-chi/chi/v5"
)

func CreateReview(w http.ResponseWriter, r *http.Request, repo *repository.MainRepository, db *sql.DB){
	var newReview models.Review

	if err := json.NewDecoder(r.Body).Decode(&newReview); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(chi.URLParam(r, "wineId"))
	if err != nil {
		http.Error(w, "Invalid wine ID", http.StatusBadRequest)
		return
	}
	newReview.WineID = id
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newReview)
}

func GetReviews(w http.ResponseWriter, r *http.Request, repo *repository.MainRepository, db *sql.DB){
	reviews, err := repository.GetReviews(repo)
	if err != nil {
		http.Error(w, "Failed to fetch reviews", http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reviews)
}

func GetReviewById(w http.ResponseWriter, r *http.Request, repo *repository.MainRepository, db *sql.DB){
	idStr := chi.URLParam(r, "reviewId")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid review ID", http.StatusBadRequest)
		return
	}
	review, err := repository.GetReviewById(repo, id)
	if err != nil {
		http.Error(w, "Failed to fetch review", http.StatusInternalServerError)
		return
	}
	if review.ID == 0 {
		http.Error(w, "Review not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(review)
}

func UpdateReview(w http.ResponseWriter, r *http.Request, repo *repository.MainRepository, db *sql.DB){
	idStr := chi.URLParam(r, "reviewId")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid review ID", http.StatusBadRequest)
		return
	}
	var updatedReview models.Review
	json.NewDecoder(r.Body).Decode(&updatedReview)
	updatedReview.ID = id
	
	err = repository.UpdateReview(repo, id, updatedReview)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Review not found", http.StatusNotFound)
		} else {
			http.Error(w, "Failed to update review", http.StatusInternalServerError)
		}
		return
	}
	w.WriteHeader(http.StatusOK)
}

func DeleteReview(w http.ResponseWriter, r *http.Request, repo *repository.MainRepository, db *sql.DB){
	idStr := chi.URLParam(r, "reviewId")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid review ID", http.StatusBadRequest)
		return
	}
	err = repository.DeleteReview(repo, id)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Review not found", http.StatusNotFound)
		} else {
			http.Error(w, "Failed to delete review", http.StatusInternalServerError)
		}
		return
	}
	w.WriteHeader(http.StatusOK)
}