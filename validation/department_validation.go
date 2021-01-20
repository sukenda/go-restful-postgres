package validation

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/sukenda/go-restful-postgres/exception"
	"github.com/sukenda/go-restful-postgres/model"
)

func ValidateDepartment(request model.CreateDepartmentRequest) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.Name, validation.Required),
	)

	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}
