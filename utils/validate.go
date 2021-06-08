package utils

import "github.com/go-playground/validator/v10"

var Validate *validator.Validate

// InitValidator return a new validator instance
func InitValidator() {
	Validate = validator.New()
}
