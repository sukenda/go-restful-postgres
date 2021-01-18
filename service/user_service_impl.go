package service

import (
	"github.com/sukenda/go-restful-postgre/entity"
	"github.com/sukenda/go-restful-postgre/exception"
	"github.com/sukenda/go-restful-postgre/model"
	"github.com/sukenda/go-restful-postgre/repository"
	"github.com/sukenda/go-restful-postgre/validation"
)

func NewUserService(userRepository *repository.UserRepository) UserService {
	return &userServiceImpl{
		repository: *userRepository,
	}
}

type userServiceImpl struct {
	repository repository.UserRepository
}

func (service userServiceImpl) Register(request model.CreateUserRequest) (response model.CreateUserResponse) {
	validation.ValidateUser(request)

	pass, err := validation.HashPassword(request.Password)
	exception.PanicIfNeeded(err)

	user := entity.User{
		Username: request.Username,
		Password: pass,
		Email:    request.Email,
		Phone:    request.Phone,
	}

	err = service.repository.Insert(user)
	exception.PanicIfNeeded(err)

	response = model.CreateUserResponse{
		Id:       user.Id.String(),
		Username: user.Username,
		Email:    user.Email,
		Phone:    user.Phone,
	}

	return response
}

func (service userServiceImpl) FindByUsername(username string) (response model.CreateUserResponse) {
	user, err := service.repository.FindByUsername(username)
	exception.PanicIfNeeded(err)

	response = model.CreateUserResponse{
		Id:       user.Id.String(),
		Username: user.Username,
		Email:    user.Email,
		Phone:    user.Phone,
	}

	return response
}

func (service userServiceImpl) Login(request model.CreateUserRequest) (response model.GetLoginResponse) {
	validation.ValidateLogin(request)

	user, err := service.repository.Login(request.Username, request.Password)
	exception.PanicIfNeeded(err)

	token, err := validation.CreateToken(user)
	exception.PanicIfNeeded(err)

	response = model.GetLoginResponse{
		AccessToken: token,
		User: model.GetUserResponse{
			Id:       user.Id.String(),
			Username: user.Username,
			Email:    user.Email,
			Phone:    user.Phone,
		},
	}

	return response
}
