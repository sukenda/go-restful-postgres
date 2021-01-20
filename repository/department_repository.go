package repository

import "github.com/sukenda/go-restful-postgres/entity"

type DepartmentRepository interface {
	Insert(param entity.Department) (department entity.Department, err error)

	Find() (departments []entity.Department)
}
