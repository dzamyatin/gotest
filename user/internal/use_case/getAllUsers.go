package use_case

import (
	"app/user/internal/entity"
	"app/user/internal/repository"
	"fmt"
)

type GetAllUsersUseCase struct {
	userRepository repository.UserRepositoryInterface
}

func (c GetAllUsersUseCase) Exec() ([]entity.User, error) {
	users, err := c.userRepository.GetAll()

	if err != nil {
		return users, fmt.Errorf("Can't get users", err)
	}

	return users, nil
}

func NewGetAllUsersUseCase(userRepository repository.UserRepositoryInterface) GetAllUsersUseCase {
	return GetAllUsersUseCase{userRepository: userRepository}
}
