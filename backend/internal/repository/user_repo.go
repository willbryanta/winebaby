package repository

import (
	"database/sql"
	"winebaby/internal/models"
)

func CreateUser(user models.User) error {
	db, err := sql.Open("postgres", "postgres://user:password@localhost:5432/winebaby_db?sslmode=disable")
	if err != nil {
		return err
	}
	defer db.Close()

	query := `INSERT INTO users (username, password, email) VALUES ($1, $2, $3) RETURNING id`
	err = db.QueryRow(query, user.Username, user.Password, user.Email).Scan(&user.ID)
	return err
}

func GetUserByUsername(username string) (models.User, error) {
	db, err := sql.Open("postgres", "postgres://user:password@localhost:5432/winebaby_db?sslmode=disable")
	if err != nil {
		return models.User{}, err
	}
	defer db.Close()
	query := `SELECT id, username, password, email FROM users WHERE username = $1`
	var user models.User
	err = db.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, nil
		}
		return models.User{}, err
	}
	return user, nil
}