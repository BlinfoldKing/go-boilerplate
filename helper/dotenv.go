package helper

import (
	"os"
)

func GetEnv(variable string, fallback string) string {
	res := os.Getenv(variable)
	if res != "" {
		return res
	}

	return fallback
}
