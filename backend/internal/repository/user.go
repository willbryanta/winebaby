package repository

import (
	"database/sql"
	"winebaby/internal/models"
)

type Repository struct {
	db *sql.DB
}

func (r *Repository) SignIn(username, password string) (models.User, error) {
	query := `SELECT id, username, email, password FROM users WHERE username = $1 AND password = $2`
	var user models.User
	var email sql.NullString
	err := r.db.QueryRow(query, username, password).Scan(&user.ID, &user.Username, &email, &user.Password)
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

func (r *Repository) SignUp(username, password, email string) (models.User, error) {
	query := `INSERT INTO users (username, password, email) VALUES ($1, $2, $3) RETURNING id`
	var user models.User
	var id int
	var emailPtr *string
	if email != "" {
		emailPtr = &email
	}
	err := r.db.QueryRow(query, username, password, emailPtr).Scan(&id)
	if err != nil {
		return models.User{}, err
	}
	user.ID = id
	user.Username = username
	user.Password = password
	user.Email = emailPtr
	return user, nil
}

func (r *Repository) CreateUser(user models.User) error {
	query := `INSERT INTO users (username, password, email) VALUES ($1, $2, $3) RETURNING id`
	err := r.db.QueryRow(query, user.Username, user.Password, user.Email).Scan(&user.ID)
	return err
}

func (r *Repository) GetUserByUsername(username string) (models.User, error) {
	query := `SELECT id, username, email, password FROM users WHERE username = $1`
	var user models.User
	var email sql.NullString
	err := r.db.QueryRow(query, username).Scan(&user.ID, &user.Username, &email, &user.Password)
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

func (r *Repository) GetUserProfile(username string) (models.UserProfile, error) {
	userQuery := `SELECT id, username, email FROM users WHERE username = $1`
	var user models.UserProfile
	var email sql.NullString
	err := r.db.QueryRow(userQuery, username).Scan(&user.ID, &user.Username, &email)
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
	rows, err := r.db.Query(winesQuery, user.ID)
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
			&wine.ServingTemp,
			&wine.ServingSize,
			&wine.ServingSizeUnit,
			&wine.ServingSizeUnitAbbreviation,
			&wine.ServingSizeUnitDescription,
			&wine.ServingSizeUnitDescriptionAbbreviation,
			&wine.ServingSizeUnitDescriptionPlural,
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
	rows, err = r.db.Query(reviewsQuery, user.ID)
	if err != nil {
		return models.UserProfile{}, err
	}
	defer rows.Close()
	for rows.Next() {
		var review models.Review
		err := rows.Scan(
			&review.ID,
			&review.WineID,
			&review.Winemaker,
			&review.WineName,
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
func (r *Repository) GetUserByEmail(email string) (models.User, error) {
	query := `SELECT id, username, email, password FROM users WHERE email = $1`
	var user models.User
	var emailVal sql.NullString
	err := r.db.QueryRow(query, email).Scan(&user.ID, &user.Username, &emailVal, &user.Password)
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

func (r *Repository) UpdateUserProfile(user models.User) error {
	query := `UPDATE users SET username = $1, email = $2 WHERE id = $3`
	_, err := r.db.Exec(query, user.Username, user.Email, user.ID)
	return err
}

func (r *Repository) DeleteUser(username string) error {
	query := `DELETE FROM users WHERE username = $1`
	_, err := r.db.Exec(query, username)
	return err
}

func (r *Repository) GetUserFavoriteWines(username string) ([]models.Wine, error) {
	query := `SELECT id, name, year, manufacturer, region, alcohol_content, serving_temp, serving_size, 
						  serving_size_unit, serving_size_unit_abbreviation, serving_size_unit_description, 
						  serving_size_unit_description_abbreviation, serving_size_unit_description_plural, 
						  price, rating, type, colour 
				   FROM favorite_wines WHERE user_id = (SELECT id FROM users WHERE username = $1)`
	rows, err := r.db.Query(query, username)
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
			&wine.ServingTemp,
			&wine.ServingSize,
			&wine.ServingSizeUnit,
			&wine.ServingSizeUnitAbbreviation,
			&wine.ServingSizeUnitDescription,
			&wine.ServingSizeUnitDescriptionAbbreviation,
			&wine.ServingSizeUnitDescriptionPlural,
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

func (r *Repository) AddUserFavoriteWine(userID int, wineID int) error {
	query := `INSERT INTO favorite_wines (user_id, wine_id) VALUES ($1, $2)`
	_, err := r.db.Exec(query, userID, wineID)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) RemoveUserFavoriteWine(userID int, wineID int) error {
	query := `DELETE FROM favorite_wines WHERE user_id = $1 AND wine_id = $2`
	_, err := r.db.Exec(query, userID, wineID)
	return err
}

func (r *Repository) GetUserReviews(username string) ([]models.Review, error) {
	query := `SELECT id, wine_id, winemaker, wine_name, comment, review_date, review_date_time, 
						  review_date_time_utc, title, description, rating 
				   FROM reviews WHERE user_id = (SELECT id FROM users WHERE username = $1)`
	rows, err := r.db.Query(query, username)
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
			&review.Winemaker,
			&review.WineName,
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

func (r *Repository) CreateUserReview(review models.Review) error {
	query := `INSERT INTO reviews (user_id, wine_id, winemaker, wine_name, comment, review_date, 
								  review_date_time, review_date_time_utc, title, description, rating) 
			  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id`
	err := r.db.QueryRow(query,
		review.UserID,
		review.WineID,
		review.Winemaker,
		review.WineName,
		review.Comment,
		review.ReviewDate,
		review.ReviewDateTime,
		review.ReviewDateTimeUTC,
		review.Title,
		review.Description,
		review.Rating).Scan(&review.ID)
	return err
}