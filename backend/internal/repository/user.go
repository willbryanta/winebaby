package repository

import (
	"database/sql"
	"winebaby/internal/models"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}
func (r *Repository) CreateUser(user models.User) error {
	query := `INSERT INTO users (username, password, email) VALUES ($1, $2, $3) RETURNING id`
	err := r.db.QueryRow(query, user.Username, user.Password, user.Email).Scan(&user.ID)
	return err
}
func (r *Repository) GetUserByUsername(username string) (models.User, error) {
	query := `SELECT id, username, password, email FROM users WHERE username = $1`
	var user models.User
	err := r.db.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, nil
		}
		return models.User{}, err
	}
	return user, nil
}

func (r *Repository) GetUserProfile(username string) (models.UserProfile, error) {

	userQuery := `SELECT id, username, email FROM users WHERE username = $1`
	var user models.UserProfile
	err := r.db.QueryRow(userQuery, username).Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		return models.UserProfile{}, err
	}

	winesQuery := `SELECT id, name, region FROM favorite_wines WHERE user_id = $1`
	rows, err := r.db.Query(winesQuery, user.ID)
	if err != nil {
		return models.UserProfile{}, err
	}
	defer rows.Close()
	for rows.Next() {
		var wine models.Wine
		if err := rows.Scan(&wine.ID, &wine.Name, &wine.Region); err != nil {
			return models.UserProfile{}, err
		}
		user.FavoriteWines = append(user.FavoriteWines, wine)
	}
	reviewsQuery := `SELECT id, wine_name, rating, comment FROM reviews WHERE user_id = $1`
	rows, err = r.db.Query(reviewsQuery, user.ID)
	if err != nil {
		return models.UserProfile{}, err
	}
	defer rows.Close()
	for rows.Next() {
		var review models.Review
		if err := rows.Scan(&review.ID, &review.WineName, &review.Rating, &review.Comment); err != nil {
			return models.UserProfile{}, err
		}
		user.Reviews = append(user.Reviews, review)
	}
	return user, nil
}
