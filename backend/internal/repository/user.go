package repository

import (
	"database/sql"
	"winebaby/internal/models"
)

func (r *MainRepository) SignUp(username, password, email string) (models.User, error) {
	query := `INSERT INTO users (username, password, email) VALUES ($1, $2, $3) RETURNING id`
	var user models.User
	var id int
	var emailPtr *string
	if email != "" {
		emailPtr = &email
	}
	err := r.DB.QueryRow(query, username, password, emailPtr).Scan(&id)
	if err != nil {
		return models.User{}, err
	}
	user.ID = id
	user.Username = username
	user.Password = password
	user.Email = emailPtr
	return user, nil
}

func (r *MainRepository) CreateUser(user models.User) error {
	query := `INSERT INTO users (username, password, email) VALUES ($1, $2, $3) RETURNING id`
	err := r.DB.QueryRow(query, user.Username, user.Password, user.Email).Scan(&user.ID)
	return err
}

func (r *MainRepository) GetUserByUsername(username string) (models.User, error) {
	query := `SELECT id, username, email, password FROM users WHERE username = $1`
	var user models.User
	var email sql.NullString
	err := r.DB.QueryRow(query, username).Scan(&user.ID, &user.Username, &email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, nil
		}
		return models.User{}, err
	}
	if email.Valid {
		user.Email = &email.String
	}
	return user, nil
}

func (r *MainRepository) GetUserProfile(username string) (models.UserProfile, error) {
	userQuery := `SELECT id, username, email FROM users WHERE username = $1`
	var user models.UserProfile
	var email sql.NullString
	err := r.DB.QueryRow(userQuery, username).Scan(&user.ID, &user.Username, &email)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.UserProfile{}, nil
		}
		return models.UserProfile{}, err
	}
	if email.Valid {
		user.Email = &email.String
	}

	winesQuery := `SELECT id, name, year, manufacturer, region, alcohol_content, serving_temp, serving_size, 
                          serving_size_unit, serving_size_unit_abbreviation, serving_size_unit_description, 
                          serving_size_unit_description_abbreviation, serving_size_unit_description_plural, 
                          price, rating, type, colour 
                   FROM favorite_wines WHERE user_id = $1`
	rows, err := r.DB.Query(winesQuery, user.ID)
	if err != nil {
		return models.UserProfile{}, err
	}
	defer rows.Close()
	for rows.Next() {
		var wine models.Wine
		err := rows.Scan(
			&wine.ID,
			&wine.Name,
			&wine.Year,
			&wine.Manufacturer,
			&wine.Region,
			&wine.AlcoholContent,
			&wine.Price,
			&wine.Rating,
			&wine.Type,
			&wine.Colour,
		)
		if err != nil {
			return models.UserProfile{}, err
		}
		user.FavoriteWines = append(user.FavoriteWines, wine)
	}

	reviewsQuery := `SELECT id, wine_id, winemaker, wine_name, comment, review_date, review_date_time, 
                            review_date_time_utc, title, description, rating 
                     FROM reviews WHERE user_id = $1`
	rows, err = r.DB.Query(reviewsQuery, user.ID)
	if err != nil {
		return models.UserProfile{}, err
	}
	defer rows.Close()
	for rows.Next() {
		var review models.Review
		err := rows.Scan(
			&review.ID,
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
			return models.UserProfile{}, err
		}
		user.Reviews = append(user.Reviews, review)
	}
	return user, nil
}
func (r *MainRepository) GetUserByEmail(email string) (models.User, error) {
	query := `SELECT id, username, email, password FROM users WHERE email = $1`
	var user models.User
	var emailVal sql.NullString
	err := r.DB.QueryRow(query, email).Scan(&user.ID, &user.Username, &emailVal, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, nil
		}
		return models.User{}, err
	}
	if emailVal.Valid {
		user.Email = &emailVal.String
	}
	return user, nil
}

func (r *MainRepository) UpdateUserProfile(user models.User) error {
	query := `UPDATE users SET username = $1, email = $2 WHERE id = $3`
	_, err := r.DB.Exec(query, user.Username, user.Email, user.ID)
	return err
}

func (r *MainRepository) DeleteUser(username string) error {
	query := `DELETE FROM users WHERE username = $1`
	_, err := r.DB.Exec(query, username)
	return err
}

