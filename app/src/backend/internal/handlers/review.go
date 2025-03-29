package handlers

import ("encoding/json"
"net/http")

func CreateReview(w http.ResponseWriter, r *http.Request){
	var review reviewjson.NewDecoder(r.Body).Decode(&review)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(review)
}

func ReadReview(w http.ResponseWriter, r *http.Request){
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	
	var review reviewjson.NewDecoder(r.Body).Decode(&review)
	w.WriteHeader(http.Statu)
}