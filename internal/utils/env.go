package utils

import "github.com/joho/godotenv"

func Env(title string) (string, error) {
	env, err := godotenv.Read(".env")
	if err != nil {
		return "", err
	}

	return env[title], nil
}
