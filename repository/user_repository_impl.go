package repository

import (
	"github.com/sukenda/go-restful-postgre/entity"
	"gorm.io/gorm"
)

func NewUserRepository(database *gorm.DB) UserRepository {
	return &userRepositoryImpl{Database: database}
}

type userRepositoryImpl struct {
	Database *gorm.DB
}

func (repository userRepositoryImpl) Insert(user entity.User) error {
	/*ctx, cancel := config.NewMongoContext()
	defer cancel()

	_, err := repository.FindByUsername(user.Username)
	if err == nil {
		exception.PanicIfNeeded(exception.ValidationError{
			Message: "USER_EXIST",
		})
	}

	_, err = repository.Collection.InsertOne(ctx, bson.M{
		"_id":      user.Id,
		"username": user.Username,
		"password": user.Password,
		"email":    user.Email,
		"phone":    user.Phone,
	})
	exception.PanicIfNeeded(err)*/

	return nil
}

func (repository userRepositoryImpl) FindByUsername(username string) (user entity.User, err error) {
	/*ctx, cancel := config.NewMongoContext()
	defer cancel()

	filter := bson.M{"username": username}
	err = repository.Collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return user, errors.New("USER_NOT_FOUND")
		}
	}*/

	return user, nil
}

func (repository userRepositoryImpl) Login(username, password string) (user entity.User, err error) {
	/*_, cancel := config.NewMongoContext()
	defer cancel()

	user, err = repository.FindByUsername(username)
	if err != nil {
		exception.PanicIfNeeded(exception.ValidationError{
			Message: "USER_NOT_FOUND",
		})
	}

	exception.PanicIfNeeded(err)

	match, _ := validation.ValidatePassword(password, user.Password)
	if !match {
		exception.PanicIfNeeded(exception.ValidationError{
			Message: "PASSWORD_WRONG",
		})
	}*/

	return user, nil
}
