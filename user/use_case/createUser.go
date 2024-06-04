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

func (c CreateUserUseCase) execute(createUserInput CreateUserInput) error {
	user := entity.InitUser(createUserInput.Login)
	err := c.userRepository.Create(user)

	if err != nil {
		return fmt.Errorf("Can't create a user: %w", err)
	}

	return nil
}

func (c CreateUserUseCase) Subscribed() lib.EventInterface {
	return CreateUserInput{}
}

func (c CreateUserUseCase) Execute(i interface{}) {
	input, ok := i.(CreateUserInput)
	if !ok {
		fmt.Printf("Unexpected execution type: %v", i)
	}

	c.execute(input)
}

func InitCreateUserUseCase() CreateUserUseCase {
	return CreateUserUseCase{userRepository: repository.InitUserRepository()}
}

func init() {
	createUser := InitCreateUserUseCase()
	lib.EventBusInstance.Subscribe(createUser)
}
