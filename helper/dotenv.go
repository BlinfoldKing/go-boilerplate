package helper

import (
	"github.com/joho/godotenv"
	"os"
)

func GetEnv(variable string, fallback string) (string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", err
	}

	res := os.Getenv(variable)
	if res != "" {
		return res, nil
	}

	return fallback, nil
}
