package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"postgres/model"
	"postgres/services"
)

func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	var user model.User

	json.NewDecoder(r.Body).Decode(&user)

	hashedPassword, err := services.HashPassword(user.Password)
	if err != nil {
		fmt.Println("Pass is not hashed:", err)
		return
	}
	user.Password = hashedPassword

	result, err := model.RegisterUser(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	json.NewEncoder(w).Encode(result)
}

func LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	result, err := model.LoginUser(user.Email, user.Password)
	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	if err := services.CheckPassword(user.Password, result.Password); err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	token, err := services.GenerateJWT(user.Email, user.Role)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	// Remove password before sending back
	result.Password = ""

	json.NewEncoder(w).Encode(map[string]string{
		"message":  "Login successful",
		"token":    token,
		"email":    result.Email,
		"password": result.Password,
	})
}
