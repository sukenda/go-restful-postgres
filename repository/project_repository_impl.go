package repository

import (
	"github.com/sukenda/go-restful-postgres/config"
	"github.com/sukenda/go-restful-postgres/entity"
	"github.com/sukenda/go-restful-postgres/exception"
	"gorm.io/gorm"
)

func NewProjectRepository(database *gorm.DB) ProjectRepository {
	return &projectRepositoryImpl{
		Database: database,
	}
}

type projectRepositoryImpl struct {
	Database *gorm.DB
}

func (repository projectRepositoryImpl) Insert(project entity.Project) (entity.Project, error) {
	database := config.GetDatabase()

	result := database.Create(&project)
	if result.Error != nil {
		exception.PanicIfNeeded(exception.ValidationError{
			Message: result.Error.Error(),
		})
	}

	return project, nil
}

func (repository projectRepositoryImpl) Find() []entity.Project {
	database := config.GetDatabase()

	var projects []entity.Project
	database.Preload("Department").Find(&projects)

	return projects
}
