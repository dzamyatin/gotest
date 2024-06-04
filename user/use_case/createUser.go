package use_case

import "app/user/repository"

type CreateUserDto struct {
}

type createUser struct {
	userRepository repository.UserRepository
}

func execute(dto CreateUserDto) {
}
