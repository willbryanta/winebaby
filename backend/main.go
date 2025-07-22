package main

import (
	"database/sql"
	"flag"
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
	log.Println("Connecting to database with connection string:", connStr)
	if err != nil {
		log.Fatal("Error opening database: ", err)
		return nil, err
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)

	return db, nil
}

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received %s request for %s from %s", r.Method, r.URL.Path, r.RemoteAddr)
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Accept")
		w.Header().Set("Access-Control-Allow-Credentials", "true")


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

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, falling back to environment variables")
	}

	dbConn, err := connectDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer dbConn.Close()

	log.Printf("Database connection stats: %+v", dbConn.Stats())
	log.Println("Successfully connected to winebaby_db!")

	if *seed {
		log.Println("Seeding the database with sample data...")
		if err := db.Seed(dbConn); err != nil {
			log.Printf("Error seeding the database: %v", err)
			return
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

	log.Printf("Server running on port %s", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}