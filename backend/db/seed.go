package db

import (
	"database/sql"
	"log"
)

func Seed(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()
	wineStmt, err := tx.Prepare(`
		INSERT INTO wines (
			id, name, year, manufacturer, region, alcohol_content, price,
			type, colour, image_url, review_count, average_rating
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
		ON CONFLICT DO NOTHING
	`)
	if err != nil {
		return err
	}
	defer wineStmt.Close()

	reviewStmt, err := tx.Prepare(`
		INSERT INTO reviews (
			id, wine_id, content, review_date, review_date_time, title, rating
		) VALUES ($1, $2, $3, $4, $5, $6, $7)
		ON CONFLICT DO NOTHING
	`)
	if err != nil {
		return err
	}
	defer reviewStmt.Close()

		type review struct {
		id             int
		wineID         int
		content        string
		reviewDate     string
		reviewDateTime string
		title          string
		rating         int
	}

	type wine struct {
		id             int
		name           string
		year           int
		manufacturer   string
		region         string
		alcoholContent float64
		price          float64
		typ            string
		colour          string
		imageURL       string
		reviewCount    int
		averageRating  interface{}
		reviews        []review
	}

	wines := []wine{
		{
			id:             1,
			name:           "Château Lafite Rothschild",
			year:           2019,
			manufacturer:   "Domaines Barons de Rothschild",
			region:         "Pauillac, Bordeaux, France",
			alcoholContent: 13.5,
			price:          650.0,
			typ:            "Cabernet Sauvignon",
			colour:          "Red",
			imageURL:       "https://media.danmurphys.com.au/dmo/product/73984-1.png?impolicy=PROD_SM",
			reviewCount:    2,
			averageRating:  4.5,
			reviews: []review{
				{
					id:             1,
					wineID:         1,
					content:        "Exquisite balance of fruit and tannins, a true masterpiece.",
					reviewDate:     "2024-05-10",
					reviewDateTime: "2024-05-10T10:30:00Z",
					title:          "Exceptional!",
					rating:         5,
				},
				{
					id:             2,
					wineID:         1,
					content:        "Very refined, but needs more time to mature.",
					reviewDate:     "2024-06-01",
					reviewDateTime: "2024-06-01T15:45:00Z",
					title:          "Great Potential",
					rating:         4,
				},
			},
		},
		{
			id:             2,
			name:           "Screaming Eagle Sauvignon Blanc",
			year:           2021,
			manufacturer:   "Screaming Eagle",
			region:         "Napa Valley, California, USA",
			alcoholContent: 13.0,
			price:          75.0,
			typ:            "Sauvignon Blanc",
			colour:          "White",
			imageURL:       "https://media.danmurphys.com.au/dmo/product/259651-1.png?impolicy=PROD_SM",
			reviewCount:    1,
			averageRating:  4.0,
			reviews: []review{
				{
					id:             3,
					wineID:         2,
					content:        "Crisp and vibrant, perfect for a summer evening.",
					reviewDate:     "2024-07-15",
					reviewDateTime: "2024-07-15T18:20:00Z",
					title:          "Refreshing Delight",
					rating:         4,
				},
			},
		},
		{
			id:             3,
			name:           "Penfolds Grange",
			year:           2017,
			manufacturer:   "Penfolds",
			region:         "South Australia",
			alcoholContent: 14.5,
			price:          800.0,
			typ:            "Shiraz",
			colour:          "Red",
			imageURL:       "https://media.danmurphys.com.au/dmo/product/261083-1.png?impolicy=PROD_SM",
			reviewCount:    2,
			averageRating:  4.0,
			reviews: []review{
				{
					id:             4,
					wineID:         3,
					content:        "Bold and complex, with a long finish.",
					reviewDate:     "2024-08-20",
					reviewDateTime: "2024-08-20T09:00:00Z",
					title:          "Iconic Aussie Shiraz",
					rating:         5,
				},
				{
					id:             5,
					wineID:         3,
					content:        "Too intense for my taste, but well-crafted.",
					reviewDate:     "2024-09-05",
					reviewDateTime: "2024-09-05T14:10:00Z",
					title:          "Powerful",
					rating:         3,
				},
			},
		},
		{
			id:             4,
			name:           "Dom Pérignon Vintage",
			year:           2012,
			manufacturer:   "Moët & Chandon",
			region:         "Champagne, France",
			alcoholContent: 12.5,
			price:          200.0,
			typ:            "Champagne",
			colour:          "Sparkling",
			imageURL:       "https://media.danmurphys.com.au/dmo/product/6699-1.png?impolicy=PROD_SM",
			reviewCount:    0,
			averageRating:  nil,
			reviews:        []review{},
		},
		{
			id:             5,
			name:           "Cloudy Bay Chardonnay",
			year:           2022,
			manufacturer:   "Cloudy Bay Vineyards",
			region:         "Marlborough, New Zealand",
			alcoholContent: 13.5,
			price:          45.0,
			typ:            "Chardonnay",
			colour:          "White",
			imageURL:       "https://media.danmurphys.com.au/dmo/product/462559-1.png?impolicy=PROD_SM",
			reviewCount:    0,
			averageRating:  nil,
			reviews:        []review{},
		},
		{
			id:             6,
			name:           "Reddoch",
			year:           2018,
			manufacturer:   "Opus One Winery",
			region:         "Napa Valley, California, USA",
			alcoholContent: 14.0,
			price:          350.0,
			typ:            "Cabernet Sauvignon Blend",
			colour:          "Red",
			imageURL:       "https://media.danmurphys.com.au/dmo/product/741918-1.png?impolicy=PROD_SM",
			reviewCount:    2,
			averageRating:  4.5,
			reviews: []review{
				{
					id:             6,
					wineID:         6,
					content:        "Rich and velvety, with deep fruit flavors.",
					reviewDate:     "2024-10-01",
					reviewDateTime: "2024-10-01T12:00:00Z",
					title:          "Luxurious Blend",
					rating:         5,
				},
				{
					id:             7,
					wineID:         6,
					content:        "Well-balanced but slightly overpriced.",
					reviewDate:     "2024-10-15",
					reviewDateTime: "2024-10-15T09:30:00Z",
					title:          "Solid but Pricey",
					rating:         4,
				},
			},
		},
		{
			id:             7,
			name:           "Little Giant",
			year:           2016,
			manufacturer:   "Château Margaux",
			region:         "Margaux, Bordeaux, France",
			alcoholContent: 13.5,
			price:          600.0,
			typ:            "Barossa Shiraz",
			colour:          "Red",
			imageURL:       "https://media.danmurphys.com.au/dmo/product/686166-1.png?impolicy=PROD_SM",
			reviewCount:    1,
			averageRating:  5.0,
			reviews: []review{
				{
					id:             8,
					wineID:         7,
					content:        "Elegant and structured, a classic Bordeaux.",
					reviewDate:     "2024-11-05",
					reviewDateTime: "2024-11-05T14:20:00Z",
					title:          "Timeless Elegance",
					rating:         5,
				},
			},
		},
		{
			id:             8,
			name:           "Sassicaia",
			year:           2019,
			manufacturer:   "Tenuta San Guido",
			region:         "Tuscany, Italy",
			alcoholContent: 14.0,
			price:          250.0,
			typ:            "Cabernet Sauvignon Blend",
			colour:          "Red",
			imageURL:       "https://media.danmurphys.com.au/dmo/product/29480-1.png?impolicy=PROD_SM",
			reviewCount:    2,
			averageRating:  4.0,
			reviews: []review{
				{
					id:             9,
					wineID:         8,
					content:        "Smooth and complex, with hints of blackberry.",
					reviewDate:     "2024-12-10",
					reviewDateTime: "2024-12-10T16:45:00Z",
					title:          "Italian Gem",
					rating:         4,
				},
				{
					id:             10,
					wineID:         8,
					content:        "Needs time to open up, but promising.",
					reviewDate:     "2024-12-15",
					reviewDateTime: "2024-12-15T11:00:00Z",
					title:          "Good Potential",
					rating:         4,
				},
			},
		},
		{
			id:             9,
			name:           "Kim Crawford Sauvignon Blanc",
			year:           2023,
			manufacturer:   "Kim Crawford Wines",
			region:         "Marlborough, New Zealand",
			alcoholContent: 12.5,
			price:          20.0,
			typ:            "Sauvignon Blanc",
			colour:          "White",
			imageURL:       "https://media.danmurphys.com.au/dmo/product/608178-1.png?impolicy=PROD_SM",
			reviewCount:    1,
			averageRating:  4.0,
			reviews: []review{
				{
					id:             11,
					wineID:         9,
					content:        "Zesty and refreshing, great value.",
					reviewDate:     "2025-01-05",
					reviewDateTime: "2025-01-05T17:30:00Z",
					title:          "Crisp and Affordable",
					rating:         4,
				},
			},
		},
		{
			id:             10,
			name:           "Vega Sicilia Unico",
			year:           2015,
			manufacturer:   "Vega Sicilia",
			region:         "Ribera del Duero, Spain",
			alcoholContent: 14.0,
			price:          400.0,
			typ:            "Tempranillo",
			colour:          "Red",
			imageURL:       "https://media.danmurphys.com.au/dmo/product/907928-1.png?impolicy=PROD_SM",
			reviewCount:    3,
			averageRating:  4.7,
			reviews: []review{
				{
					id:             12,
					wineID:         10,
					content:        "Powerful yet refined, a Spanish masterpiece.",
					reviewDate:     "2025-02-01",
					reviewDateTime: "2025-02-01T10:15:00Z",
					title:          "Stunning Complexity",
					rating:         5,
				},
				{
					id:             13,
					wineID:         10,
					content:        "Bold flavors, but a bit tannic.",
					reviewDate:     "2025-02-10",
					reviewDateTime: "2025-02-10T13:50:00Z",
					title:          "Needs Time",
					rating:         4,
				},
				{
					id:             14,
					wineID:         10,
					content:        "Absolutely divine, worth every penny.",
					reviewDate:     "2025-02-15",
					reviewDateTime: "2025-02-15T08:00:00Z",
					title:          "Exceptional",
					rating:         5,
				},
			},
		},
		{
			id:             11,
			name:           "Rombauer Chardonnay",
			year:           2022,
			manufacturer:   "Rombauer Vineyards",
			region:         "Carneros, California, USA",
			alcoholContent: 14.5,
			price:          40.0,
			typ:            "Chardonnay",
			colour:          "White",
			imageURL:       "https://media.danmurphys.com.au/dmo/product/260252-1.png?impolicy=PROD_SM",
			reviewCount:    1,
			averageRating:  4.0,
			reviews: []review{
				{
					id:             15,
					wineID:         11,
					content:        "Buttery and smooth, a crowd-pleaser.",
					reviewDate:     "2025-03-01",
					reviewDateTime: "2025-03-01T19:00:00Z",
					title:          "Rich and Creamy",
					rating:         4,
				},
			},
		},
		{
			id:             12,
			name:           "Château d’Yquem",
			year:           2018,
			manufacturer:   "Château d’Yquem",
			region:         "Sauternes, Bordeaux, France",
			alcoholContent: 13.5,
			price:          300.0,
			typ:            "Sémillon",
			colour:          "White",
			imageURL:       "https://media.danmurphys.com.au/dmo/product/793361-1.png?impolicy=PROD_SM",
			reviewCount:    2,
			averageRating:  4.0,
			reviews: []review{
				{
					id:             16,
					wineID:         12,
					content:        "Luscious and sweet, with perfect balance.",
					reviewDate:     "2025-04-10",
					reviewDateTime: "2025-04-10T14:30:00Z",
					title:          "Dessert Perfection",
					rating:         5,
				},
				{
					id:             17,
					wineID:         12,
					content:        "Too sweet for my palate, but well-made.",
					reviewDate:     "2025-04-15",
					reviewDateTime: "2025-04-15T09:20:00Z",
					title:          "Very Sweet",
					rating:         3,
				},
			},
		},
	}

	// Insert wines
	for _, w := range wines {
		_, err := wineStmt.Exec(
			w.id, w.name, w.year, w.manufacturer, w.region, w.alcoholContent,
			w.price, w.typ, w.colour, w.imageURL, w.reviewCount, w.averageRating,
		)
		if err != nil {
			log.Printf("Error inserting wine %s: %v", w.name, err)
			return err
		}
	}

	// Insert reviews
	for _, w := range wines {
		for _, r := range w.reviews {
			_, err := reviewStmt.Exec(
				r.id, r.wineID, r.content, r.reviewDate, r.reviewDateTime, r.title, r.rating,
			)
			if err != nil {
				log.Printf("Error inserting review for wine ID %d: %v", r.wineID, err)
				return err
			}
		}
	}

	log.Println("Database seeded successfully with wines and reviews")
	return nil
}