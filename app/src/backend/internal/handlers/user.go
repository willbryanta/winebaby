package handlers

import (
	"encoding/json"
	"net/http"
	"winebaby/internal/repository"
)

func SignUp(w http.ResponseWriter, r *http.Request){
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := repository.CreateUser(user); err != nil {
		http.Error(w, "Failed to create new user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
