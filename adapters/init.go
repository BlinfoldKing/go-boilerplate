package adapters

import (
	"go-boilerplate/adapters/postgres"
	validation "go-boilerplate/adapters/validator"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Adapters struct {
	Postgres  *gorm.DB
	Validator *validator.Validate
}

func Init() (Adapters, error) {
	postgres, err := postgres.Init()
	if err != nil {
		return Adapters{}, err
	}

	validator, err := validation.Init()
	if err != nil {
		return Adapters{}, err
	}

	return Adapters{
		postgres,
		validator,
	}, err
}
