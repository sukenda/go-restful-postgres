package repository

import "github.com/sukenda/go-restful-postgres/entity"

type ProjectRepository interface {
	Insert(project entity.Project) (entity.Project, error)

	Find() []entity.Project
}
