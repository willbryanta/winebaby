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

	log.Println("Server running on port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Server failed: ", err)
	}
}