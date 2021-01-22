package service

import "github.com/sukenda/go-restful-postgres/model"


type DepartmentService interface {
	Insert(request model.CreateDepartmentRequest) model.CreateDepartmentResponse

	Find() []model.CreateDepartmentResponse
}
