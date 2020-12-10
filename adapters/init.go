package adapters

import (
	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis"
	"go-boilerplate/adapters/postgres"
	rd "go-boilerplate/adapters/redis"
	validation "go-boilerplate/adapters/validator"
	// "gorm.io/gorm"
	"xorm.io/xorm"
)

// Adapters is wrapper for lib/drivers that needed to be injected
type Adapters struct {
	Postgres  *xorm.Engine
	Validator *validator.Validate
	Redis     *redis.Client
}

// Init create new Adapters
func Init() (Adapters, error) {
	postgres, err := postgres.Init()
	if err != nil {
		return Adapters{}, err
	}

	validator, err := validation.Init()
	if err != nil {
		return Adapters{}, err
	}

	redis, err := rd.Init()
	if err != nil {
		return Adapters{}, err
	}

	return Adapters{
		postgres,
		validator,
		redis,
	}, err
}
