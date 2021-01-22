package service

import (
	"github.com/sukenda/go-restful-postgres/model"
)

type UserService interface {

	Register(request model.CreateUserRequest) model.CreateUserResponse

	FindByUsername(username string) model.CreateUserResponse

	Login(request model.CreateUserRequest) model.GetLoginResponse
}
