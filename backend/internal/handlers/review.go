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

func CreateReview(w http.ResponseWriter, r *http.Request, db *sql.DB){
	var newReview models.Review
	json.NewDecoder(r.Body).Decode(&newReview)
	repository.CreateReview(newReview)
	w.WriteHeader(http.StatusCreated)
}

func GetReviews(w http.ResponseWriter, r *http.Request, db *sql.DB){
	reviews:= repository.GetReviews()
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reviews)
}

func GetReviewById(w http.ResponseWriter, r *http.Request, db *sql.DB){
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	review := repository.GetReviewById(id)
	if review == nil {
		http.Error(w, "Review not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(review)
}

func UpdateReview(w http.ResponseWriter, r *http.Request, db *sql.DB){
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	var updatedReview models.Review
	json.NewDecoder(r.Body).Decode(&updatedReview)
	updatedReview.ID = id
	
	if repository.UpdateReview(id, updatedReview){
		w.WriteHeader(http.StatusOK)
	} else {
		http.Error(w, "Review not found", http.StatusNotFound)
	}
}

func DeleteReview(w http.ResponseWriter, r *http.Request, db *sql.DB){
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	if repository.DeleteReview(id){
		w.WriteHeader(http.StatusOK)
	} else {
		http.Error(w, "Review not found", http.StatusNotFound)
	}
}