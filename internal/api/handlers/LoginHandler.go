package handlers

import (
	"encoding/json"
	"mydiary/internal/database"
	"mydiary/internal/models"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// LoginRequest represents the data for a login request
type LoginRequest struct {
    UserName string `json:"username"`
    Password string `json:"password"`
}


func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var loginReq LoginRequest

    err := json.NewDecoder(r.Body).Decode(&loginReq)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Validate user credentials
    var user models.User
    result := database.DB.Where("username = ?", loginReq.UserName).First(&user)
    if result.Error != nil {
        http.Error(w, "User not found", http.StatusUnauthorized)
        return
    }

	 // Compare the stored hashed password with the provided password
	 err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(loginReq.Password))
	 if err != nil {
		 http.Error(w, "Invalid password", http.StatusUnauthorized)
		 return
	 }
}