package validator

import "github.com/go-playground/validator/v10"

func Init() (*validator.Validate, error) {
	return validator.New(), nil
}
