package repository

import "github.com/sukenda/go-restful-postgre/entity"

type UserRepository interface {

	Insert(user entity.User) error

	FindByUsername(username string) (user entity.User, err error)

	Login(username, password string) (user entity.User, err error)
}
