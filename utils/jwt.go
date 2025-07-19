package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"time"
)

const secretKey = "supersecret"

func CreateToken(username string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"userId":   userId,
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	return token.SignedString([]byte(secretKey))
}

func VerifyToken(tokenString string) (int64, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Unexpected signing method")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return 0, errors.New("Could not parse token")
	}

	if !token.Valid {
		return 0, errors.New("Invalid token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("Invalid token claims")
	}

	userId := int64(claims["userId"].(float64))
	log.Println("userId:", userId)
	return userId, nil
}
