package adapters

import (
	"go-boilerplate/adapters/postgres"
	validation "go-boilerplate/adapters/validator"

	"github.com/go-playground/validator/v10"
	// "gorm.io/gorm"
	"xorm.io/xorm"
)

// Adapters is wrapper for lib/drivers that needed to be injected
type Adapters struct {
	Postgres  *xorm.Engine
	Validator *validator.Validate
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

	return Adapters{
		postgres,
		validator,
	}, err
}
