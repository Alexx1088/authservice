package jwtutil

import (
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

var secretKey = []byte(getSecretKey())

func getSecretKey() string {
	if val := os.Getenv("JWT_SECRET_KEY"); val != "" {
		return val
	}
	return "default_dev_secret"
}

func GenerateToken(userID string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}
