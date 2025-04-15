package models

type Review struct {
	ID				int 	`json:"id"`
	UserID			int 	`json:"user_id"`
	WineID 			int		`json:"wine_id"`
	Winemaker		string	`json:"winemaker"`
	WineName		string	`json:"wine_name"`
	Comment			string 	`json:"comment"`
	ReviewDate		string 	`json:"review_date"`
	ReviewDateTime	string 	`json:"review_date_time"`
	ReviewDateTimeUTC	string 	`json:"review_date_time_utc"`
	Title			string 	`json:"title"`
	Description 	string 	`json:"description"`
	Rating			int 	`json:"rating"`
}

