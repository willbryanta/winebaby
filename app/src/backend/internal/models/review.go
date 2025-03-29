package models

type Review struct {
	ID				int 	`json:"id"`
	WineID 			int		`json:"wine_id"`
	Manufacturer	string	`json:"maker"`
	Title			string 	`json:"title"`
	Rating			int 	`json:"rating"`
}