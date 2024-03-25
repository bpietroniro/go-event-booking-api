package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken() {
	jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  "",
		"userID": "",
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})
}
