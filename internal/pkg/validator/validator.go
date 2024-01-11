package validator

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type ValidatorError struct {
	Property string `json:"property"`
	Message  string `json:"message"`
	Value    string `json:"value"`
	Tag      string `json:"tag"`
}

var HttpStatusCode = map[string]int{
	"OK":                  200,
	"BadRequest":          400,
	"Unauthorized":        401,
	"Forbidden":           403,
	"NotFound":            404,
	"InternalServerError": 500,
}

func GetErrorValidator(err error) *[]ValidatorError {
	var e []ValidatorError
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		for _, err := range ve {
			e = append(e, ValidatorError{
				Property: err.Field(),
				Message:  err.Error(),
				Value:    err.Value().(string),
				Tag:      err.Tag(),
			})
		}
		return &e
	}
	return nil
}