func (r *MainRepository) GetUserFavoriteWines(username string) ([]models.Wine, error) {
	query := `SELECT id, name, year, manufacturer, region, alcohol_content, serving_temp, serving_size, 
						  serving_size_unit, serving_size_unit_abbreviation, serving_size_unit_description, 
						  serving_size_unit_description_abbreviation, serving_size_unit_description_plural, 
						  price, rating, type, colour 
				   FROM favorite_wines WHERE user_id = (SELECT id FROM users WHERE username = $1)`
	rows, err := r.DB.Query(query, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var wines []models.Wine
	for rows.Next() {
		var wine models.Wine
		err := rows.Scan(
			&wine.ID,
			&wine.Name,
			&wine.Year,
			&wine.Manufacturer,
			&wine.Region,
			&wine.AlcoholContent,

			&wine.Price,
			&wine.Rating,
			&wine.Type,
			&wine.Colour,
		)
		if err != nil {
			return nil, err
		}
		wines = append(wines, wine)
	}
	return wines, nil
}

func (r *MainRepository) AddUserFavoriteWine(userID int, wineID int) error {
	query := `INSERT INTO favorite_wines (user_id, wine_id) VALUES ($1, $2)`
	_, err := r.DB.Exec(query, userID, wineID)
	if err != nil {
		return err
	}
	return nil
}

func (r *MainRepository) RemoveUserFavoriteWine(userID int, wineID int) error {
	query := `DELETE FROM favorite_wines WHERE user_id = $1 AND wine_id = $2`
	_, err := r.DB.Exec(query, userID, wineID)
	return err
}

func (r *MainRepository) GetUserReviews(username string) ([]models.Review, error) {
	query := `SELECT id, wine_id, winemaker, wine_name, comment, review_date, review_date_time, 
						  review_date_time_utc, title, description, rating 
				   FROM reviews WHERE user_id = (SELECT id FROM users WHERE username = $1)`
	rows, err := r.DB.Query(query, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reviews []models.Review
	for rows.Next() {
		var review models.Review
		err := rows.Scan(
			&review.ID,
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

func (r *MainRepository) CreateUserReview(review models.Review) error {
	query := `INSERT INTO reviews (user_id, wine_id, winemaker, wine_name, comment, review_date, 
								  review_date_time, review_date_time_utc, title, description, rating) 
			  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id`
	err := r.DB.QueryRow(query,
		review.UserID,
		review.WineID,
		review.Comment,
		review.ReviewDate,
		review.ReviewDateTime,
		review.ReviewDateTimeUTC,
		review.Title,
		review.Description,
		review.Rating).Scan(&review.ID)
	return err
}

func (r *MainRepository) UpdateUserReview(review models.Review) error {
	query := `UPDATE reviews SET wine_id = $1, winemaker = $2, wine_name = $3, comment = $4, review_date = $5, 
								  review_date_time = $6, review_date_time_utc = $7, title = $8, description = $9, rating = $10 
			  WHERE id = $11`
	_, err := r.DB.Exec(query,
		review.WineID,
		review.Comment,
		review.ReviewDate,
		review.ReviewDateTime,
		review.ReviewDateTimeUTC,
		review.Title,
		review.Description,
		review.Rating,
		review.ID)
	return err
}
func (r *MainRepository) DeleteUserReview(reviewID int) error {
	query := `DELETE FROM reviews WHERE id = $1`
	_, err := r.DB.Exec(query, reviewID)
	return err
}

func (r *MainRepository) GetUserReviewById(reviewID int) (models.Review, error) {
	query := `SELECT id, wine_id, winemaker, wine_name, comment, review_date, review_date_time, 
						  review_date_time_utc, title, description, rating 
				   FROM reviews WHERE id = $1`
	var review models.Review
	err := r.DB.QueryRow(query, reviewID).Scan(
		&review.ID,
		&review.WineID,
		&review.Comment,
		&review.ReviewDate,
		&review.ReviewDateTime,
		&review.ReviewDateTimeUTC,
		&review.Title,
		&review.Description,
		&review.Rating)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Review{}, nil
		}
		return models.Review{}, err
	}
	return review, nil
}
