package validator

import "github.com/go-playground/validator/v10"

// Init Create new validator
func Init() (*validator.Validate, error) {
	return validator.New(), nil
}
