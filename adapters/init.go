package adapters

import (
	"go-boilerplate/adapters/enforcer"
	"go-boilerplate/adapters/postgres"
	rd "go-boilerplate/adapters/redis"
	validation "go-boilerplate/adapters/validator"

	"github.com/casbin/casbin/v2"
	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis"
)

// Adapters is wrapper for lib/drivers that needed to be injected
type Adapters struct {
	Postgres  *postgres.Postgres
	Validator *validator.Validate
	Redis     *redis.Client
	Enforcer  *casbin.Enforcer
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

	enforcer, err := enforcer.Init()
	if err != nil {
		return Adapters{}, err
	}

	return Adapters{
		postgres,
		validator,
		redis,
		enforcer,
	}, err
}
