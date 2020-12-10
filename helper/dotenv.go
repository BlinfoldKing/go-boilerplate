package helper

import (
	"os"
)

// GetEnv read env variable or use a fallback
func GetEnv(variable string, fallback string) string {
	res := os.Getenv(variable)
	if res != "" {
		return res
	}

	return fallback
}
