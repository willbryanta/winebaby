package handlers

import ("encoding/json"
"net/http"
"winebaby/internal/models")

func GetWines(w http.ResponseWriter, r *http.Request){
	wines := []Wine{}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(wines)
}

func AddWine(w http.ResponseWriter, r *http.Request){
	var wine Wine
	if err := json.NewDecoder(r.Body).Decode(&wine); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(wine)
}