package use_case

import (
	"app/user/entity"
	"app/user/repository"
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

func InitGetAllUsersUseCase() GetAllUsersUseCase {
	return GetAllUsersUseCase{userRepository: repository.InitUserRepository()}
}
