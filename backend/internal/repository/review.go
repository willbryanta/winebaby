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

func (r *MainRepository) CreateReview(review models.Review) (int, error) {
	query := `INSERT INTO reviews (user_id, wine_id, comment, review_date, review_date_time, review_date_time_utc, title, description, rating) 
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`
	var id int
	err := r.DB.QueryRow(query,
		review.UserID,
		review.WineID,
		review.Comment,
		review.ReviewDate,
		review.ReviewDateTime,
		review.ReviewDateTimeUTC,
		review.Title,
		review.Description,
		review.Rating,
	).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *MainRepository) UpdateReview(id int, updated models.Review) error {
	query := `UPDATE reviews SET user_id = $1, wine_id = $2, comment = $3, review_date = $4, review_date_time = $5, 
                             review_date_time_utc = $6, title = $7, description = $8, rating = $9 
              WHERE id = $10`
	_, err := r.DB.Exec(query,
		updated.UserID,
		updated.WineID,
		updated.Comment,
		updated.ReviewDate,
		updated.ReviewDateTime,
		updated.ReviewDateTimeUTC,
		updated.Title,
		updated.Description,
		updated.Rating,
		id,
	)
	return err
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