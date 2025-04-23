package repository

import (
	"database/sql"
	"winebaby/internal/models"
)

var reviews = []models.Review{
	{ID: 1, WineID:1, Title: "McClaren Vale 2024", Description: "Zesty and fruity with notes of licorice", Rating: 7},
}

// Get all reviews
func GetReviews(r *MainRepository) ([]models.Review, error){
	query := `SELECT id, user_id, wine_id, comment, review_date, review_date_time, review_date_time_utc, title, description, rating 
              FROM reviews`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reviews []models.Review
	for rows.Next(){
		var review models.Review
		err := rows.Scan(
			&review.ID,
			&review.UserID,
			&review.WineID,
			&review.Comment,
			&review.ReviewDate,
			&review.ReviewDateTime,
			&review.ReviewDateTimeUTC,
			&review.Title,
			&review.Description,
			&review.Rating,
		)
		if err != nil {
			return nil, err
		}
		reviews = append(reviews, review)
	}
	return reviews, nil
}

// Get a review by ID
func (r *MainRepository) GetReviewById(id int) (models.Review, error) {
	query := `SELECT id, user_id, wine_id, comment, review_date, review_date_time, review_date_time_utc, title, description, rating 
              FROM reviews WHERE id = $1`
	var review models.Review
	err := r.DB.QueryRow(query, id).Scan(
		&review.ID,
		&review.UserID,
		&review.WineID,
		&review.Comment,
		&review.ReviewDate,
		&review.ReviewDateTime,
		&review.ReviewDateTimeUTC,
		&review.Title,
		&review.Description,
		&review.Rating,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Review{}, nil
		}
		return models.Review{}, err
	}
	return review, nil
}


func CreateReview(r models.Review){
	reviews = append(reviews, r)
	
}

func UpdateReview(id int, updated models.Review) bool {
	for i, r := range reviews{
		if r.ID == id {
			reviews[i] = updated
			return true
		}
	}
	return false
}

func DeleteReview(id int) bool {
	for i, r:= range reviews {
		if r.ID ==id {
			reviews = append(reviews[:i], reviews[i+1:]...)
			return true
		}
	}
	return false
}