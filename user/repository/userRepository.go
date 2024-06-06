package repository

import (
	"app/user/entity"
	"app/user/lib"
)

type UserRepositoryInterface interface {
	Create(user entity.User) error
	GetById(uid string) (entity.User, error)
}

type UserRepository struct {
}

func (u UserRepository) Create(user entity.User) error {
	_, err := lib.DB.Exec(`INSERT INTO user
    (uid, login, passwordHash)
VALUES
    (?, ?, ?)`,
		user.Uid(),
		user.Login(),
		user.PasswordHash(),
	)

	return err
}

func (u UserRepository) GetById(uid string) (entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func InitUserRepository() UserRepositoryInterface {
	return UserRepository{}
}
