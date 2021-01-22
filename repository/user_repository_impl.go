package repository

import (
	"errors"
	"github.com/sukenda/go-restful-postgres/config"
	"github.com/sukenda/go-restful-postgres/entity"
	"github.com/sukenda/go-restful-postgres/exception"
	"github.com/sukenda/go-restful-postgres/validation"
	"gorm.io/gorm"
)

func NewUserRepository(database *gorm.DB) UserRepository {
	return &userRepositoryImpl{Database: database}
}

type userRepositoryImpl struct {
	Database *gorm.DB
}

func (repository userRepositoryImpl) Insert(param entity.User) (entity.User, error) {
	database := config.GetDatabase()

	_, err := repository.FindByUsername(param.Username)
	if err == nil {
		exception.PanicIfNeeded(exception.ValidationError{
			Message: "USER_EXIST",
		})
	}

	result := database.Create(&param)
	if result.Error != nil {
		exception.PanicIfNeeded(exception.ValidationError{
			Message: "EMAIL_EXIST",
		})
	}

	return param, nil
}

func (repository userRepositoryImpl) FindByUsername(username string) (entity.User, error) {
	database := config.GetDatabase()

	var user entity.User
	database.Where("username = ?", username).Preload("Employee").First(&user)

	if len(user.Username) == 0 {
		return user, errors.New("USER_NOT_FOUND")
	}

	return user, nil
}

func (repository userRepositoryImpl) Login(username, password string) (entity.User, error) {
	database := config.GetDatabase()

	var user entity.User
	database.Where("username = ?", username).Preload("Employee").First(&user)

	match, _ := validation.ValidatePassword(password, user.Password)
	if !match {
		exception.PanicIfNeeded(exception.ValidationError{
			Message: "PASSWORD_WRONG",
		})
	}

	return user, nil
}
