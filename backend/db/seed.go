package db

import (
	"database/sql"
	"log"
	"winebaby/internal/models"

	_ "github.com/lib/pq"
)

func Seed(db *sql.DB) error {
	// Truncate the wines table and restart the ID sequence
	_, err := db.Exec("TRUNCATE TABLE wines RESTART IDENTITY CASCADE")
	if err != nil {
		return err
	}

	// Define sample wines based on the updated Wine model
	wines := []models.Wine{
		{
			Name:            "Chateau Margaux",
			Year:            2015,
			Manufacturer:    "Chateau Margaux",
			Region:          "Bordeaux",
			AlcoholContent:  13.5,
			ServingTemp:     18.0,
			ServingSize:     150.0,
			ServingSizeUnit: "ml",
			ServingSizeUnitAbbreviation:          "mL",
			ServingSizeUnitDescription:           "milliliter",
			ServingSizeUnitDescriptionAbbreviation: "mL",
			ServingSizeUnitDescriptionPlural:     "milliliters",
			Price:           250.00,
			Rating:          4.8,
			Type:            "Cabernet Sauvignon",
			Colour:          "Red",
		},
		{
			Name:            "Pinot Grigio Santa Margherita",
			Year:            2020,
			Manufacturer:    "Santa Margherita",
			Region:          "Veneto",
			AlcoholContent:  12.5,
			ServingTemp:     10.0,
			ServingSize:     150.0,
			ServingSizeUnit: "ml",
			ServingSizeUnitAbbreviation:          "mL",
			ServingSizeUnitDescription:           "milliliter",
			ServingSizeUnitDescriptionAbbreviation: "mL",
			ServingSizeUnitDescriptionPlural:     "milliliters",
			Price:           25.00,
			Rating:          4.2,
			Type:            "Pinot Grigio",
			Colour:          "White",
		},
		{
			Name:            "Opus One",
			Year:            2018,
			Manufacturer:    "Opus One Winery",
			Region:          "Napa Valley",
			AlcoholContent:  14.0,
			ServingTemp:     18.0,
			ServingSize:     150.0,
			ServingSizeUnit: "ml",
			ServingSizeUnitAbbreviation:          "mL",
			ServingSizeUnitDescription:           "milliliter",
			ServingSizeUnitDescriptionAbbreviation: "mL",
			ServingSizeUnitDescriptionPlural:     "milliliters",
			Price:           350.00,
			Rating:          4.9,
			Type:            "Bordeaux Blend",
			Colour:          "Red",
		},
		{
			Name:            "Sancerre",
			Year:            2021,
			Manufacturer:    "Domaine Vacheron",
			Region:          "Loire Valley",
			AlcoholContent:  13.0,
			ServingTemp:     12.0,
			ServingSize:     150.0,
			ServingSizeUnit: "ml",
			ServingSizeUnitAbbreviation:          "mL",
			ServingSizeUnitDescription:           "milliliter",
			ServingSizeUnitDescriptionAbbreviation: "mL",
			ServingSizeUnitDescriptionPlural:     "milliliters",
			Price:           40.00,
			Rating:          4.5,
			Type:            "Sauvignon Blanc",
			Colour:          "White",
		},
	}

	// Prepare the INSERT statement with all fields from the Wine model
	stmt, err := db.Prepare(`
		INSERT INTO wines (
			name, year, manufacturer, region, alcohol_content, serving_temp, 
			serving_size, serving_size_unit, serving_size_unit_abbreviation, 
			serving_size_unit_description, serving_size_unit_description_abbreviation, 
			serving_size_unit_description_plural, price, rating, type, colour
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Insert each wine
	for _, wine := range wines {
		_, err = stmt.Exec(
			wine.Name, wine.Year, wine.Manufacturer, wine.Region, wine.AlcoholContent,
			wine.ServingTemp, wine.ServingSize, wine.ServingSizeUnit,
			wine.ServingSizeUnitAbbreviation, wine.ServingSizeUnitDescription,
			wine.ServingSizeUnitDescriptionAbbreviation, wine.ServingSizeUnitDescriptionPlural,
			wine.Price, wine.Rating, wine.Type, wine.Colour,
		)
		if err != nil {
			return err
		}
	}

	log.Println("Database seeded successfully with wines!")
	return nil
}