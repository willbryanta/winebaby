package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"winebaby/db"
	"winebaby/internal/routes"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
)

func connectDB() (*sql.DB, error){
	connStr := "postgres://"+os.Getenv("DB_USER") + ":" +
	os.Getenv("DB_PASSWORD") + "@" +
	os.Getenv("DB_HOST") + ":" +
	os.Getenv("DB_PORT") + "/" +
	os.Getenv("DB_NAME") + "?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error opening database: ", err)
		return nil, err
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	
	return db, nil
}

// TODO: Move this to a separate file potentially
func initSchema(db *sql.DB) error {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			username VARCHAR(50) UNIQUE NOT NULL,
			password TEXT NOT NULL,
			email VARCHAR(100) UNIQUE,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE TABLE IF NOT EXISTS wines (
			id SERIAL PRIMARY KEY,
			name VARCHAR(100) NOT NULL,
			year INT,
			manufacturer VARCHAR(100),
			region VARCHAR(100),
			alcohol_content FLOAT,
			serving_temp FLOAT,
			serving_size FLOAT,
			serving_size_unit VARCHAR(50),
			serving_size_unit_abbreviation VARCHAR(10),
			serving_size_unit_description VARCHAR(100),
			serving_size_unit_description_abbreviation VARCHAR(10),
			serving_size_unit_description_plural VARCHAR(100),
			price FLOAT,
			rating FLOAT,
			type VARCHAR(50),
			colour VARCHAR(50),
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE TABLE IF NOT EXISTS favorite_wines (
			user_id INT REFERENCES users(id) ON DELETE CASCADE,
			wine_id INT REFERENCES wines(id) ON DELETE CASCADE,
			PRIMARY KEY (user_id, wine_id),
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE TABLE IF NOT EXISTS reviews (
			id SERIAL PRIMARY KEY,
			user_id INT REFERENCES users(id) ON DELETE CASCADE,
			wine_id INT REFERENCES wines(id) ON DELETE CASCADE,
			comment TEXT,
			review_date VARCHAR(50),
			review_date_time VARCHAR(50),
			review_date_time_utc VARCHAR(50),
			title VARCHAR(100),
			description TEXT,
			rating INT,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`,
	}

	for _, query := range queries {
		_, err := db.Exec(query)
		if err != nil {
			return fmt.Errorf("error executing schema query: %w", err)
		}
	}
	return nil
}

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")


		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Println("Starting winebaby server...")

	seed := flag.Bool("seed", false, "Seed the database with sample data")
	flag.Parse()
	if *seed {
		log.Println("Seeding the database with sample data...")
	}

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, falling back to environment variables")
	}

	dbConn, err := connectDB()
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	log.Println(dbConn.Stats())
	defer dbConn.Close()

	if err = dbConn.Ping(); err != nil {
		log.Fatal("Error pinging the database: ", err)
	}
	log.Println("Successfully connected to winebaby_db!")

	if *seed {
		log.Println("Seeding the database with sample data...")
		if err := db.Seed(dbConn); err != nil {
			log.Fatal("Error seeding the database: ", err)
		}
		log.Println("Database seeded successfully!")
		return
	}

	r := chi.NewRouter()
	r.Use(CORSMiddleware) 
	r.Mount("/", routes.RegisterRoutes(dbConn))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server running on port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Server failed: ", err)
	}
}