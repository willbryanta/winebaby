package models

type User struct {
	ID			int		`json:"id"`
	Username	string	`json:"username"`
	Email		*string	`json:"email"`
	Password	string	`json:"password"`
}

type UserProfile struct {
	ID			int		`json:"id"`
	Username	string	`json:"username"`
	Email		*string	`json:"email"`
	FavoriteWines	[]Wine	`json:"favorite_wines"`
	Reviews		[]Review	`json:"reviews"`
}

