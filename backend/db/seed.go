package db

import (
	"database/sql"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func Seed(db *sql.DB) error {
	// Start a transaction to ensure atomicity
	tx, err := db.Begin()
	if err != nil {
		log.Printf("Failed to start transaction: %v", err)
		return err
	}
	defer func() {
		if err != nil {
			log.Printf("Rolling back transaction due to error: %v", err)
			tx.Rollback()
			return
		}
		err = tx.Commit()
		if err != nil {
			log.Printf("Failed to commit transaction: %v", err)
		}
	}()

	// Seed users
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Failed to hash password: %v", err)
		return err
	}

	users := []struct {
		username, password, email string
	}{
		{"user1", string(hashedPassword), "user1@example.com"},
		{"user2", string(hashedPassword), "user2@example.com"},
	}
	for _, u := range users {
		log.Printf("Inserting user: %s", u.username)
		_, err := tx.Exec(`INSERT INTO users (username, password, email) VALUES ($1, $2, $3) ON CONFLICT DO NOTHING`,
			u.username, u.password, u.email)
		if err != nil {
			log.Printf("Failed to insert user %s: %v", u.username, err)
			return err
		}
	}

	// Seed wines
	wines := []struct {
		name, manufacturer, region, wineType, colour string
		year                                    int
		alcoholContent, servingTemp, servingSize, price, rating float32
		servingSizeUnit, servingSizeUnitAbbreviation, servingSizeUnitDescription string
		servingSizeUnitDescriptionAbbreviation, servingSizeUnitDescriptionPlural string
	}{
		{
			name: "McClaren Vale Shiraz", manufacturer: "Penfolds", region: "Australia", wineType: "Red", colour: "Deep Red",
			year: 2024, alcoholContent: 14.5, servingTemp: 18.0, servingSize: 150.0, price: 29.99, rating: 4.5,
			servingSizeUnit: "Milliliter", servingSizeUnitAbbreviation: "mL",
			servingSizeUnitDescription: "Glass", servingSizeUnitDescriptionAbbreviation: "gl",
			servingSizeUnitDescriptionPlural: "Glasses",
		},
		{
			name: "Chablis", manufacturer: "Domaine Laroche", region: "France", wineType: "White", colour: "Pale Gold",
			year: 2023, alcoholContent: 12.5, servingTemp: 10.0, servingSize: 125.0, price: 24.99, rating: 4.0,
			servingSizeUnit: "Milliliter", servingSizeUnitAbbreviation: "mL",
			servingSizeUnitDescription: "Glass", servingSizeUnitDescriptionAbbreviation: "gl",
			servingSizeUnitDescriptionPlural: "Glasses",
		},
	}
	for _, w := range wines {
		log.Printf("Inserting wine: %s", w.name)
		_, err := tx.Exec(`INSERT INTO wines (name, year, manufacturer, region, alcohol_content, serving_temp, 
                                            serving_size, serving_size_unit, serving_size_unit_abbreviation, 
                                            serving_size_unit_description, serving_size_unit_description_abbreviation, 
                                            serving_size_unit_description_plural, price, rating, type, colour) 
                          VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16) 
                          ON CONFLICT DO NOTHING`,
			w.name, w.year, w.manufacturer, w.region, w.alcoholContent, w.servingTemp, w.servingSize,
			w.servingSizeUnit, w.servingSizeUnitAbbreviation, w.servingSizeUnitDescription,
			w.servingSizeUnitDescriptionAbbreviation, w.servingSizeUnitDescriptionPlural,
			w.price, w.rating, w.wineType, w.colour)
		if err != nil {
			log.Printf("Failed to insert wine %s: %v", w.name, err)
			return err
		}
	}

	// Seed favorite_wines
	favoriteWines := []struct {
		username string
		wineID   int
	}{
		{"user1", 1},
		{"user1", 2},
		{"user2", 1},
	}
	for _, fw := range favoriteWines {
		log.Printf("Inserting favorite wine for user %s, wine ID %d", fw.username, fw.wineID)
		_, err := tx.Exec(`INSERT INTO favorite_wines (user_id, wine_id) 
                          SELECT id, $2 FROM users WHERE username = $1 
                          ON CONFLICT DO NOTHING`,
			fw.username, fw.wineID)
		if err != nil {
			log.Printf("Failed to insert favorite wine for user %s, wine ID %d: %v", fw.username, fw.wineID, err)
			return err
		}
	}

	// Seed reviews
	reviews := []struct {
		username, comment, reviewDate, reviewDateTime, reviewDateTimeUTC, title, description string
		userID, wineID, rating                                                      int
	}{
		{
			username: "user1", userID: 1, wineID: 1, comment: "Great wine!", title: "McClaren Vale 2024",
			description: "Zesty and fruity with notes of licorice", rating: 7,
			reviewDate: time.Now().Format("2006-01-02"),
			reviewDateTime: time.Now().Format("2006-01-02 15:04:05"),
			reviewDateTimeUTC: time.Now().UTC().Format("2006-01-02 15:04:05"),
		},
		{
			username: "user2", userID: 2, wineID: 2, comment: "Crisp and refreshing", title: "Chablis 2023",
			description: "Perfect for a summer evening", rating: 8,
			reviewDate: time.Now().Format("2006-01-02"),
			reviewDateTime: time.Now().Format("2006-01-02 15:04:05"),
			reviewDateTimeUTC: time.Now().UTC().Format("2006-01-02 15:04:05"),
		},
	}
	for _, r := range reviews {
		log.Printf("Inserting review for user ID %d, wine ID %d", r.userID, r.wineID)
		_, err := tx.Exec(`INSERT INTO reviews (user_id, wine_id, comment, review_date, review_date_time, 
                                              review_date_time_utc, title, description, rating) 
                          VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) ON CONFLICT DO NOTHING`,
			r.userID, r.wineID, r.comment, r.reviewDate, r.reviewDateTime, r.reviewDateTimeUTC,
			r.title, r.description, r.rating)
		if err != nil {
			log.Printf("Failed to insert review for user ID %d, wine ID %d: %v", r.userID, r.wineID, err)
			return err
		}
	}

	log.Println("Database seeding completed successfully")
	return nil
}