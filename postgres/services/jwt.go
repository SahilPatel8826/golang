package services

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// A simple secret key (keep it safe in .env in real apps)
var jwtKey = []byte("my_secret_key")

// Struct that holds data in token
type JWTClaim struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.RegisteredClaims
}

// GenerateJWT creates a token
func GenerateJWT(email, role string) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour) // token valid for 1 hour

	claims := &JWTClaim{
		Email: email,
		Role:  role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateJWT checks if token is valid
func ValidateJWT(tokenString string) (*JWTClaim, error) {
	claims := &JWTClaim{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("invalid token: %v", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("token is not valid")
	}

	return claims, nil
}
