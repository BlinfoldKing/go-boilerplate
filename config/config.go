package config

import (
	"fmt"
	"go-boilerplate/helper"
	"strconv"
	"time"
)

// ENV get env
func ENV() string {
	return helper.
		GetEnv("ENV", "development")
}

// PREFIX get env
func PREFIX() string {
	return helper.
		GetEnv("PREFIX", "/v1")
}

// APPURL gets app url
func APPURL() string {
	return helper.
		GetEnv("APP_URL", "localhost")
}

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
	password := DBPASSWORD()
	if password == "" {
		return fmt.Sprintf(
			"dbname=%s user=%s host=%s port=%s sslmode=%s",
			DBDATABASE(),
			DBUSER(),
			DBHOST(),
			DBPORT(),
			DBSSLMODE(),
		)
	}
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

// MINIOENDPOINT gets minio endpoint
func MINIOENDPOINT() string {
	return helper.
		GetEnv("MINIO_ENDPOINT", "play.min.io")
}

// MINIOACCESSKEY gets access key for minio, default is a public key for default users in minio
func MINIOACCESSKEY() string {
	return helper.
		GetEnv("MINIO_ACCESS_KEY", "Q3AM3UQ867SPQQA43P2F")
}

// MINIOSECRET gets minio secret key, default is a public key for default users in minio
func MINIOSECRET() string {
	return helper.
		GetEnv("MINIO_SECRET", "zuf+tfteSlswRu7BJ86wekitnifILbZam1KYY3TG")
}

// MINIOEXPIRE gets minio link expire duration
func MINIOEXPIRE() time.Duration {
	exp, _ := time.ParseDuration(helper.GetEnv("MINIO_EXPIRE", "86400s"))
	return exp
}

// MINIOREGION gets minio bucket region
func MINIOREGION() string {
	return helper.
		GetEnv("MINIO_REGION", "us-east-1")
}

// MAILGUNDOMAIN gets mailgun domain
func MAILGUNDOMAIN() string {
	return helper.
		GetEnv("MAILGUN_DOMAIN", " ")
}

// MAILGUNAPIKEY gets mailgun api key
func MAILGUNAPIKEY() string {
	return helper.
		GetEnv("MAILGUN_API_KEY", " ")
}

// MAILGUNPUBLICAPIKEY gets mailgun api key
func MAILGUNPUBLICAPIKEY() string {
	return helper.
		GetEnv("MAILGUN_PUBLIC_API_KEY", " ")
}

// MAILGUNURL gets mailgun url
func MAILGUNURL() string {
	return helper.
		GetEnv("MAILGUN_URL", " ")
}

// NATSURI gets nats uri
func NATSURI() string {
	return helper.
		GetEnv("NATS_URI", "nats://localhost:4222")

}

// EMAILACTIVATION gets whether or not email activation is needed
func EMAILACTIVATION() bool {
	activation, _ := strconv.ParseBool(
		helper.
			GetEnv("EMAIL_ACTIVATION", "true"))
	return activation
}

// OTPDURATION gets otp expire duration
func OTPDURATION() time.Duration {
	dur, _ := time.ParseDuration(helper.GetEnv("OTP_DURATION", "3600s"))
	return dur
}
