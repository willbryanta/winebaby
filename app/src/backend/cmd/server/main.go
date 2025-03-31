package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"winebaby/db"
	"github.com/go-chi/chi/v5"
	"winebaby/internal/routes"
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

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, falling back to environment variables")
	}

	db, err := connectDB()
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal("Error pinging the database: ", err)
	}
	log.Println("Successfully connected to winebaby_db!")

	if err := db.Seed(db); err != nil{
		log.Fatal("Error seeding the database", err)
	}

	r := chi.NewRouter()
	r.Mount("/", routes.RegisterRoutes())

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", r)
}