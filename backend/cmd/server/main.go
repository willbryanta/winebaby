package main

import (
	"database/sql"
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
	return sql.Open("postgres", connStr)
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

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, falling back to environment variables")
	}

	dbConn, err := connectDB()
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	defer dbConn.Close()

	if err = dbConn.Ping(); err != nil {
		log.Fatal("Error pinging the database: ", err)
	}
	log.Println("Successfully connected to winebaby_db!")

	if err := db.Seed(dbConn); err != nil{
		log.Fatal("Error seeding the database", err)
	}

	r := chi.NewRouter()
	r.Use(CORSMiddleware)
	r.Mount("/", routes.RegisterRoutes())

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", r)
}