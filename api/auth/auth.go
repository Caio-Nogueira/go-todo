package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/Caio-Nogueira/go-todo/database"
	"github.com/Caio-Nogueira/go-todo/models"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	DB := database.DB
	var loginUser models.LoginUser

	// get LoginUser from request body
	err := json.NewDecoder(r.Body).Decode(&loginUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// check if user exists
	var user models.User
	DB.Where("username = ?", loginUser.Username).First(&user)

	if user.Username == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// check if password is correct
	validPassword := models.ComparePassword(user.Password, loginUser.Password)
	if !validPassword {
		fmt.Println("Invalid password")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// generate JWT
	token, err := GenerateJWT(user.Username)
	
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// return JWT in cookie
	w.Header().Set("Content-Type", "application/json")
	http.SetCookie(w, &http.Cookie{
		Name: "jwt",
		Value: token,
		MaxAge: 3600, // 1 hour ttl
		HttpOnly: true,
	})

	json.NewEncoder(w).Encode(user)
}