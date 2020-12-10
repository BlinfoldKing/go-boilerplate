package config

import "go-boilerplate/helper"

// DBURL get database url
func DBURL() string {
	return helper.
		GetEnv("DB_URL", "postgresql://root@database:26257/postgres?sslmode=disable")
}

// PORT get server port
func PORT() string {
	return helper.
		GetEnv("PORT", "8000")
}
