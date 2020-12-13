package config

import (
	"fmt"
	"go-boilerplate/helper"
	"strconv"
)

// DBUSER get database user
func DBUSER() string {
	return helper.
		GetEnv("DB_USER", "root")
}

// DBHOST get database host
func DBHOST() string {
	return helper.
		GetEnv("DB_HOST", "127.0.0.1")
}

// DBPORT get database port
func DBPORT() string {
	return helper.
		GetEnv("DB_PORT", "26257")
}

// DBSSLMODE get database sslmode
func DBSSLMODE() string {
	return helper.
		GetEnv("DB_SSLMODE", "disable")
}

// DBDATABASE get database name
func DBDATABASE() string {
	return helper.
		GetEnv("DB_DATABASE", "postgres")
}

// DBPASSWORD get database password
func DBPASSWORD() string {
	return helper.
		GetEnv("DB_PASSWORD", "secret123")
}

// DBCONFIG get connection string of db
func DBCONFIG() string {
	return fmt.Sprintf(
		"dbname=%s user=%s password=%s host=%s port=%s sslmode=%s",
		DBDATABASE(),
		DBUSER(),
		DBPASSWORD(),
		DBHOST(),
		DBPORT(),
		DBSSLMODE(),
	)
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

// TOKENDURATION get redis password
func TOKENDURATION() int {
	res := helper.
		GetEnv("TOKEN_DURATION", "3600")

	dur, err := strconv.Atoi(res)
	if err != nil {
		return 3600
	}
	return dur
}
