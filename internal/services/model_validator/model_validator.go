package model_validator

import (
	"github.com/go-playground/validator/v10"
)

type ModelValidator interface {
	Validate(model any) error
}

type modelValidator struct {
	validator *validator.Validate
}

func NewModelValidator() ModelValidator {
	newValidator := validator.New()

	return &modelValidator{
		validator: newValidator,
	}
}

func (mv *modelValidator) Validate(model any) error {
	if err := mv.validator.Struct(model); err != nil {
		return err
	}

	return nil
}
