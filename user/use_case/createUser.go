package use_case

import (
	"app/user/entity"
	"app/user/lib"
	"app/user/repository"
	"fmt"
)

type CreateUserInput struct {
	Login string
}

func (c CreateUserInput) GetEventName() string {
	return "CreateUser"
}

type CreateUserUseCase struct {
	userRepository repository.UserRepositoryInterface
}

func (c CreateUserUseCase) Exec(createUserInput CreateUserInput) (entity.User, error) {
	user := entity.InitUser(createUserInput.Login)
	err := c.userRepository.Create(user)

	if err != nil {
		return user, fmt.Errorf("Can't create a user: %w", err)
	}

	return user, nil
}

func (c CreateUserUseCase) Subscribed() lib.EventInterface {
	return CreateUserInput{}
}

func (c CreateUserUseCase) Execute(i interface{}) {
	input, ok := i.(CreateUserInput)
	if !ok {
		fmt.Printf("Unexpected execution type: %v", i)
	}

	c.Exec(input)
}

func InitCreateUserUseCase() CreateUserUseCase {
	return CreateUserUseCase{userRepository: repository.InitUserRepository()}
}

func init() {
	createUser := InitCreateUserUseCase()
	lib.EventBusInstance.Subscribe(createUser)
}
