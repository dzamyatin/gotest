package repository

import "app/user/entity"

type UserRepositoryInterface interface {
	Create(user entity.User) error
	GetById(uid string) (entity.User, error)
}

type UserRepository struct {
}

func (u UserRepository) Create(user entity.User) error {
	//TODO implement me
	panic("implement me")
}

func (u UserRepository) GetById(uid string) (entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func InitUserRepository() UserRepositoryInterface {
	return UserRepository{}
}
