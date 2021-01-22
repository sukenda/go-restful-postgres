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

func (repository departmentRepositoryImpl) Insert(department entity.Department) (entity.Department, error) {
	database := config.GetDatabase()

	result := database.Create(&department)
	if result.Error != nil {
		exception.PanicIfNeeded(exception.ValidationError{
			Message: result.Error.Error(),
		})
	}

	return department, nil
}

func (repository departmentRepositoryImpl) Find() []entity.Department {
	database := config.GetDatabase()

	var departments []entity.Department
	database.Preload("Projects").Find(&departments)

	return departments
}
