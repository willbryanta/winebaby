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