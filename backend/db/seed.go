package db

import (
	"database/sql"
	"log"
)

type review struct {
	id             int
	userID         int
	wineID         int
	content        string
	reviewDate     string
	reviewDateTime string
	title          string
	rating         int
}

type wine struct {
	name           string
	year           int
	manufacturer   string
	region         string
	alcoholContent float32
	price          float32
	rating         float32
	typ            string
	colour         string
	image_url      string
	reviewCount    int
	averageRating  sql.NullFloat64
	reviews        []review
}

type user struct {
	username string
	email    *string
	password string
}

func stringPtr(s string) *string {
	return &s
}

func Seed(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(`
		DROP TABLE IF EXISTS reviews CASCADE;
		DROP TABLE IF EXISTS user_favorite_wines CASCADE;
		DROP TABLE IF EXISTS wines CASCADE;
		DROP TABLE IF EXISTS users CASCADE;
	`)
	if err != nil {
		log.Println("Error dropping tables:", err)
		return err
	}
	_, err = tx.Exec(`
		CREATE TABLE users (
			id SERIAL PRIMARY KEY,
			username VARCHAR(255) NOT NULL UNIQUE,
			email VARCHAR(255),
			password VARCHAR(255) NOT NULL
		);
		CREATE TABLE wines (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			year INT NOT NULL,
			manufacturer VARCHAR(255) NOT NULL,
			region VARCHAR(255) NOT NULL,
			alcohol_content FLOAT NOT NULL,
			price FLOAT NOT NULL,
			rating FLOAT NOT NULL,
			type VARCHAR(50) NOT NULL,
			colour VARCHAR(50) NOT NULL,
			image_url VARCHAR(255),
			review_count INT DEFAULT 0,
			average_rating FLOAT DEFAULT 0.0,
			reviews JSONB DEFAULT '[]' NOT NULL
		);
		CREATE TABLE reviews (
			id SERIAL PRIMARY KEY,
			user_id INT NOT NULL REFERENCES users(id),
			wine_id INT NOT NULL REFERENCES wines(id),
			content TEXT NOT NULL,
			review_date DATE NOT NULL,
			review_date_time TIMESTAMP NOT NULL,
			title VARCHAR(255) NOT NULL,
			rating INT NOT NULL CHECK (rating >= 1 AND rating <= 5)
		);
		CREATE TABLE user_favorite_wines (
			user_id INT NOT NULL REFERENCES users(id),
			wine_id INT NOT NULL REFERENCES wines(id),
			PRIMARY KEY (user_id, wine_id)
		);
	`)
	if err != nil {
		log.Println("Error creating tables:", err)
		return err
	}

	userStmt, err := tx.Prepare(`
		INSERT INTO users (
			username, email, password
		) VALUES ($1, $2, $3)
		RETURNING id
	`)
	if err != nil {
		log.Println("Error preparing user statement:", err)
		return err
	}
	defer userStmt.Close()

	wineStmt, err := tx.Prepare(`
		INSERT INTO wines (
			name, year, manufacturer, region, alcohol_content, price,
			rating, type, colour, image_url, review_count, average_rating, reviews
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
		RETURNING id
	`)
	if err != nil {
		log.Println("Error preparing wine statement:", err)
		return err
	}
	defer wineStmt.Close()

	reviewStmt, err := tx.Prepare(`
		INSERT INTO reviews (
			user_id, wine_id, content, review_date, review_date_time, title, rating
		) VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id
	`)
	if err != nil {
		log.Println("Error preparing review statement:", err)
		return err
	}
	defer reviewStmt.Close()

	users := []user{
		{
			username: "user1",
			email:    stringPtr("user1@example.com"),
			password: "hashed_password_1", 
		},
		{
			username: "user2",
			email:    stringPtr("user2@example.com"),
			password: "hashed_password_2",
		},
		{
			username: "user3",
			email:    nil,
			password: "hashed_password_3",
		},
	}

	userIDs := make(map[int]int)
	for i, u := range users {
		var id int
		err := userStmt.QueryRow(
			u.username, u.email, u.password,
		).Scan(&id)
		if err != nil {
			log.Printf("Error inserting user %s: %v", u.username, err)
			return err
		}
		userIDs[i + 1] = id
	}

	wines := []wine{
		{
			name:           "Château Lafite Rothschild",
			year:           2019,
			manufacturer:   "Domaines Barons de Rothschild",
			region:         "Pauillac, Bordeaux, France",
			alcoholContent: 13.5,
			price:          650.0,
			rating:         4.5,
			typ:            "Cabernet Sauvignon",
			colour:         "Red",
			image_url:      "https://media.danmurphys.com.au/dmo/product/73984-1.png?impolicy=PROD_SM",
			reviewCount:    2,
			averageRating:  sql.NullFloat64{Float64: 4.5, Valid: true},
			reviews:        []review{},
		},
		{
			name:           "Screaming Eagle Sauvignon Blanc",
			year:           2021,
			manufacturer:   "Screaming Eagle",
			region:         "Napa Valley, California, USA",
			alcoholContent: 13.0,
			price:          75.0,
			rating:         4.0,
			typ:            "Sauvignon Blanc",
			colour:         "White",
			image_url:      "https://media.danmurphys.com.au/dmo/product/259651-1.png?impolicy=PROD_SM",
			reviewCount:    1,
			averageRating:  sql.NullFloat64{Float64: 4.0, Valid: true},
			reviews:        []review{},
		},
		{
			name:           "Penfolds Grange",
			year:           2017,
			manufacturer:   "Penfolds",
			region:         "South Australia",
			alcoholContent: 14.5,
			price:          800.0,
			rating:         4.0,
			typ:            "Shiraz",
			colour:         "Red",
			image_url:      "https://media.danmurphys.com.au/dmo/product/261083-1.png?impolicy=PROD_SM",
			reviewCount:    2,
			averageRating:  sql.NullFloat64{Float64: 4.0, Valid: true},
			reviews:        []review{},
		},
		{
			name:           "Dom Pérignon Vintage",
			year:           2012,
			manufacturer:   "Moët & Chandon",
			region:         "Champagne, France",
			alcoholContent: 12.5,
			price:          200.0,
			rating:         0.0,
			typ:            "Champagne",
			colour:         "Sparkling",
			image_url:      "https://media.danmurphys.com.au/dmo/product/6699-1.png?impolicy=PROD_SM",
			reviewCount:    0,
			averageRating:  sql.NullFloat64{Valid: false},
			reviews:        []review{},
		},
		{
			name:           "Cloudy Bay Chardonnay",
			year:           2022,
			manufacturer:   "Cloudy Bay Vineyards",
			region:         "Marlborough, New Zealand",
			alcoholContent: 13.5,
			price:          45.0,
			rating:         0.0,
			typ:            "Chardonnay",
			colour:         "White",
			image_url:      "https://media.danmurphys.com.au/dmo/product/462559-1.png?impolicy=PROD_SM",
			reviewCount:    0,
			averageRating:  sql.NullFloat64{Valid: false},
			reviews:        []review{},
		},
		{
			name:           "Opus One",
			year:           2018,
			manufacturer:   "Opus One Winery",
			region:         "Napa Valley, California, USA",
			alcoholContent: 14.0,
			price:          350.0,
			rating:         4.5,
			typ:            "Cabernet Sauvignon Blend",
			colour:         "Red",
			image_url:      "https://media.danmurphys.com.au/dmo/product/741918-1.png?impolicy=PROD_SM",
			reviewCount:    2,
			averageRating:  sql.NullFloat64{Float64: 4.5, Valid: true},
			reviews:        []review{},
		},
		{
			name:           "Château Margaux",
			year:           2016,
			manufacturer:   "Château Margaux",
			region:         "Margaux, Bordeaux, France",
			alcoholContent: 13.5,
			price:          600.0,
			rating:         5.0,
			typ:            "Cabernet Sauvignon",
			colour:         "Red",
			image_url:      "https://media.danmurphys.com.au/dmo/product/686166-1.png?impolicy=PROD_SM",
			reviewCount:    1,
			averageRating:  sql.NullFloat64{Float64: 5.0, Valid: true},
			reviews:        []review{},
		},
		{
			name:           "Sassicaia",
			year:           2019,
			manufacturer:   "Tenuta San Guido",
			region:         "Tuscany, Italy",
			alcoholContent: 14.0,
			price:          250.0,
			rating:         4.0,
			typ:            "Cabernet Sauvignon Blend",
			colour:         "Red",
			image_url:      "https://media.danmurphys.com.au/dmo/product/29480-1.png?impolicy=PROD_SM",
			reviewCount:    2,
			averageRating:  sql.NullFloat64{Float64: 4.0, Valid: true},
			reviews:        []review{},
		},
		{
			name:           "Kim Crawford Sauvignon Blanc",
			year:           2023,
			manufacturer:   "Kim Crawford Wines",
			region:         "Marlborough, New Zealand",
			alcoholContent: 12.5,
			price:          20.0,
			rating:         4.0,
			typ:            "Sauvignon Blanc",
			colour:         "White",
			image_url:      "https://media.danmurphys.com.au/dmo/product/608178-1.png?impolicy=PROD_SM",
			reviewCount:    1,
			averageRating:  sql.NullFloat64{Float64: 4.0, Valid: true},
			reviews:        []review{},
		},
		{
			name:           "Vega Sicilia Unico",
			year:           2015,
			manufacturer:   "Vega Sicilia",
			region:         "Ribera del Duero, Spain",
			alcoholContent: 14.0,
			price:          400.0,
			rating:         4.7,
			typ:            "Tempranillo",
			colour:         "Red",
			image_url:      "https://media.danmurphys.com.au/dmo/product/907928-1.png?impolicy=PROD_SM",
			reviewCount:    3,
			averageRating:  sql.NullFloat64{Float64: 4.7, Valid: true},
			reviews:        []review{},
		},
		{
			name:           "Rombauer Chardonnay",
			year:           2022,
			manufacturer:   "Rombauer Vineyards",
			region:         "Carneros, California, USA",
			alcoholContent: 14.5,
			price:          40.0,
			rating:         4.0,
			typ:            "Chardonnay",
			colour:         "White",
			image_url:      "https://media.danmurphys.com.au/dmo/product/260252-1.png?impolicy=PROD_SM",
			reviewCount:    1,
			averageRating:  sql.NullFloat64{Float64: 4.0, Valid: true},
			reviews:        []review{},
		},
		{
			name:           "Château d'Yquem",
			year:           2018,
			manufacturer:   "Château d'Yquem",
			region:         "Sauternes, Bordeaux, France",
			alcoholContent: 13.5,
			price:          300.0,
			rating:         4.0,
			typ:            "Sémillon",
			colour:         "White",
			image_url:      "https://media.danmurphys.com.au/dmo/product/793361-1.png?impolicy=PROD_SM",
			reviewCount:    2,
			averageRating:  sql.NullFloat64{Float64: 4.0, Valid: true},
			reviews:        []review{},
		},
	}

	wineIDs := make(map[int]int)
	for i, w := range wines {
		var id int
		err := wineStmt.QueryRow(
			i+1, w.name, w.year, w.manufacturer, w.region, w.alcoholContent,
			w.price, w.rating, w.typ, w.colour, w.image_url, w.reviewCount, w.averageRating,
		).Scan(&id)
		if err != nil {
			log.Printf("Error inserting wine %s: %v", w.name, err)
			return err
		}
		wineIDs[i+1] = id
	}

	reviews := []review{
		{       
			userID:         userIDs[1],
			wineID:         wineIDs[1],
			content:        "Exquisite balance of fruit and tannins, a true masterpiece.",
			reviewDate:     "2024-05-10",
			reviewDateTime: "2024-05-10T10:30:00Z",
			title:          "Exceptional!",
			rating:         5,
		},
		{

			userID:         userIDs[2],
			wineID:         wineIDs[1],
			content:        "Very refined, but needs more time to mature.",
			reviewDate:     "2024-06-01",
			reviewDateTime: "2024-06-01T15:45:00Z",
			title:          "Great Potential",
			rating:         4,
		},
		{
			userID:         userIDs[2],
			wineID:         wineIDs[2],
			content:        "Crisp and vibrant, perfect for a summer evening.",
			reviewDate:     "2024-07-15",
			reviewDateTime: "2024-07-15T18:20:00Z",
			title:          "Refreshing Delight",
			rating:         4,
		},
		{
			userID:         userIDs[1],
			wineID:         wineIDs[3],
			content:        "Bold and complex, with a long finish.",
			reviewDate:     "2024-08-20",
			reviewDateTime: "2024-08-20T09:00:00Z",
			title:          "Iconic Aussie Shiraz",
			rating:         5,
		},
		{
			userID:         userIDs[3],
			wineID:         wineIDs[3],
			content:        "Too intense for my taste, but well-crafted.",
			reviewDate:     "2024-09-05",
			reviewDateTime: "2024-09-05T14:10:00Z",
			title:          "Powerful",
			rating:         3,
		},
		{
			userID:         userIDs[3],
			wineID:         wineIDs[6],
			content:        "Rich and velvety, with deep fruit flavors.",
			reviewDate:     "2024-10-01",
			reviewDateTime: "2024-10-01T12:00:00Z",
			title:          "Luxurious Blend",
			rating:         5,
		},
		{
			userID:         userIDs[3],
			wineID:         wineIDs[6],
			content:        "Well-balanced but slightly overpriced.",
			reviewDate:     "2024-10-15",
			reviewDateTime: "2024-10-15T09:30:00Z",
			title:          "Solid but Pricey",
			rating:         4,
		},
	}

	for _, r := range reviews {
		_, err := reviewStmt.Exec(
			r.id, r.userID, r.wineID, r.content, r.reviewDate, r.reviewDateTime, r.title, r.rating,
		)
		if err != nil {
			log.Printf("Error inserting review for wine ID %d: %v", r.wineID, err)
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	log.Println("Database seeded successfully with users, wines, and reviews")
	return nil
	}