package handlers 

type Wine struct {
	ID				int 	`json:"id"`
	Title			string 	`json:"title"`
	Year			int		`json:"date"` // May want to change this if using GORM to grab a date type
	Manufacturer	string 	`json:"manufacturer`
	Type			string 	`json:"type"` // Type refers to cab sav/etc
	Colour			string 	`json:"colour"` // Red or White
}