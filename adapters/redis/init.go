package redis

import (
	"go-boilerplate/config"

	"github.com/go-redis/redis"
)

// Init create redis client
func Init() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     config.REDISHOST(),
		Password: config.REDISPASSWORD(),
		DB:       0,
	})

	err := client.Ping().Err()
	if err != nil {
		return nil, err
	}

	return client, nil
}
