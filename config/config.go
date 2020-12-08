package config

import "go-boilerplate/helper"

func DB_URL() (string, error) {
	return helper.
		GetEnv("DB_URL", "postgresql://root@database:26257/postgres?sslmode=disable")
}

func PORT() (string, error) {
	return helper.
		GetEnv("PORT", "8000")
}
