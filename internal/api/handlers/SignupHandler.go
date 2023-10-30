package handlers

import (
	"encoding/json"
	"mydiary/internal/database"
	"mydiary/internal/models"
	"net/http"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

var validate *validator.Validate

func init() {
    validate = validator.New()
}

//registration request struct separates the concerns of incoming data (which includes a plaintext password)
// and the stored data (which only includes the hashed password).
type RegistrationRequest struct {
	UserName string `json:"username" validate:"required,min=5,max=15"`
	Password string `json:"password" validate:"required,min=12"`
	Email    string `json:"email" validate:"required,email"`
}


func SignupHandler(w http.ResponseWriter, r *http.Request) {
	var regReq RegistrationRequest


	// decode the request body 
	err := json.NewDecoder(r.Body).Decode(&regReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	 // Validate request
	err = validate.Struct(regReq)
	if err != nil {
		 http.Error(w, err.Error(), http.StatusBadRequest)
		 return
	}

	// Check for existing user with the same username or email
	var existingUser models.User
	if err := database.DB.Where("UserName = ? OR Email = ?", regReq.UserName, regReq.Email).First(&existingUser).Error; err == nil {
		   http.Error(w, "User already exists", http.StatusConflict)
		   return
	}

	 //hass password 
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(regReq.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Server Error Password", http.StatusInternalServerError)
		return
	}

	// Create user with hashed password
	user := models.User{
		UserName:     regReq.UserName,
		PasswordHash: string(hashedPassword),
		Email:        regReq.Email, 
	}

    // Save the user to the database
    result := database.DB.Create(&user)
    if result.Error != nil {
        http.Error(w, "Server error", http.StatusInternalServerError)
        return
    }

    // Respond to the request
    w.WriteHeader(http.StatusCreated)
}
