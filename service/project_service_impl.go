package service

import (
	"github.com/sukenda/go-restful-postgres/entity"
	"github.com/sukenda/go-restful-postgres/exception"
	"github.com/sukenda/go-restful-postgres/model"
	"github.com/sukenda/go-restful-postgres/repository"
	"github.com/sukenda/go-restful-postgres/validation"
)

func NewProjectService(projectRepository *repository.ProjectRepository) ProjectService {
	return &projectServiceImpl{
		repository: *projectRepository,
	}
}

type projectServiceImpl struct {
	repository repository.ProjectRepository
}

func (service projectServiceImpl) Insert(request model.CreateProjectRequest) model.CreateProjectResponse {
	validation.ValidateProject(request)

	project := entity.Project{
		Entity:       entity.Entity{},
		Name:         request.Name,
		Description:  request.Description,
		DepartmentID: request.DepartmentID,
	}

	result, err := service.repository.Insert(project)
	exception.PanicIfNeeded(err)

	response := model.CreateProjectResponse{
		Id:          result.Id,
		Name:        project.Name,
		Description: project.Description,
		/*Department: model.CreateDepartmentResponse{
			Id:          result.Id,
			Name:        result.Department.Name,
			Description: result.Department.Description,
		},*/
	}

	return response
}

func (service projectServiceImpl) Find() []model.CreateProjectResponse {
	results := service.repository.Find()
	var projects []model.CreateProjectResponse

	for i := range results {
		projects = append(projects, model.CreateProjectResponse{
			Id:          results[i].Id,
			Name:        results[i].Name,
			Description: results[i].Description,
		})
	}

	return projects
}
