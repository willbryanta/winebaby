package models

type Review struct {
	ID				int 	`json:"id"`
	UserID			int 	`json:"user_id"`
	WineID 			int		`json:"wine_id"`
	Content			string 	`json:"comment"`
	ReviewDate		string 	`json:"review_date"`
	ReviewDateTime	string 	`json:"review_date_time"`
	Title			string 	`json:"title"`
	Rating			int 	`json:"rating"`
}

