package handlers

import ("encoding/json"
"net/http"
"winebaby/internal/models"
"winebaby/internal/repository")

func GetWines(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	wines := []models.Wine{}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(wines)
}

func AddWine(w http.ResponseWriter, r *http.Request){
	var wine models.Wine
	if err := json.NewDecoder(r.Body).Decode(&wine); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(wine)
}