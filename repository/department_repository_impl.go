package repository

import (
	"github.com/sukenda/go-restful-postgres/config"
	"github.com/sukenda/go-restful-postgres/entity"
	"github.com/sukenda/go-restful-postgres/exception"
	"gorm.io/gorm"
)

func NewDepartmentRepository(database *gorm.DB) DepartmentRepository {
	return &departmentRepositoryImpl{Database: database}
}

type departmentRepositoryImpl struct {
	Database *gorm.DB
}

func (repository departmentRepositoryImpl) Insert(param entity.Department) (department entity.Department, err error) {
	database := config.GetDatabase()

	result := database.Create(&param)
	if result.Error != nil {
		exception.PanicIfNeeded(exception.ValidationError{
			Message: "ERROR_INSERT",
		})
	}

	return param, nil
}

func (repository departmentRepositoryImpl) Find() (departments []entity.Department) {
	database := config.GetDatabase()
	database.Find(&departments)

	return departments
}
