package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
	"winebaby/internal/models"
	"winebaby/internal/repository"

	"github.com/go-chi/chi"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var JWTSecret = []byte(os.Getenv("JWT_SECRET"))
var JWTExpiration = 3600

func VerifyToken(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    cookie, err := r.Cookie("token")
	fmt.Print("Verifying token: ", cookie)
    if err != nil {
        w.WriteHeader(http.StatusUnauthorized)
        json.NewEncoder(w).Encode(map[string]any{
            "isAuthenticated": false,
            "message":         "No token provided",
        })
        return
    }

    token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (any, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return JWTSecret, nil
    })

    if err != nil || !token.Valid {
        w.WriteHeader(http.StatusUnauthorized)
        json.NewEncoder(w).Encode(map[string]any{
            "isAuthenticated": false,
            "message":         "Invalid or expired token",
        })
        return
    }

    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok || !token.Valid {
        w.WriteHeader(http.StatusUnauthorized)
        json.NewEncoder(w).Encode(map[string]any{
            "isAuthenticated": false,
            "message":         "Invalid token claims",
        })
        return
    }

    username, ok := claims["username"].(string)
    if !ok {
        w.WriteHeader(http.StatusUnauthorized)
        json.NewEncoder(w).Encode(map[string]any{
            "isAuthenticated": false,
            "message":         "Username not found in token",
        })
        return
    }
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]any{
        "isAuthenticated": true,
        "username":        username,
    })
}


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


	if err := repo.CreateUser(&user); err != nil {
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
		Secure: false,
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
        "id":       user.ID,
        "username": user.Username,
        "email":    user.Email,
        "exp":      time.Now().Add(time.Hour * 24).Unix(),
    })

    tokenString, err := token.SignedString(JWTSecret)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(map[string]string{"message": "Failed to generate token: " + err.Error()})
        return
    }

    http.SetCookie(w, &http.Cookie{
        Name:     "token",
        Value:    tokenString,
        Path:     "/",
        HttpOnly: true,
        Secure:   false,
        MaxAge:   86400,
    })

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "Sign-in successful"})
}

func SignOut(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name: "token",
		Value: "",
		Path: "/",
		HttpOnly: true,
		Secure: false,
		MaxAge: -1,
	})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]bool{"success": true})
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

	if user.Password != "" {
		var reqBody struct {
			OldPassword string `json:"old_password"`
		}
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil || reqBody.OldPassword == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"message": "Old password required for password change"})
			return
		}
		existingUser, err := repo.GetUserByUsername(user.Username)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"message": "Failed to fetch user for password verification"})
			return
		}
		if err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(reqBody.OldPassword)); err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"message": "Invalid old password"})
			return
		}
	}

	if err := repo.UpdateUserProfile(user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Failed to update user profile"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "User profile updated successfully"})
}

func GetUserProfile(w http.ResponseWriter, r *http.Request, repo *repository.MainRepository, db *sql.DB) {
	username := chi.URLParam(r, "username")
	if username == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "Username is required"})
		return
	}
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

