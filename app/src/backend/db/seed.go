package db

import (
	"database/sql"
	"log"
	"winebaby/internal/models"
	_ "github.com/lib/pq"
)

func Seed(db *sql.DB) error {
	_, err := db.Exec("TRUNCATE TABLE wines RESTART IDENTITY")
	if err != nil {
		return err
	}

	wines := []models.Wine{
		{Title: "Chateau Margaux", Year: 2015, Manufacturer: "Chateau Margaux", Type: "Cabernet Savignon", Colour: "Red"},
		{Title: "Pinot Grigio Santa Margherita", Year: 2020, Manufacturer: "Santa Margherita", Type: "Pinot Grigio", Colour: "White"},
		{Title: "Opus One", Year: 2018, Manufacturer: "Opus One Winery", Type: "Bordeaux Blend", Colour: "Red"},
		{Title: "Sancerre", Year: 2021, Manufacturer: "Domaine Vacheron", Type: "Sauvignon Blanc", Colour: "White"},
	}

	stmt, err := db.Prepare(`
	INSERT INTO wines (title, year, manufacturer, type, colour)
	VALUES ($1, $2, $3, $4, $5)`)
	if err != nil {
		return err
	}
	defer stmt.Close()
	
	for _, wine := range wines {
		_, err = stmt.Exec(wine.Title, wine.Year, wine.Manufacturer, wine.Type, wine.Colour)
		if err != nil {
			return err
		}
	}
	log.Println("Database seeded successfully!")
	return nil
}