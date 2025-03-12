package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

func CreateJWT(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	env, err := godotenv.Read(".env")
	if err != nil {
		return "", err
	}

	tokenString, err := token.SignedString([]byte(env["JWT_SECRET_KEY"]))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
