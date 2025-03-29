package main

import (
	"net/http"
	"winebaby/handlers"
)

func main() {
    mux := http.NewServeMux()
	mux.HandleFunc("/wines", handlers.GetWines)
	mux.HandleFunc("/reviews", handlers.CreateReview)
	mux.HandleFunc("/", handlers.Home)
}