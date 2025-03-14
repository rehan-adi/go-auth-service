package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
    "github.com/rehan-adi/go-auth-service/config"
)


func GenerateToken(userID int, email string) (string, error) {

	payload := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Expires in 24 hours
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, payload)
	return token.SignedString([]byte(config.JWT_SECRET))

}
