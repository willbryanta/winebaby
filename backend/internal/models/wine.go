package models

type Wine struct {
	ID				int 	`json:"id"`
	Name			string 	`json:"name"`
	Year			int		`json:"date"` // May want to change this if using GORM to grab a date type
	Manufacturer	string 	`json:"manufacturer"`
	Region			string 	`json:"region"`
	AlcoholContent	float32	`json:"alcohol_content"`
	ServingTemp	float32	`json:"serving_temp"`
	ServingSize	float32	`json:"serving_size"`
	ServingSizeUnit	string	`json:"serving_size_unit"`
	Price			float32	`json:"price"`
	Rating			float32	`json:"rating"`
	Reviews			[]Review	`json:"reviews"`
	ReviewCount		int		`json:"review_count"`
	AverageRating	float32	`json:"average_rating"`
	Type			string 	`json:"type"` // Type refers to cab sav/etc
	Colour			string 	`json:"colour"` // Red or White
}
