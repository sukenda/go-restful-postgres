package service

import (
	"github.com/sukenda/go-restful-postgres/entity"
	"github.com/sukenda/go-restful-postgres/exception"
	"github.com/sukenda/go-restful-postgres/model"
	"github.com/sukenda/go-restful-postgres/repository"
	"github.com/sukenda/go-restful-postgres/validation"
)

func NewDepartmentService(departmentRepository *repository.DepartmentRepository) DepartmentService {
	return &departmentServiceImpl{
		repository: *departmentRepository,
	}
}

type departmentServiceImpl struct {
	repository repository.DepartmentRepository
}

func (service departmentServiceImpl) Insert(request model.CreateDepartmentRequest) model.CreateDepartmentResponse {
	validation.ValidateDepartment(request)

	department := entity.Department{
		Entity:      entity.Entity{},
		Name:        request.Name,
		Description: request.Description,
	}

	result, err := service.repository.Insert(department)
	exception.PanicIfNeeded(err)

	response := model.CreateDepartmentResponse{
		Id:          result.Id,
		Name:        department.Name,
		Description: department.Description,
	}

	return response
}

func (service departmentServiceImpl) Find() []model.CreateDepartmentResponse {
	results := service.repository.Find()
	var departments []model.CreateDepartmentResponse

	for i := range results {

		var projects []model.CreateProjectResponse
		for p := range results[i].Projects {
			projects = append(projects, model.CreateProjectResponse{
				Id:          results[i].Projects[p].Id,
				Name:        results[i].Projects[p].Name,
				Description: results[i].Projects[p].Description,
			})
		}

		departments = append(departments, model.CreateDepartmentResponse{
			Id:          results[i].Id,
			Name:        results[i].Name,
			Description: results[i].Description,
			Projects:    projects,
		})
	}

	return departments
}
