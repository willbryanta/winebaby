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

func (r *MainRepository) CreateUser(user *models.User) error {
	query := `INSERT INTO users (username, password, email) VALUES ($1, $2, $3) RETURNING id`
	email := ""
	if user.Email != nil {
		email = *user.Email
	}
	err := r.DB.QueryRow(query, user.Username, user.Password, email).Scan(&user.ID)
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

func (r *MainRepository) GetUserProfile(username string) (models.User, error) {
	userQuery := `SELECT id, username, email FROM users WHERE username = $1`
	var user models.User
	var email sql.NullString
	err := r.DB.QueryRow(userQuery, username).Scan(&user.ID, &user.Username, &email)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, nil
		}
		return models.User{}, err
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
		return models.User{}, err
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
			return models.User{}, err
		}
		user.FavoriteWines = append(user.FavoriteWines, wine)
	}

	reviewsQuery := `SELECT id, wine_id, winemaker, wine_name, comment, review_date, review_date_time, 
                            review_date_time_utc, title, description, rating 
                     FROM reviews WHERE user_id = $1`
	rows, err = r.DB.Query(reviewsQuery, user.ID)
	if err != nil {
		return models.User{}, err
	}
	defer rows.Close()
	for rows.Next() {
		var review models.Review
		err := rows.Scan(
			&review.ID,
			&review.WineID,
			&review.Content,
			&review.ReviewDate,
			&review.ReviewDateTime,
			&review.Title,
			&review.Rating,
		)
		if err != nil {
			return models.User{}, err
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
			&review.Content,
			&review.ReviewDate,
			&review.ReviewDateTime,
			&review.Title,
			&review.Rating,
		)
		if err != nil {
			return nil, err
		}
		reviews = append(reviews, review)
	}
	return reviews, nil
}




