package utils

import (
	"github.com/go-playground/validator/v10"
)

// Validator struct that holds the validator instance
type Validator struct {
	validate *validator.Validate
}

// NewValidator creates a new instance of the Validator
func NewValidator() *Validator {
	return &Validator{
		validate: validator.New(),
	}
}

// ValidateStruct validates the given struct based on its tags
func (v *Validator) ValidateStruct(i interface{}) error {
	return v.validate.Struct(i)
}
