package service

import "github.com/sukenda/go-restful-postgres/model"

type ProjectService interface {
	Insert(request model.CreateProjectRequest) model.CreateProjectResponse

	Find() []model.CreateProjectResponse
}
