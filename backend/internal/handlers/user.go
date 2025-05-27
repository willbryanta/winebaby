package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
	"winebaby/internal/models"
	"winebaby/internal/repository"

	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var JWTSecret = []byte(os.Getenv("JWT_SECRET"))
var JWTExpiration = 3600

func SignUp(w http.ResponseWriter, r *http.Request, repo *repository.MainRepository, db *sql.DB) {
	var user models.User
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "Invalid request body"})
		return
	}
	
	user.Username = strings.TrimSpace(user.Username)
	user.Password = strings.TrimSpace(user.Password)
	if user.Email != nil {
		*user.Email = strings.TrimSpace(*user.Email)
	}

	if user.Username == "" || user.Password == "" || (user.Email != nil && *user.Email == "") {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "Username, password, and email are required"})
		return
	}
	if user.Email != nil && *user.Email != "" {
		if !regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`).MatchString(*user.Email) {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"message": "Invalid email format"})
			return
		}
	}

	existingUser, err := repo.GetUserByUsername(user.Username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Failed to check existing user: " + err.Error()})
		return
	}
	if existingUser.Username != "" {
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(map[string]string{"message": "Username already exists"})
		return
	}
	if user.Email != nil {
		existingEmail, err := repo.GetUserByEmail(*user.Email)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"message": "Failed to check existing email: " + err.Error()})
			return
		}
		if existingEmail.ID != 0 {
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(map[string]string{"message": "Email already exists"})
			return
		}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Failed to hash password"})
		return
	}
	user.Password = string(hashedPassword)


	if err := repo.CreateUser(user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Failed to create user: " + err.Error()})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": 	 user.ID,
		"username": user.Username,
		"email":    user.Email,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err:= token.SignedString(JWTSecret)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Failed to generate token: " + err.Error()})
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name: "token",
		Value: tokenString,
		Path: "/",
		HttpOnly: true,
		Secure: true,
		MaxAge: 86400,
	})

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User created successfully"})
}

func SignIn(w http.ResponseWriter, r *http.Request, repo *repository.MainRepository, db *sql.DB) {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "Invalid request body"})
		return
	}

	credentials.Username = strings.TrimSpace(credentials.Username)
	credentials.Password = strings.TrimSpace(credentials.Password)
	
	if credentials.Username == "" || credentials.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "Username and password are required"})
		return
	}
	user, err := repo.GetUserByUsername(credentials.Username)
	if err != nil {	
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Failed to fetch user: " + err.Error()})
		return
	}
	if user.Username == "" {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"message": "Invalid username or password"})
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"message": "Invalid username or password"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": 	 user.ID,
		"username": user.Username,
		"email":    user.Email,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err:= token.SignedString(JWTSecret)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Failed to generate token: " + err.Error()})
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name: "token",
		Value: tokenString,
		Path: "/",
		HttpOnly: true,
		Secure: true,
		MaxAge: 86400,
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}

func UpdateUserProfile(w http.ResponseWriter, r *http.Request, repo *repository.MainRepository, db *sql.DB) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "Invalid request body"})
		return
	}

	if user.Username == "" || (user.Email != nil && *user.Email == "") {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "Username and email are required"})
		return
	}

	if err := repo.UpdateUserProfile(user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Failed to update user profile"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "User profile updated successfully"})
}
func DeleteUser(w http.ResponseWriter, r *http.Request, repo *repository.MainRepository, db *sql.DB) {
	username := r.URL.Path[len("/api/users/"):]
	if err := repo.DeleteUser(username); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Failed to delete user"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "User deleted successfully"})
}
func GetUserFavoriteWines(w http.ResponseWriter, r *http.Request, repo *repository.MainRepository, db *sql.DB) {
	username := r.URL.Path[len("/api/users/"):]
	favoriteWines, err := repo.GetUserFavoriteWines(username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Failed to fetch favorite wines"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(favoriteWines)
}
func AddUserFavoriteWine(w http.ResponseWriter, r *http.Request, repo *repository.MainRepository, db *sql.DB) {
	var wine models.Wine
	err := json.NewDecoder(r.Body).Decode(&wine)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "Invalid request body"})
		return
	}

	userID := chi.URLParam(r, "userID")
	wineID := chi.URLParam(r, "wineID")

	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "Invalid user ID"})
		return
	}
	wineIDInt, err := strconv.Atoi(wineID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "Invalid wine ID"})
		return
	}
	if err := repo.AddUserFavoriteWine(userIDInt, wineIDInt); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Failed to add favorite wine"})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Favorite wine added successfully"})
}
func RemoveUserFavoriteWine(w http.ResponseWriter, r *http.Request, repo *repository.MainRepository, db *sql.DB) {
	userID := chi.URLParam(r, "userID") 
	wineID := chi.URLParam(r, "wineID")

	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "Invalid user ID"})
		return
	}
	wineIDInt, err := strconv.Atoi(wineID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "Invalid wine ID"})
		return
	}

	if err := repo.RemoveUserFavoriteWine(userIDInt, wineIDInt); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Failed to remove favorite wine"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Favorite wine removed successfully"})
}
func GetUserReviews(w http.ResponseWriter, r *http.Request, repo *repository.MainRepository, db *sql.DB) {
	username := r.URL.Path[len("/api/users/"):] 
	reviews, err := repo.GetUserReviews(username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Failed to fetch user reviews"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reviews)
}
func CreateUserReview(w http.ResponseWriter, r *http.Request, repo *repository.MainRepository, db *sql.DB) {
	var review models.Review
	err := json.NewDecoder(r.Body).Decode(&review)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "Invalid request body"})
		return
	}

	if err := repo.CreateUserReview(review); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Failed to create user review"})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User review created successfully"})
}
func UpdateUserReview(w http.ResponseWriter, r *http.Request, repo *repository.MainRepository, db *sql.DB) {
	var review models.Review
	err := json.NewDecoder(r.Body).Decode(&review)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "Invalid request body"})
		return
	}

	if err := repo.UpdateUserReview(review); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Failed to update user review"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "User review updated successfully"})
}
func DeleteUserReview(w http.ResponseWriter, r *http.Request, repo *repository.MainRepository, db *sql.DB) {
	reviewIdStr := chi.URLParam(r, "reviewId")
	reviewId, err := strconv.Atoi(reviewIdStr)
	if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(map[string]string{"message": "Invalid review ID"})
        return
    }
	if err := repo.DeleteUserReview(reviewId); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Failed to delete user review"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "User review deleted successfully"})
}
func GetUserReviewById(w http.ResponseWriter, r *http.Request, repo *repository.MainRepository, db *sql.DB) {
	reviewIdStr := chi.URLParam(r, "reviewId")
	reviewId, err := strconv.Atoi(reviewIdStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "Invalid review ID"})
		return
	}
	review, err := repo.GetUserReviewById(reviewId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Failed to fetch user review"})
		return
	}
	if review.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"message": "User review not found"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(review)
}

func GetUserProfile(w http.ResponseWriter, r *http.Request, repo *repository.MainRepository, db *sql.DB) {
	username := r.URL.Path[len("/api/users/"):] 
	user, err := repo.GetUserProfile(username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Failed to fetch user profile"})
		return
	}
	if user.Username == "" {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"message": "User not found"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

