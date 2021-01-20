package service

import (
	"github.com/sukenda/go-restful-postgres/model"
)

type UserService interface {

	Register(request model.CreateUserRequest) (response model.CreateUserResponse)

	FindByUsername(username string) (response model.CreateUserResponse)

	Login(request model.CreateUserRequest) (response model.GetLoginResponse)
}
