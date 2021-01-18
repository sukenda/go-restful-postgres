package repository

import "github.com/sukenda/go-restful-postgre/entity"

type UserRepository interface {

	Insert(param entity.User) (user entity.User, err error)

	FindByUsername(username string) (user entity.User, err error)

	Login(username, password string) (user entity.User, err error)
}
