package validation

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/sukenda/go-restful-postgre/exception"
	"github.com/sukenda/go-restful-postgre/model"
)

func ValidateUser(request model.CreateUserRequest) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.Id, validation.Required),
		validation.Field(&request.Username, validation.Required),
		validation.Field(&request.Password, validation.Required),
		validation.Field(&request.Email, validation.Required, is.Email),
	)

	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}

func ValidateLogin(request model.CreateUserRequest) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.Username, validation.Required),
		validation.Field(&request.Password, validation.Required),
	)

	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}
