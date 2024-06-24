package auth

import (
	"e-commerce-back/internal/models"
	"e-commerce-back/pkg/db"
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User

	fmt.Println("RegisterHandler")

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user.Password = string(hashedPassword)
	if err := db.DB.Create(&user).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var storedUser models.User
	if err := db.DB.Where("email = ?", user.Email).First(&storedUser).Error; err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)); err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
}
