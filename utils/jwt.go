package utils

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

var secretKey = GoDotEnvVariable("SECRET_KEY")

func GoDotEnvVariable(key string) string {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func GenerateToken(email string, userId int64) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})
	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) error {
	parseToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing method")
		}
		return secretKey, nil
	})

	if err != nil {
		return errors.New("could not parse token")
	}

	if !parseToken.Valid {
		return errors.New("invalid token")
	}

	// claims, ok := parseToken.Claims.(jwt.MapClaims)

	// if !ok {
	// 	return errors.New("Invalid token claims.")
	// }

	// email := claims["email"].(string)
	// userId := claims["userId"].(int64)

	return nil
}
