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
	ServingSizeUnitAbbreviation	string	`json:"serving_size_unit_abbreviation"`
	ServingSizeUnitDescription	string	`json:"serving_size_unit_description"`
	ServingSizeUnitDescriptionAbbreviation	string	`json:"serving_size_unit_description_abbreviation"`
	ServingSizeUnitDescriptionPlural	string	`json:"serving_size_unit_description_plural"`
	Price			float32	`json:"price"`
	Rating			float32	`json:"rating"`
	Reviews			[]Review	`json:"reviews"`
	ReviewCount		int		`json:"review_count"`
	AverageRating	float32	`json:"average_rating"`
	Review			Review	`json:"review"`
	Type			string 	`json:"type"` // Type refers to cab sav/etc
	Colour			string 	`json:"colour"` // Red or White
}

type WineProfile struct {
	ID				int 	`json:"id"`
	Title			string 	`json:"title"`
	Year			int		`json:"date"` // May want to change this if using GORM to grab a date type
	Manufacturer	string 	`json:"manufacturer"`
	Type			string 	`json:"type"` // Type refers to cab sav/etc
	Colour			string 	`json:"colour"` // Red or White
	Reviews			[]Review	`json:"reviews"`
	Rating			float32	`json:"rating"`
	Price			float32	`json:"price"`
	Region			string	`json:"region"`
	AlcoholContent	float32	`json:"alcohol_content"`
	ServingTemp	float32	`json:"serving_temp"`
	ServingSize	float32	`json:"serving_size"`
	ServingSizeUnit	string	`json:"serving_size_unit"`
	ServingSizeUnitAbbreviation	string	`json:"serving_size_unit_abbreviation"`
	ServingSizeUnitDescription	string	`json:"serving_size_unit_description"`
	ServingSizeUnitDescriptionAbbreviation	string	`json:"serving_size_unit_description_abbreviation"`
	ServingSizeUnitDescriptionPlural	string	`json:"serving_size_unit_description_plural"`
}