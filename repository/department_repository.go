package repository

import "github.com/sukenda/go-restful-postgres/entity"

type DepartmentRepository interface {
	Insert(param entity.Department) (entity.Department, error)

	Find() []entity.Department
}
