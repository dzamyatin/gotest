package use_case

import (
	"app/user/internal/entity"
	"app/user/internal/repository"
	"fmt"
)

type GetUserInput struct {
	Uid string
}

type GetUserUseCase struct {
	userRepository repository.UserRepositoryInterface
}

func (c GetUserUseCase) Exec(getUserInput GetUserInput) (entity.User, error) {
	user, err := c.userRepository.GetById(getUserInput.Uid)

	if err != nil {
		return user, fmt.Errorf("Can't get user: %w", err)
	}

	return user, nil
}

func NewGetUserUseCase(repository repository.UserRepositoryInterface) GetUserUseCase {
	return GetUserUseCase{userRepository: repository}
}
