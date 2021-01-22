package repository

import "github.com/sukenda/go-restful-postgres/entity"

type UserRepository interface {

	Insert(param entity.User) (entity.User, error)

	FindByUsername(username string) (entity.User, error)

	Login(username, password string) (entity.User, error)
}
