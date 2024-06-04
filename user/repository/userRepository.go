package repository

import "app/user/entity"

type UserRepositoryInterface interface {
	create(user entity.User)
	getById(uid string) (entity.User, error)
}

type UserRepository struct {
}

func (r UserRepository) create(user entity.User) {

}

func (r UserRepository) getById(uid string) entity.User {
	return entity.User{}
}
