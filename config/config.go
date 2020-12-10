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

// REDISHOST get redis host
func REDISHOST() string {
	return helper.
		GetEnv("REDIS_HOST", "localhost:6379")
}

// REDISPASSWORD get redis password
func REDISPASSWORD() string {
	return helper.
		GetEnv("REDIS_PASSWORD", "secret123")
}
