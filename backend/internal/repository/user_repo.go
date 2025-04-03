package repository

import "winebaby/internal/models"

var users []models.User

func CreateUser(u models.User) error {
	if len(users) > 0 {
		u.ID = users[len(users)-1].ID + 1
	} else {
		u.ID = 1
	}
	users = append(users, u)
	return nil
}