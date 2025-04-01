package models

type Review struct {
	ID				int 	`json:"id"`
	WineID 			int		`json:"wine_id"`
	Winemaker		string	`json:"winemaker"`
	Title			string 	`json:"title"`
	Description 	string 	`json:"description"`
	Rating			int 	`json:"rating"`
}

