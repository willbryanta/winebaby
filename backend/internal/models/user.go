package models

type User struct {
	ID			int		`json:"id"`
	Username	string	`json:"username"`
	Email		*string	`json:"email"`
	Password	string	`json:"password"`
	FavoriteWines	[]Wine	`json:"favorite_wines"`
	Reviews		[]Review	`json:"reviews"`
}