package model

import (
	"fmt"
	"postgres/middleware"
)

type User struct {
	ID        string `json:id`
	Email     string `json:email`
	Password  string `json:password`
	Role      string `json:"role"`
	AccountID string `json:"account_id"`
}

func RegisterUser(user User) (User, error) {
	db := middleware.CreateConnection()
	defer db.Close()
	sqlStatement2 := `SELECT email FROM users WHERE email=$1`
	var existingUser User
	err := db.QueryRow(sqlStatement2, user.Email).Scan(&existingUser.ID)
	if err == nil {
		return User{}, fmt.Errorf("email already exists")

	}

	sqlStatement := `
		INSERT INTO users (email, password, role, account_id)
		VALUES ($1, $2, $3, $4)
		RETURNING id, email, password, role, account_id;`

	var newUser User
	err = db.QueryRow(sqlStatement, user.Email, user.Password, user.Role, user.AccountID).
		Scan(&newUser.ID, &newUser.Email, &newUser.Password, &newUser.Role, &newUser.AccountID)

	if err != nil {
		return User{}, fmt.Errorf("failed to register user: %v", err)
	}

	fmt.Println("✅ User registered successfully:", newUser)
	return newUser, nil
}
func LoginUser(email string, password string) (User, error) {
	db := middleware.CreateConnection()
	defer db.Close()
	var user User
	sqlStatement := `SELECT * FROM users WHERE email=$1`
	err := db.QueryRow(sqlStatement, email).Scan(&user.ID, &user.Email, &user.Password, &user.Role, &user.AccountID)

	if err != nil {
		fmt.Errorf("failed login:", err)
	}
	return user, nil
}
