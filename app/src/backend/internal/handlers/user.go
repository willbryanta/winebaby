package handlers

import (
	"encoding/json"
	"net/http"
)

func SignUp(w http.ResponseWriter, r *http.Request){
	var user userjson.NewDecoder(r.Body).Decode(&user)
	w.WriteHeader(http.StatusCreated)
}

func Login(w http.ResponseWriter, r *http.Request){
	
}