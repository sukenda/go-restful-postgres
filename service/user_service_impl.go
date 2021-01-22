package service

import (
	"github.com/sukenda/go-restful-postgres/entity"
	"github.com/sukenda/go-restful-postgres/exception"
	"github.com/sukenda/go-restful-postgres/model"
	"github.com/sukenda/go-restful-postgres/repository"
	"github.com/sukenda/go-restful-postgres/validation"
)

func NewUserService(userRepository *repository.UserRepository) UserService {
	return &userServiceImpl{
		repository: *userRepository,
	}
}

type userServiceImpl struct {
	repository repository.UserRepository
}

func (service userServiceImpl) Register(request model.CreateUserRequest) model.CreateUserResponse {
	validation.ValidateUser(request)

	pass, err := validation.HashPassword(request.Password)
	exception.PanicIfNeeded(err)

	user := entity.User{
		Username: request.Username,
		Password: pass,
		Email:    request.Email,
		Employee: entity.Employee{
			FullName: request.FullName,
			Phone:    request.Phone,
			Age:      request.Age,
			Status:   "Active",
		},
	}

	result, err := service.repository.Insert(user)
	exception.PanicIfNeeded(err)

	response := model.CreateUserResponse{
		Id:       result.Id,
		Username: user.Username,
		Email:    user.Email,
		Employee: model.CreateEmployeeResponse{
			FullName: request.FullName,
			Phone:    request.Phone,
			Age:      request.Age,
			Status:   user.Employee.Status,
		},
	}

	return response
}

func (service userServiceImpl) FindByUsername(username string) model.CreateUserResponse {
	user, err := service.repository.FindByUsername(username)
	exception.PanicIfNeeded(err)

	response := model.CreateUserResponse{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
		Employee: model.CreateEmployeeResponse{
			Id:       user.Employee.Id,
			FullName: user.Employee.FullName,
			Phone:    user.Employee.Phone,
			Age:      user.Employee.Age,
			Status:   user.Employee.Status,
		},
	}

	return response
}

func (service userServiceImpl) Login(request model.CreateUserRequest) model.GetLoginResponse {
	validation.ValidateLogin(request)

	user, err := service.repository.Login(request.Username, request.Password)
	exception.PanicIfNeeded(err)

	token, err := validation.CreateToken(user)
	exception.PanicIfNeeded(err)

	response := model.GetLoginResponse{
		AccessToken: token,
		User: model.GetUserResponse{
			Id:       user.Id,
			Username: user.Username,
			Email:    user.Email,
			Employee: model.CreateEmployeeResponse{
				Id:       user.Employee.Id,
				FullName: user.Employee.FullName,
				Phone:    user.Employee.Phone,
				Age:      user.Employee.Age,
				Status:   user.Employee.Status,
			},
		},
	}

	return response
}
