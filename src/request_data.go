package src

import (
	"github.com/go-playground/validator/v10"
)

type RegisterRequest struct {
	Email                string `json:"email" validate:"required,email"`
	Name                 string `json:"name,omitempty"`
	Gender               string `json:"gender,omitempty"`
	Address              string `json:"address,omitempty"`
	Password             string `json:"password" validate:"required,min=6"`
	PasswordConfirmation string `json:"password_confirmation" validate:"required,min=6,eqfield=Password"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type ErrorResponse struct {
	Field string `json:"field,omitempty"`
	Tag   string `json:"tag,omitempty"`
	Value string `json:"value,omitempty"`
}

var validate = validator.New()

func ValidateStruct(data interface{}) []ErrorResponse {
	var errs []ErrorResponse
	err := validate.Struct(data)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.Field = err.Field()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errs = append(errs, element)
		}
	}
	return errs
}
